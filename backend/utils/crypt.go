package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic("unable to hash passwords, err="+ err.Error())
	}

	return string(hashed)
}

func ComparePasswords(hashed, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)) == nil
}