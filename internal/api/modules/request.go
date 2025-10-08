package address

type BalanceRequest struct {
	Address    string `uri:"address" binding:"required"`
	Blockchain string `form:"blockchain"`
	ChainID    string `form:"chain_id"`
	Contract   string `form:"contract"`
}
