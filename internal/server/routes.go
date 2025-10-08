package server

import (
	address "api-service/internal/api/modules"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func HandleOption(c *gin.Context) {
	allowedOriginsStr := os.Getenv("ALLOWED_ORIGINS")
	requestOrigin := c.Request.Header.Get("Origin")

	// By default, block
	allowOrigin := ""

	// Special case: allow any origin by echoing back
	if allowedOriginsStr == "*" {
		allowOrigin = requestOrigin
	} else {
		allowedOrigins := make(map[string]bool)
		for _, origin := range strings.Split(allowedOriginsStr, ",") {
			allowedOrigins[strings.TrimSpace(origin)] = true
		}
		if allowedOrigins[requestOrigin] {
			allowOrigin = requestOrigin
		}
	}

	if allowOrigin != "" {
		c.Header("Access-Control-Allow-Origin", allowOrigin)
		c.Header("Access-Control-Allow-Credentials", "true")
	}

	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, withcredentials, X-CSRF-Token")
	c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

	// Prevent caching
	c.Header("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
	c.Header("Pragma", "no-cache")
	c.Header("Expires", "0")

	// Handle preflight
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Next()
}

func ConfigRoutesAndSchedulers(server *Server) {
	server.Gin.Use(HandleOption)
	// evmService, err := ethprovider.NewEvmService(
	// 	server.Cfg.EVM.EthRpcURL,
	// )
	// if err != nil {
	// 	log.Fatal("Failed to create evm provider: ", err)
	// }
	// fmt.Println("ethSyncService:::", evmService)

	// addressRepo := repository.NewCmcRepository(server.DB)
	addressService := address.NewApiService()
	addressHandler := address.NewAddressHandler(addressService)
	intentRoutes := server.Gin.Group("/api/v1")
	address.RegisterAddressRoutes(intentRoutes, addressHandler)
	// cmcProvider := cmcprovider.NewCMCProvider(server.Cfg.CMC.CmcApiKey)
	// cmcService := service.NewPriceService(server.DB, cmcProvider)
	// cmcScheduler := cron.NewPriceSyncScheduler(cmcService)
	// fmt.Println("cmcScheduler:::", cmcScheduler)
	// cmcScheduler.Start()

}
