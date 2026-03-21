package usecase

import (
	"yakiimo-notifier/internal/gen"
	"yakiimo-notifier/internal/repository"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	repo repository.IUserRepository
}

// NewUserUsecase はUserUsecaseを生成します
func NewUserUsecase(repo repository.IUserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

// CreateUser は会員登録を実行します
func (uu *UserUsecase) CreateUser(email, name, password string) (gen.CreateUserResponseData, error) {
	userID, err := uuid.NewV7()
	if err != nil {
		return gen.CreateUserResponseData{}, err
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return gen.CreateUserResponseData{}, err
	}
	data, err := uu.repo.CreateUser(userID, email, name, string(hashed))
	if err != nil {
		return gen.CreateUserResponseData{}, err
	}

	user := gen.CreateUserResponseData{
		Name:                   data.Name,
		Email:                  data.Email,
		NotificationPermission: data.NotificationPermission,
	}

	return user, nil
}
