package main

import (
	"context"
	"github.com/amir2002-js/digital-shop/internal/interface/http/handler/gallery"
	productsHandler "github.com/amir2002-js/digital-shop/internal/interface/http/handler/products"
	tagsHandler "github.com/amir2002-js/digital-shop/internal/interface/http/handler/tags"
	cacheRepo "github.com/amir2002-js/digital-shop/internal/repository/cache"
	repository "github.com/amir2002-js/digital-shop/internal/repository/postgres"
	cacheService "github.com/amir2002-js/digital-shop/internal/services/cache"
	galleryService "github.com/amir2002-js/digital-shop/internal/services/gallery"
	productsService "github.com/amir2002-js/digital-shop/internal/services/products"
	tagService "github.com/amir2002-js/digital-shop/internal/services/tags"
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
		Addr:         addrRedis,
		Password:     passRedis,
		DB:           0,
		PoolSize:     50,
		MinIdleConns: 10,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
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
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(time.Hour)

	migrationsPkg.RunMigrations(sqlDB)

	validation := validator.New()

	// layer
	dbRepo := repository.NewGormDb(db)
	cache := cacheRepo.NewRedisCacheRepo(client)

	userServe := usersServices.NewUsersServices(dbRepo)
	productServe := productsService.NewProductsService(dbRepo)
	galleryServe := galleryService.NewGalleryService(dbRepo)
	tagsServe := tagService.NewTagService(dbRepo)
	redisServe := cacheService.NewRedisCacheService(cache)

	userHndlr := usersHandler.NewUsersHandler(userServe, redisServe, validation)
	productHndlr := productsHandler.NewProductsHandler(productServe, redisServe, validation)
	galleryHndlr := galleryHandler.NewGalleryHandler(galleryServe, redisServe, validation)
	tagsHndlr := tagsHandler.NewTagsHandler(tagsServe, redisServe, validation)

	mainHandler := handler.NewHandler(userHndlr, productHndlr, galleryHndlr, tagsHndlr)

	app := fiber.New(fiber.Config{
		ReadTimeout:       3 * time.Second,
		WriteTimeout:      3 * time.Second,
		IdleTimeout:       5 * time.Second,
		ReduceMemoryUsage: true,
		AppName:           "DigitalShop",
	})

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
