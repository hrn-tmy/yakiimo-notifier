package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UserID                 uuid.UUID      `json:"user_id" gorm:"primaryKey"`
	Name                   string         `json:"name"`
	Email                  string         `json:"email"`
	Password               string         `json:"password"`
	NotificationPermission bool           `json:"notification_permission"`
	CreatedAt              time.Time      `json:"created_at" gorm:"type:timestamp(0)"`
	UpdatedAt              time.Time      `json:"updated_at" gorm:"type:timestamp(0)"`
	DeletedAt              gorm.DeletedAt `json:"deleted_at" gorm:"type:timestamp(0)"`
}
