package utilities

import (
	"golang.org/x/crypto/bcrypt"
)

func CheckPassword(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil, err
}

func EncryptPassword(password []byte) (string, error) {
	password, err := bcrypt.GenerateFromPassword(password, 9)
	return string(password), err
}
