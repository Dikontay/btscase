package api

import (
	"log"
	"os"

	"github.com/Dikontay/btscase/internal/auth"
	"github.com/Dikontay/btscase/internal/auth/repository"
	"github.com/Dikontay/btscase/internal/auth/service"
	atransport "github.com/Dikontay/btscase/internal/auth/transport"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type api struct {
}

func New() *api {
	return &api{}
}

func getEnv(key, defaultValue string) string {
	value, found := os.LookupEnv(key)
	if !found {
		return defaultValue
	}
	return value
}

func (a *api) Run() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file found")
	}

	dsn := getEnv("DSN", "postgres://postgres:password@localhost/hacknu")

	eng := gin.New()
	errLogger := gin.ErrorLogger()
	logger := gin.Logger()

	eng.Use(errLogger)
	eng.Use(logger)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	authgroup := eng.Group("/auth")
	ar := repository.New(db)
	as := service.New(ar)
	at := atransport.New(as)
	auth.SetRoutes(authgroup, at)

	eng.Use(at.Authorize())

	// handlers := transport.New()

	// _ = handlers

	if err := eng.Run("localhost:4000"); err != nil {
		panic(err)
	}
}
