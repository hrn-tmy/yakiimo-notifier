package infra

import (
	"fmt"
	"os"
	"yakiimo-notifier/internal/domain"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	user     string
	pass     string
	host     string
	port     string
	name     string
	sslmode  string
	timezone string
}

func NewDB() (*gorm.DB, error) {
	godotenv.Load()
	cfg := DBConfig{
		user:     os.Getenv("DB_USER"),
		pass:     os.Getenv("DB_PASS"),
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		name:     os.Getenv("DB_NAME"),
		sslmode:  os.Getenv("SSL_MODE"),
		timezone: os.Getenv("DB_TIMEZONE"),
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", cfg.host, cfg.user, cfg.pass, cfg.name, cfg.port, cfg.sslmode, cfg.timezone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&domain.User{}, &domain.Favorite{}, &domain.Machine{}); err != nil {
		return nil, err
	}

	return db, nil
}
