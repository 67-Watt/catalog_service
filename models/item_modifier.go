package models

import (
	"github.com/google/uuid"
	"time"
)

type ItemModifier struct {
	ItemModifierID uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"item_modifier_id"`
	ItemID         uuid.UUID `gorm:"type:uuid;not null" json:"item_id"`
	ModifierID     uuid.UUID `gorm:"type:uuid;not null" json:"modifier_id"`
	RestaurantID   uuid.UUID `gorm:"type:uuid;not null" json:"restaurant_id"`
	CountryCode    string    `gorm:"type:char(2);not null" json:"country_code"`
	CreatedAt      time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
}
