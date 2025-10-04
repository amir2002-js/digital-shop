package whoIs

import (
	"errors"
	"os"
)

func IsAdmin(password, email, username string) (bool, error) {
	adminUsername := os.Getenv("USERNAME_ADMIN")
	adminPassword := os.Getenv("PASSWORD_ADMIN")
	adminEmail := os.Getenv("EMAIL_ADMIN")
	if adminUsername == "" || adminPassword == "" || adminEmail == "" {
		return false, errors.New("credentials not set in environment variables")
	}

	// کاربر ادمین
	if username == adminUsername && email == adminEmail && password == adminPassword {
		return true, nil
	}

	// کاربر ساده که از ایمیل و یوزرنیم ادمین داره استفاده میکنه
	if username == adminUsername || email == adminEmail {
		return false, errors.New("username or email is already exist")
	}

	// کاربر ساده
	return false, nil
}
