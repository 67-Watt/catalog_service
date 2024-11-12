package models

import (
	"github.com/google/uuid"
	"time"
)

type ItemAllergen struct {
	AllergenID   uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"allergen_id"`
	ItemID       uuid.UUID `gorm:"type:uuid;not null" json:"item_id"`
	AllergenName string    `gorm:"type:varchar(50);not null" json:"allergen_name"`
	DietaryFlag  string    `gorm:"type:varchar(50)" json:"dietary_flag,omitempty"`
	RestaurantID uuid.UUID `gorm:"type:uuid;not null" json:"restaurant_id"`
	CountryCode  string    `gorm:"type:char(2);not null" json:"country_code"`
	CreatedAt    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
}
