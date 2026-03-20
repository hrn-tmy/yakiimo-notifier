package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UserID                 uuid.UUID
	Name                   string
	Email                  string
	Password               string
	NotificationPermission bool
	CreatedAt              time.Time
	UpdatedAt              time.Time
	DeletedAt              gorm.DeletedAt
}
