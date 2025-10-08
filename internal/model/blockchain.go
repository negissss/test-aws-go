package model

import (
	"time"

	"gorm.io/gorm"
)

type Blockchain struct {
	ID          uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string         `gorm:"type:varchar(50);uniqueIndex:idx_blockchain_name_network;not null" json:"name"`             // e.g., "Bitcoin", "Ethereum"
	Symbol      string         `gorm:"type:varchar(10);unique;not null" json:"symbol"`                                            // e.g., "BTC", "ETH"
	Network     string         `gorm:"type:varchar(20);uniqueIndex:idx_blockchain_name_network;default:'mainnet'" json:"network"` // e.g., "mainnet", "testnet"
	Status      string         `gorm:"type:varchar(20);default:'active'" json:"status"`                                           // e.g., active, maintenance, deprecated
	RPCEndpoint string         `gorm:"type:varchar(255)" json:"rpc_endpoint,omitempty"`                                           // Node RPC URL
	ChainID     string         `gorm:"type:varchar(50);uniqueIndex;null" json:"chain_id,omitempty"`                               // For EVM chains
	Convert     string         `gorm:"type:varchar(10);default:'USD'" json:"convert"`                                             // Currency for price conversion, e.g., USD, INR
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
