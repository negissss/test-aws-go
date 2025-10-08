package model

import (
	"time"

	"gorm.io/gorm"
)

type CryptoPrice struct {
	ID          uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Symbol      string         `gorm:"type:varchar(20);not null;index:idx_symbol_currency,unique" json:"symbol"`
	Currency    string         `gorm:"type:varchar(10);not null;index:idx_symbol_currency,unique" json:"currency"`
	Price       float64        `gorm:"type:decimal(30,10);not null" json:"price"`
	Status      string         `gorm:"type:varchar(20);default:'active'" json:"status"`
	LastUpdated time.Time      `gorm:"not null" json:"last_updated"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
