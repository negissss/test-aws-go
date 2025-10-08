package cmcprovider

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type CMCProvider interface {
	GetPrice(symbols []string, convert string) (map[string]float64, error)
}

type cmcProvider struct {
	apiKey string
	client *http.Client
}

func NewCMCProvider(apiKey string) CMCProvider {
	return &cmcProvider{
		apiKey: apiKey,
		client: &http.Client{Timeout: 10 * time.Second},
	}
}

type cmcResponse struct {
	Data map[string]struct {
		Quote map[string]struct {
			Price float64 `json:"price"`
		} `json:"quote"`
	} `json:"data"`
}

// GetPrice fetches crypto prices for given symbols from CMC
func (c *cmcProvider) GetPrice(symbols []string, convert string) (map[string]float64, error) {
	// Join symbols as comma-separated string
	symbolParam := strings.Join(symbols, ",")

	url := fmt.Sprintf(
		"https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest?symbol=%s&convert=%s",
		symbolParam, convert,
	)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-CMC_PRO_API_KEY", c.apiKey)
	req.Header.Set("Accepts", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var body []byte
		body, _ = io.ReadAll(resp.Body)
		return nil, fmt.Errorf("CMC API error: %s", string(body))
	}

	var result cmcResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	prices := make(map[string]float64)
	for sym, data := range result.Data {
		if quote, ok := data.Quote[convert]; ok {
			prices[sym] = quote.Price
		}
	}

	return prices, nil
}
