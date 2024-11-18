package tools

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"regexp"
)

func ValidateEmail(email string) bool {
	re := `\S+@\S+\.\S+`
	matched := regexp.MustCompile(re).MatchString(email)
	return matched
}

func SendActivationEmail(to string, token string) bool {
	domain := os.Getenv("MAILGUN_FROM_DOMAIN")
	apiKey := os.Getenv("MAILGUN_API_KEY")
	frontendOrigin := os.Getenv("FRONTEND_ORIGIN")

	apiEndpoint := fmt.Sprintf("https://api.mailgun.net/v3/%s/messages", domain)

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	writer.WriteField("from", fmt.Sprintf("Real-time Notification <postmaster@%s>", domain))
	writer.WriteField("to", to)
	writer.WriteField("subject", "Please activate your account!")
	writer.WriteField("text", "Please activate your account!")
	writer.WriteField("html", ActivationEmail(fmt.Sprintf("%s/verify-email?email=%s&token=%s", frontendOrigin, to, token)))

	writer.Close()

	req, err := http.NewRequest("POST", apiEndpoint, &body)
	if err != nil {
		log.Println("Error creating request:", err)
		return false
	}

	authHeader := "Basic " + base64.StdEncoding.EncodeToString([]byte("api:"+apiKey))
	req.Header.Set("Authorization", authHeader)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request:", err)
		return false
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response:", err)
		return false
	}

	fmt.Println("Response:", string(respBody))
	return resp.StatusCode == http.StatusOK
}
