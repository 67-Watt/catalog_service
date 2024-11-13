package model

import (
	"github.com/google/uuid"
	"time"
)

type ItemAvailability struct {
	AvailabilityID uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"availability_id"`
	ItemID         uuid.UUID  `gorm:"type:uuid;not null" json:"item_id"`
	StartDate      *time.Time `gorm:"type:date" json:"start_date,omitempty"`
	EndDate        *time.Time `gorm:"type:date" json:"end_date,omitempty"`
	IsAvailable    bool       `gorm:"type:boolean;default:true" json:"is_available"`
	RestaurantID   uuid.UUID  `gorm:"type:uuid;not null" json:"restaurant_id"`
	CountryCode    string     `gorm:"type:char(2);not null" json:"country_code"`
	CreatedAt      time.Time  `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
}
