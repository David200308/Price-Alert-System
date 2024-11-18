package controllers

import (
	"os"
	"strings"

	"github.com/David200308/go-api/Backend/models"
	"github.com/David200308/go-api/Backend/mq"
	"github.com/David200308/go-api/Backend/services"
	"github.com/David200308/go-api/Backend/tools"
	"github.com/gin-gonic/gin"
)

// @Router /user/register [Post]
// @Param username formData string true "Username"
// @Param email formData string true "Email"
// @Param password formData string true "Password"
// @Success 200
// @Failure 400
// @Failure 500
func CreateUser(c *gin.Context) {
	user := models.User{
		Username: c.PostForm("username"),
		Email:    c.PostForm("email"),
		Password: c.PostForm("password"),
	}

	if user.Username == "" || user.Email == "" || user.Password == "" {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Username, Email and password are required",
		})
		return
	}

	if !tools.ValidateEmail(strings.ToLower(user.Email)) {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Invalid email",
		})
		return
	}

	userUuid, err := services.InsertUser(&user)
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": "Failed to create user",
		})
		return
	}

	emailVerifyToken, err := tools.GenerateToken(strings.ToLower(user.Email), userUuid, "email_verification", true)
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": "Failed to generate token",
		})
		return
	}

	token, err := tools.GenerateToken(strings.ToLower(user.Email), userUuid, "signup", true)
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": "Failed to generate token",
		})
		return
	}

	tools.SendActivationEmail(strings.ToLower(user.Email), emailVerifyToken)

	c.SetCookie("token", token, 60*5, "/", os.Getenv("FRONTEND_DOMAIN"), true, true)

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Registration successful, please verify your email",
	})
}

// @Router /user/login [Post]
// @Param email formData string true "Email"
// @Param password formData string true "Password"
// @Success 200
// @Failure 400
// @Failure 500
func LoginUser(c *gin.Context) {
	user := models.User{
		Email:    c.PostForm("email"),
		Password: c.PostForm("password"),
	}

	if user.Email == "" || user.Password == "" {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Email and password are required",
		})
		return
	}

	res, err := services.GetUserByEmail(strings.ToLower(user.Email))
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "User not found",
		})
		return
	}

	existingUser := *res
	if existingUser.Status != "active" {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "User is not active",
		})
		return
	}

	if !tools.VerifyUserPassword(user.Password, existingUser.Password) {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Invalid password",
		})
		return
	}

	token, err := tools.GenerateToken(strings.ToLower(user.Email), existingUser.UserUUID, "auth", false)
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": "Failed to generate token",
		})
		return
	}

	c.SetCookie("token", token, 60*60*1, "/", os.Getenv("FRONTEND_DOMAIN"), false, true)

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Login successful",
	})
}

// @Router /user/token [Post]
// @Success 200
// @Failure 400
// @Failure 500
func TokenVerification(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Unauthorized",
		})
		return
	}

	res, err := tools.VerifyToken(token)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Unauthorized",
		})
		return
	}

	if res == nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Unauthorized",
		})
		return
	}

	c.JSON(200, gin.H{
		"status":   "success",
		"message":  "Token verified",
		"function": res["function"],
	})
}

// @Router /user/email/verification [Post]
// @Param email_verify_token formData string true "Email Verification Token"
// @Param email formData string true "Email"
// @Success 200
// @Failure 400
// @Failure 500
func EmailVerification(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Unauthorized",
		})
		return
	}

	emailVerificationToken := c.PostForm("email_verify_token")
	email := c.PostForm("email")
	if emailVerificationToken == "" || email == "" {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Email verification token and email are required",
		})
		return
	}

	res, err := tools.VerifyToken(token)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Unauthorized",
		})
		return
	}

	uuid := res["sub"].(string)
	if email != res["email"].(string) {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Unauthorized",
		})
		return
	}

	if uuid == "" {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Unauthorized",
		})
		return
	}

	if res["function"] != "signup" {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Unauthorized",
		})
		return
	}

	emailTokenRes, err := tools.VerifyToken(emailVerificationToken)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Failed to verify email token",
		})
		return
	}

	if emailTokenRes["email"] != email {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Email does not match",
		})
		return
	}
	if (emailTokenRes["sub"].(string)) != uuid {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "UUID does not match",
		})
		return
	}

	if emailTokenRes["function"] != "email_verification" {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Function does not match",
		})
		return
	}

	err = services.UpdateUserStatus(uuid, "active")
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": "Failed to update user status",
		})
		return
	}

	mq.UserVerify(c, uuid, email)

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Email verified",
	})
}

// @Router /user [Get]
// @Success 200
// @Failure 400
// @Failure 404
func GetUser(c *gin.Context) {
	uuid, email, err := tools.NormalRequestVerifyToken(c)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Unauthorized",
			"error":   err,
		})
		return
	}

	userData, err := services.GetUserByUUIDAndEmail(uuid, email)

	if err != nil {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "User not found or not active",
		})
		return
	}

	mq.UserCreated(c, userData.UserUUID, userData.Email)

	userData.Password = ""

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Get user successful",
		"data":    userData,
	})
}

// @Router /user/logout [Post]
// @Success 200
// @Failure 400
func LogoutUser(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Unauthorized",
		})
		return
	}

	res, err := tools.VerifyToken(token)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Unauthorized",
		})
		return
	}

	if res == nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Unauthorized",
		})
		return
	}

	c.SetCookie("token", "", 0, "/", os.Getenv("FRONTEND_DOMAIN"), true, true)

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Logout successful",
	})
}

// func UpdateUser(c *gin.Context) {
// 	c.JSON(200, gin.H{
// 		"status":  "success",
// 		"message": "User updated",
// 	})
// }

// func DeleteUser(c *gin.Context) {
// 	c.JSON(200, gin.H{
// 		"status":  "success",
// 		"message": "User deleted",
// 	})
// }
