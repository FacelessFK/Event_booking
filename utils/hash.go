package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {

	byteHash, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(byteHash), err
}

func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil // give me true read  bcrypt.CompareHashAndPassword return
}
