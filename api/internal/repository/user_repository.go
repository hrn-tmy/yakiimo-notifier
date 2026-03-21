package repository

import (
	"time"
	"yakiimo-notifier/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(userID uuid.UUID, email, name, password string) (domain.User, error)
	GetTargetUsers(machineID string) ([]string, error)
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

func (ur *UserRepository) GetTargetUsers(machineID string) ([]string, error) {
	var targets []string
	err := ur.db.Table("favorites").
		Select("users.email").
		Joins("JOIN users ON users.user_id = favorites.user_id").
		Where("favorites.machine_id = ?", machineID).
		Where("notification_permission = ?", "TRUE").
		Scan(&targets).
		Error
	if err != nil {
		return nil, err
	}

	return targets, nil
}