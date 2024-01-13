package schema

import (
	"errors"
	"regexp"
	"strings"
	"unicode/utf8"
)

type RegistrationSchema struct {
	Login    string `json:"login" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
	FullName string `json:"full_name" validate:"required"`
}

func (me RegistrationSchema) Validate() error {

	if strings.TrimSpace(me.FullName) == "" {
		return errors.New("name cannot be empty")
	}

	if utf8.RuneCountInString(me.FullName) >= 50 {
		return errors.New("name cannot exceed 100 characters")
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`)
	if !emailRegex.MatchString(me.Email) {
		return errors.New("email is entered incorrectly")
	}

	return nil
}
