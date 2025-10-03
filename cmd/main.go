package main

import (
	"context"
	cacheRepo "github.com/amir2002-js/digital-shop/internal/repository/cache"
	repository "github.com/amir2002-js/digital-shop/internal/repository/postgres"
	cacheService "github.com/amir2002-js/digital-shop/internal/services/cache"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/amir2002-js/digital-shop/internal/interface/http"
	"github.com/amir2002-js/digital-shop/internal/interface/http/handler"
	usersHandler "github.com/amir2002-js/digital-shop/internal/interface/http/handler/user"
	usersServices "github.com/amir2002-js/digital-shop/internal/services/users"
	migrationsPkg "github.com/amir2002-js/digital-shop/pkg"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	passRedis := os.Getenv("REDIS_PASSWORD")
	addrRedis := os.Getenv("REDIS_ADDR")

	client := redis.NewClient(&redis.Options{
		Addr:     addrRedis,
		Password: passRedis,
		DB:       0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := client.Ping(ctx).Result(); err != nil {
		log.Fatalf("failed to connect to redis: %v", err)
	}

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

	// layer
	dbRepo := repository.NewGormDb(db)
	cache := cacheRepo.NewRedisCacheRepo(client)

	userServe := usersServices.NewUsersServices(dbRepo)
	redisServe := cacheService.NewRedisCacheService(cache)

	userHndlr := usersHandler.NewUsersHandler(userServe, redisServe, validation)

	mainHandler := handler.NewHandler(userHndlr)

	app := fiber.New()

	http.Router(app, mainHandler)

	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Fatalf("Fiber listen failed: %v", err)
		}
	}()
	log.Println("Server started on :3000")

	// --- Graceful shutdown ---
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	if err := app.ShutdownWithContext(shutdownCtx); err != nil {
		log.Printf("Fiber shutdown error: %v", err)
	}

	if err := sqlDB.Close(); err != nil {
		log.Printf("DB close error: %v", err)
	}

	if err := client.Close(); err != nil {
		log.Printf("Redis close error: %v", err)
	}

	log.Println("Server exited cleanly")
}
