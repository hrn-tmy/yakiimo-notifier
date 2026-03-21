package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Machine struct {
	MachineID uuid.UUID
	Name      string
	Location  string
	IsActive  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
