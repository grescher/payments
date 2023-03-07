package models

import (
	"errors"
	"regexp"
)

type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	IsActive     bool   `json:"is_active"`
}

func (u *User) SignUpValidation() (err error) {
	if isEmpty(u.Name) {
		return errors.New("name of the user is empty")
	}
	if isEmpty(u.Email) {
		return errors.New("email of the user is empty")
	}
	if !isEmailValid(u.Email) {
		return errors.New("email of the user is not valid")
	}
	if isEmpty(u.PasswordHash) {
		return errors.New("password of the user is empty")
	}
	return nil
}

func isEmpty(s string) bool {
	return s == ""
}

func isEmailValid(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}
