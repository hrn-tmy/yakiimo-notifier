package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Favorite struct {
	FavoriteID uuid.UUID
	UserID     uuid.UUID
	MachineID  uuid.UUID
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}
