package helper

import (
	"fmt"
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

func ValidateNik(nik string) bool {
	nikRegex := regexp.MustCompile(`^[0-9]*$`)
	length := len(nik)
	if length !=16 {
		return false
	}

	return nikRegex.MatchString(nik)
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

func LatLongIsEmpty(data float64)bool{
	s := fmt.Sprintf("%f", data)
	if s!=""{return true}else{return false}
}