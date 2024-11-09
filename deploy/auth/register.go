package auth

import (
	"errors"
	"example/v3/db"
	"example/v3/models"
	"example/v3/utils"
	"fmt"
)

func SignUpUser(user *models.User) error {
	db, err := db.OpenDBConnection()
	if err != nil {
		return err
	}

	fmt.Println("DB connection established...")

	//is Existed LOGIN?
	var existingUser *models.User
	if err := db.Where("login = ?", user.Login).First(&existingUser).Error; err == nil {
		return errors.New("user already exists")
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.PassHash = hashedPassword

	if err := db.Create(&user).Error; err != nil {
		return fmt.Errorf("error for registration: %v", err)
	}

	return nil
}
