package model

import (
	"github.com/google/uuid"
	"time"
)

type ItemPrice struct {
	PriceID      uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"price_id"`
	ItemID       uuid.UUID  `gorm:"type:uuid;not null" json:"item_id"`
	Price        float64    `gorm:"type:decimal(15,2);not null" json:"price"`
	CostPrice    float64    `gorm:"type:decimal(15,2)" json:"cost_price,omitempty"`
	Currency     string     `gorm:"type:char(3);default:'USD'" json:"currency"`
	StartDate    time.Time  `gorm:"type:date;not null" json:"start_date"`
	EndDate      *time.Time `gorm:"type:date" json:"end_date,omitempty"`
	RestaurantID uuid.UUID  `gorm:"type:uuid;not null" json:"restaurant_id"`
	CountryCode  string     `gorm:"type:char(2);not null" json:"country_code"`
	CreatedAt    time.Time  `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
