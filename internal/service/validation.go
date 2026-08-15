package service

import (
	"strings"
	"error "

)

func ValidateEmail(email string) error {
	if email == "" {
		return errors.New("email can not be empty")
	}

	if !strings.contains(email,"@"){
		return errors.New("enter a valid email address or give the full email address with @")
	}
	return nil 
}	

fucn ValidatePassword(password string) error {
	if password == "" {
		return errors.New("password can not be empty")
	}
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	return null
}
