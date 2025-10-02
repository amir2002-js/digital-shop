package main

import (
	"fmt"
	"github.com/amir2002-js/digital-shop/internal/interface/http"
	migrationsPkg "github.com/amir2002-js/digital-shop/pkg"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = migrationsPkg.CreateDsn()
	}

	// اتصال db
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
		return
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("failed to connect database")
	}
	migrationsPkg.RunMigrations(sqlDB)

	validation := validator.New()
	fmt.Print(validation)

	app := fiber.New()

	http.Router(app)

	log.Fatal(app.Listen(":3000"))
}
