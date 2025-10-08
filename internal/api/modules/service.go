package address

// ApiServiceI defines the service interface for address operations.
type ApiServiceI interface {
}

type CryptoPricesResponse struct {
	Data []map[string]map[string]float64 `json:"data"`
}

// ApiService is the concrete implementation of ApiServiceI.
type ApiService struct {
	// httpClient  *http.Client
	// Repo        repository.AddressRepositoryI
	// RedisClient *cache.RedisClient
	// BtcClient   *config.BtcConfig
	// EvmClient   *ethprovider.EvmService
	// db          *gorm.DB
	// MempoolBase string
	// Sleep       func(time.Duration)
	// CallBTCRPC  func(client *http.Client, url, token, method string, params []interface{}, out interface{}) error
}

// NewApiService creates a new ApiService instance.
func NewApiService() ApiServiceI {
	return &ApiService{}
}
