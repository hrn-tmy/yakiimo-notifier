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

// NewUserRepository はUserRepositoryを生成します
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}

// CreateUser は新しい会員をDBに登録し、登録した会員データを返す
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

// GetTargetUsers は指定した機械IDをお気に入り登録しており、通知許可フラグがONの会員のメールアドレス一覧を返す
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
