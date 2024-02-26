package utils

import (
	"errors"
	"math/rand"
	"os"
	"strings"
)

func GenerateRandomString(length int) string {
    const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    result := make([]byte, length)
    for i := range result {
        result[i] = charset[rand.Intn(len(charset))]
    }
    return string(result)
}

func GenerateExternalId(length int) string {
    const charset = "0123456789"
    result := make([]byte, length)
    for i := range result {
        result[i] = charset[rand.Intn(len(charset))]
    }
    return "FS-" + string(result)
}

func GenerateEmailVerificationBody(name, url string) (emailBody string, err error) {
	configPath := "./resources/template/verification_email.html"
	tempByte, err := os.ReadFile(configPath)
	if err != nil {
		return
	}

	emailBody = string(tempByte)
	emailBody = strings.ReplaceAll(emailBody, "{.url_verification}", url)
    emailBody = strings.ReplaceAll(emailBody, "{.name}", name)
	return
}

func CorrectPhoneNumber(phoneNumber string) string {
	if strings.HasPrefix(phoneNumber, "0") {
        phoneNumber = "62" + phoneNumber[1:]
    }
	return phoneNumber
}

func GenerateEmailVerificationSuccesBody(name string) (emailBody string) {
	configPath := "./resources/template/verification_email_success.html"
	tempByte, err := os.ReadFile(configPath)
	if err != nil {
		return
	}

	emailBody = string(tempByte)
    emailBody = strings.ReplaceAll(emailBody, "{.name}", name)
	return
}

func ValidationImages(imageName string, imageSize int) error {
	name := strings.ToLower(imageName)
	if !strings.HasSuffix(strings.ToLower(name), ".jpg") && !strings.HasSuffix(strings.ToLower(name), ".png") && !strings.HasSuffix(strings.ToLower(name), ".jpeg") {
		return errors.New("only supported file formats are jpg, jpeg and png")
	}

	if imageSize > 5003000 {
		return errors.New("image size cannot be more than 5 MB")
	}

	return nil
}


