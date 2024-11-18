package services

import (
	"log"
	"strings"

	"github.com/David200308/go-api/Backend/initializers"
	"github.com/David200308/go-api/Backend/models"
	"github.com/David200308/go-api/Backend/tools"
)

func InsertUser(user *models.User) (string, error) {
	hashedPassword, err := tools.HashingPassword(user.Password)
	if err != nil {
		log.Println("Error hashing password:", err)
		return "", err
	}
	user.Email = strings.ToLower(user.Email)
	user.Password = hashedPassword
	user.Status = "pending"
	user.UserUUID = tools.GenerateUUID()

	if err := initializers.DB.Create(&user).Error; err != nil {
		log.Println("Error inserting user into database:", err)
		return "", err
	}

	return user.UserUUID, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := initializers.DB.Where("email = ?", email).First(&user).Error; err != nil {
		log.Println("Error getting user:", err)
		return nil, err
	}

	return &user, nil
}

func GetUserByUUIDAndEmail(uuid, email string) (*models.User, error) {
	var user models.User
	if err := initializers.DB.Where("user_uuid = ? AND email = ? AND status = 'active'", uuid, email).First(&user).Error; err != nil {
		log.Println("Error getting user:", err)
		return nil, err
	}

	return &user, nil
}

func UpdateUserStatus(uuid, email string) error {
	if err := initializers.DB.Model(&models.User{}).
		Where("user_uuid = ? AND email = ? AND status = ?", uuid, email, "pending").
		Update("status", "active").
		Error; err != nil {
		log.Println("Error updating user status:", err)
		return err
	}

	return nil
}
