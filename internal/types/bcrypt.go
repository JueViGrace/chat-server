package types

import "golang.org/x/crypto/bcrypt"

func HashPassword(pass string) (string, error) {

	encpw, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	encPass := string(encpw)

	return encPass, nil
}

func ValidatePassword(reqPass, encPass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(encPass), []byte(reqPass)) == nil
}
