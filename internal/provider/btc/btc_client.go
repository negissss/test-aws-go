package btcprovider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type rpcRequest struct {
	JSONRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      int           `json:"id"`
}

// CallRPC sends a JSON-RPC request to the given URL with the API token and unmarshals the result into 'result'.
func CallRPC(
	httpClient *http.Client,
	rpcURL string,
	apiToken string,
	method string,
	params []interface{},
	result interface{},
) error {
	// fmt.Println("CallRPC params", httpClient, rpcURL, apiToken)
	reqBody := rpcRequest{
		JSONRPC: "2.0",
		Method:  method,
		Params:  params,
		ID:      1,
	}
	data, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON RPC request: %w", err)
	}

	req, err := http.NewRequest("POST", rpcURL, bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if apiToken != "" {
		req.Header.Set("x-api-key", apiToken)
	}
	// req.Header.Set("x-api-key", apiToken)

	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("RPC HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read RPC response body: %w", err)
	}

	var respWrapper struct {
		Result json.RawMessage  `json:"result"`
		Error  *json.RawMessage `json:"error"`
		ID     int              `json:"id"`
	}

	if err := json.Unmarshal(respBody, &respWrapper); err != nil {
		return fmt.Errorf("failed to unmarshal RPC response: %w", err)
	}

	if respWrapper.Error != nil {
		return fmt.Errorf("RPC error: %s", string(*respWrapper.Error))
	}

	if result != nil {
		if err := json.Unmarshal(respWrapper.Result, &result); err != nil {
			return fmt.Errorf("failed to unmarshal RPC result: %w", err)
		}
	}

	return nil
}
