package models

import (
	"github.com/google/uuid"
	"time"
)

type MenuItem struct {
	ItemID          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"item_id"`
	Name            string    `gorm:"type:varchar(50);not null" json:"name"`
	Description     string    `gorm:"type:text" json:"description,omitempty"`
	AvailableStatus bool      `gorm:"type:boolean;default:true" json:"available_status"`
	PreparationTime int       `gorm:"type:int" json:"preparation_time"`
	IsCustomizable  bool      `gorm:"type:boolean;default:false" json:"is_customizable"`
	CategoryID      uuid.UUID `gorm:"type:uuid" json:"category_id"`
	RestaurantID    uuid.UUID `gorm:"type:uuid;not null" json:"restaurant_id"`
	CountryCode     string    `gorm:"type:char(2);not null" json:"country_code"`
	CreatedAt       time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
