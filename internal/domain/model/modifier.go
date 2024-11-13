package model

import (
	"github.com/google/uuid"
	"time"
)

type Modifier struct {
	ModifierID   uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"modifier_id"`
	Name         string    `gorm:"type:varchar(50);not null" json:"name"`
	Description  string    `gorm:"type:text" json:"description,omitempty"`
	RestaurantID uuid.UUID `gorm:"type:uuid;not null" json:"restaurant_id"`
	CountryCode  string    `gorm:"type:char(2);not null" json:"country_code"`
	CreatedAt    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
