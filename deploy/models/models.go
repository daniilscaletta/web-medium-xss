package models

import (
	"database/sql"
	"errors"
	"example/v3/db"
	"example/v3/utils"
	"fmt"
)

type User struct {
	Name           string
	Surname        string
	DateOfBirthday string
	Email          string
	PhoneNumber    string
	Passport       string
	Login          string
	Password       string
	PassHash       string
}

func RegisterUser(user *User) error {
	db, err := db.OpenDBConnection()
	if err != nil {
		return err
	}

	fmt.Println("DB connection established...")
	defer db.Close()

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO users (name, surname, dateofbirthday ,email, phonenumber, passport, login, passhash) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		user.Name, user.Surname, user.DateOfBirthday, user.Email, user.PhoneNumber, user.Passport, user.Login, hashedPassword)
	if err != nil {
		return fmt.Errorf("ошибка при регистрации пользователя: %v", err)
	}

	return nil
}

func AuthenticateUser(login, password string) (*User, error) {
	db, err := db.OpenDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var user User
	err = db.QueryRow("SELECT Login, PassHash FROM users WHERE login = ?", login).Scan(&user.Login, &user.PassHash)
	if err == sql.ErrNoRows {
		return nil, errors.New("пользователь не найден")
	} else if err != nil {
		return nil, err
	}

	if !utils.CheckPasswordHash(password, user.PassHash) {
		return nil, errors.New("неверный пароль")
	}

	return &user, nil
}
