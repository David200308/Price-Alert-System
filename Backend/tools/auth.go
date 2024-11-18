package tools

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func HashingPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func VerifyUserPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateUUID() string {
	newUUID := uuid.New()
	return newUUID.String()
}

func GenerateToken(email string, uuid string, function string, expFast bool) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"email": email,
		"sub":   uuid,
		"exp": func() int64 {
			if expFast {
				return time.Now().Add(time.Minute * 5).Unix()
			}
			return time.Now().Add(time.Hour).Unix()
		}(),
		"function": func() string {
			switch function {
			case "auth":
				return "auth"
			case "email_verification":
				return "email_verification"
			case "signup":
				return "signup"
			case "password_reset":
				return "password_reset"
			default:
				return ""
			}
		}(),
	})

	privateKeyPEM := os.Getenv("JWT_PRIVATE_KEY")
	block, _ := pem.Decode([]byte(privateKeyPEM))
	if block == nil || block.Type != "EC PRIVATE KEY" {
		return "", errors.New("failed to decode EC private key")
	}

	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	publicKeyPEM := os.Getenv("JWT_PUBLIC_KEY")

	if publicKeyPEM == "" {
		return nil, errors.New("JWT_PUBLIC_KEY environment variable is not set")
	}

	block, _ := pem.Decode([]byte(publicKeyPEM))

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}

	ecdsaPublicKey, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("public key is not of type *ecdsa.PublicKey")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return ecdsaPublicKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	validFunctions := map[string]struct{}{
		"auth":               {},
		"email_verification": {},
		"signup":             {},
		"password_reset":     {},
	}

	functionClaim, ok := claims["function"].(string)
	if !ok || functionClaim == "" {
		return nil, errors.New("invalid or missing 'function' claim")
	}

	if _, valid := validFunctions[functionClaim]; !valid {
		return nil, errors.New("invalid function claim value")
	}

	return claims, nil
}

type IPLocationResponse struct {
	CountryName string `json:"country_name"`
}

func GetIPLocation(ipAddress string) (string, error) {
	url := "https://api.iplocation.net/?ip=" + ipAddress

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var ipLocation IPLocationResponse
	if err := json.NewDecoder(resp.Body).Decode(&ipLocation); err != nil {
		return "", err
	}

	return ipLocation.CountryName, nil
}

func GetIPDeviceNameLocation(c *gin.Context) (string, string, string, error) {
	loginIPAddress := c.GetHeader("X-Forwarded-For")
	if loginIPAddress == "" {
		ip, _, err := net.SplitHostPort(c.Request.RemoteAddr)
		if err != nil {
			return "", "", "", err
		}
		loginIPAddress = ip
	} else {
		if strings.Contains(loginIPAddress, ",") {
			loginIPAddress = strings.TrimSpace(strings.Split(loginIPAddress, ",")[0])
		}
	}

	device := c.GetHeader("User-Agent")

	location, err := GetIPLocation(loginIPAddress)
	if err != nil {
		return "", "", "", err
	}

	return loginIPAddress, device, location, nil
}

func NormalRequestVerifyToken(c *gin.Context) (string, string, error) {
	token, err := c.Cookie("token")
	if err != nil {
		return "", "", err
	}

	res, err := VerifyToken(token)
	if err != nil {
		return "", "", err
	}

	uuid := res["sub"].(string)
	email := res["email"].(string)

	if uuid == "" || email == "" {
		return "", "", errors.New("invalid uuid or email claim value")
	}

	if res["function"] != "auth" {
		return "", "", errors.New("invalid function claim value")
	}

	return uuid, email, nil
}
