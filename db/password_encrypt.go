package db

import "golang.org/x/crypto/bcrypt"

// PasswordEncrypt returns the bcrypt hash of the given password
func PasswordEncrypt(password string) (string, error) {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}
