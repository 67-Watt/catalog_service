package model

import (
	"github.com/google/uuid"
	"time"
)

type Promotion struct {
	PromotionID        uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"promotion_id"`
	Name               string     `gorm:"type:varchar(50);not null" json:"name"`
	Description        string     `gorm:"type:text" json:"description,omitempty"`
	DiscountPercentage float64    `gorm:"type:decimal(5,2)" json:"discount_percentage"`
	DiscountAmount     float64    `gorm:"type:decimal(15,2)" json:"discount_amount"`
	StartDate          time.Time  `gorm:"type:date;not null" json:"start_date"`
	EndDate            *time.Time `gorm:"type:date" json:"end_date,omitempty"`
	RestaurantID       uuid.UUID  `gorm:"type:uuid" json:"restaurant_id"`
	CountryCode        string     `gorm:"type:char(2)" json:"country_code"`
	CreatedAt          time.Time  `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt          time.Time  `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
