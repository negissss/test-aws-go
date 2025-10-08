package address

import (
	"api-service/internal/common"

	"github.com/gin-gonic/gin"
)

// AddressHandler handles HTTP requests related to addresses.
type AddressHandler struct {
	ApiService ApiServiceI
}

// AddressRequest represents the expected request payload for setting an address.
type AddressRequest struct {
	UserAddress string `json:"address" binding:"required"`
}

// NewAddressHandler creates a new instance of AddressHandler.
func NewAddressHandler(intentService ApiServiceI) *AddressHandler {
	return &AddressHandler{ApiService: intentService}
}

func (h *AddressHandler) HealthCheck(c *gin.Context) {
	common.SuccessResponse(c, common.MsgHealthOK, gin.H{
		"status":  "ok",
		"service": "crypto-api-service",
	})

}
