package main

import (
	"errors"
	"net/mail"
	"strings"
)

func ValidateUser(u User) error {
	if err := ValidateName(u.Name); err != nil {
		return err
	}

	if err := ValidateEmail(u.Email); err != nil {
		return err
	}

	if err := ValidateAge(u.Age); err != nil {
		return err
	}

	return nil
}

func ValidateName(name string) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("name cannot be empty")
	}

	if len(name) > 100 {
		return errors.New("name is too long")
	}

	return nil
}

func ValidateEmail(email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return errors.New("invalid email")
	}

	return nil
}

func ValidateAge(age int) error {
	if age < 0 || age > 150 {
		return errors.New("invalid age")
	}

	return nil
}