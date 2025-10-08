package address

import (
	"github.com/gin-gonic/gin"
)

// RegisterAddressRoutes registers all address-related routes.
func RegisterAddressRoutes(router gin.IRoutes, handler *AddressHandler) {
	router.GET("/health", handler.HealthCheck)
}
