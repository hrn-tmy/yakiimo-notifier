package repository

import (
	"time"
	"yakiimo-notifier/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(userID uuid.UUID, email, name, password string) (domain.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) CreateUser(userID uuid.UUID, email, name, password string) (domain.User, error) {
	now := time.Now().Truncate(time.Second)
	user := domain.User{
		UserID:                 userID,
		Name:                   name,
		Email:                  email,
		Password:               password,
		NotificationPermission: true,
		CreatedAt:              now,
		UpdatedAt:              now,
	}
	if err := ur.db.Create(&user).Error; err != nil {
		return domain.User{}, err
	}

	return user, nil
}
