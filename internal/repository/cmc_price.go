package repository

import (
	"api-service/internal/model"

	"gorm.io/gorm"
)

// AddressRepositoryI defines the repository interface for address operations.
type AddressRepositoryI interface {
	FindAll() ([]model.CryptoPrice, error)
	Create(address *model.CryptoPrice) (*model.CryptoPrice, error)
}

// AddressRepository is the concrete implementation of AddressRepositoryI.
type AddressRepository struct {
	db *gorm.DB
}

// NewAddressRepository creates a new AddressRepository instance.
func NewCmcRepository(db *gorm.DB) AddressRepositoryI {
	return &AddressRepository{db: db}
}

// FindAll fetches all addresses from the database.
func (r *AddressRepository) FindAll() ([]model.CryptoPrice, error) {
	var blockchains []model.CryptoPrice
	if err := r.db.Find(&blockchains).Error; err != nil {
		return nil, err
	}
	return blockchains, nil
}

// Create inserts a new address into the database.
func (r *AddressRepository) Create(Address *model.CryptoPrice) (*model.CryptoPrice, error) {
	if err := r.db.Create(Address).Error; err != nil {
		return nil, err
	}
	return Address, nil
}
