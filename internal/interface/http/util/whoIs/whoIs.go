package whoIs

import (
	"errors"
	"os"
)

func IsAdmin(password, email, username string) (bool, error) {
	adminUsername := os.Getenv("USERNAME_ADMIN")
	adminPassword := os.Getenv("PASSWORD_ADMIN")
	adminEmail := os.Getenv("EMAIL_ADMIN")

	if username == adminUsername && email == adminEmail && password == adminPassword {
		return true, nil
	}

	if username == adminUsername || email == adminEmail {
		return false, errors.New("invalid email")
	}

	return false, nil
}
