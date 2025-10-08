package cron

import (
	"api-service/internal/service"
	"log"

	"github.com/robfig/cron/v3"
)

type PriceSyncScheduler struct {
	cronJob *cron.Cron
	service service.PriceService
}

func NewPriceSyncScheduler(s service.PriceService) *PriceSyncScheduler {
	return &PriceSyncScheduler{
		cronJob: cron.New(),
		service: s,
	}
}

func (p *PriceSyncScheduler) Start() {
	if _, err := p.cronJob.AddFunc("@every 30s", func() {
		if err := p.service.SyncPrices(); err != nil {
			log.Printf("Price sync failed: %v", err)
		}
	}); err != nil {
		log.Fatalf("Failed to schedule price sync job: %v", err)
	}

	p.cronJob.Start()
}
