package service

import (
	"api-service/internal/model"
	cmcprovider "api-service/internal/provider/cmc"
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PriceService interface {
	SyncPrices() error
}

type priceService struct {
	db       *gorm.DB
	provider cmcprovider.CMCProvider
}

func NewPriceService(db *gorm.DB, provider cmcprovider.CMCProvider) PriceService {
	return &priceService{db: db, provider: provider}
}

func (s *priceService) SyncPrices() error {
	// 1. Load active blockchains from DB
	var blockchains []model.Blockchain
	if err := s.db.Where("status = ?", "active").Find(&blockchains).Error; err != nil {
		return fmt.Errorf("failed to fetch blockchains: %w", err)
	}

	// 2. Group blockchains by convert currency
	currencyMap := make(map[string][]string) // convert -> []symbols
	for _, bc := range blockchains {
		cur := bc.Convert
		currencyMap[cur] = append(currencyMap[cur], bc.Symbol)
	}

	// 3. For each currency, fetch prices from CMC and upsert
	for convert, symbols := range currencyMap {
		prices, err := s.provider.GetPrice(symbols, convert)
		fmt.Println("prices:::", prices)
		if err != nil {
			return fmt.Errorf("failed to fetch prices for %s: %w", convert, err)
		}

		for sym, price := range prices {
			record := model.CryptoPrice{
				Symbol:      sym,
				Currency:    convert,
				Price:       price,
				LastUpdated: time.Now().UTC(),
			}

			if err := s.db.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "symbol"}, {Name: "currency"}},
				DoUpdates: clause.AssignmentColumns([]string{"price", "last_updated", "updated_at"}),
			}).Create(&record).Error; err != nil {
				return fmt.Errorf("failed to upsert price for %s-%s: %w", sym, convert, err)
			}
		}
	}

	return nil
}
