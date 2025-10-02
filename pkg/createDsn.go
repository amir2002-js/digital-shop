package migrationsPkg

import (
	"fmt"
	"os"
)

func CreateDsn() (dsn string) {
	name := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")

	if port == "" || host == "" || user == "" || password == "" || name == "" {
		return
	}

	dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", host, user, password, name, port)
	return
}
