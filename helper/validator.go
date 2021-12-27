package helper

import (
	"regexp"
	"strings"
)

func ValidateEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	length := len(email)
	if length < 3 || length > 254 {
		return false
	}

	return emailRegex.MatchString(email)
}

func ValidatePhoneNumber(phoneNumber string) bool {
	phoneNumberRegex := regexp.MustCompile(`^[0-9]*$`)
	length := len(phoneNumber)
	if length < 8 || length > 15 {
		return false
	}

	return phoneNumberRegex.MatchString(phoneNumber)
}

func ValidatePassword(password string) bool {
	length := len(password)
	if length < 6 || length > 100 {
		return false
	}

	return true
}

func IsEmpty(str string) bool {
	return len(strings.Trim(str, " ")) == 0
}
