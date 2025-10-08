package ethprovider

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EvmServiceI interface {
	GetBalance(address string) (float64, error)
}

type EvmService struct {
	ethereumClient *ethclient.Client
}

func NewEvmService(
	EthRpcURL string,
) (*EvmService, error) {
	ethereumClient, err := ethclient.Dial(EthRpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum client: %w", err)
	}

	return &EvmService{
		ethereumClient: ethereumClient,
	}, nil
}

// GetBalance fetches the balance of an address in ETH (float64)
func (s *EvmService) GetBalance(address string) (float64, error) {
	account := common.HexToAddress(address)

	// nil = latest block
	balanceWei, err := s.ethereumClient.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch balance: %w", err)
	}

	// Convert from wei (big.Int) → ETH (float64)
	ethValue := new(big.Float).Quo(
		new(big.Float).SetInt(balanceWei),
		big.NewFloat(1e18),
	)

	result, _ := ethValue.Float64()
	return result, nil
}

func (s *EvmService) GetERC20Balance(address string, contract string) (float64, error) {
	parsedABI, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return 0, fmt.Errorf("failed to parse ABI: %w", err)
	}

	contractAddr := common.HexToAddress(contract)
	caller := bind.NewBoundContract(contractAddr, parsedABI, s.ethereumClient, s.ethereumClient, s.ethereumClient)

	account := common.HexToAddress(address)

	// 1. Get decimals
	var decimals uint8
	decimalsResult := []interface{}{&decimals}
	err = caller.Call(&bind.CallOpts{}, &decimalsResult, "decimals")
	if err != nil {
		// fallback default (most tokens use 18 decimals)
		decimals = 8
	}

	// 2. Get balance
	balance := new(big.Int)
	balanceResult := []interface{}{&balance}
	err = caller.Call(&bind.CallOpts{}, &balanceResult, "balanceOf", account)
	if err != nil {
		return 0, fmt.Errorf("failed to get token balance: %w", err)
	}

	// Convert using decimals - handle precision properly
	if balance.Sign() == 0 {
		return 0, nil
	}

	// Create the divisor (10^decimals)
	divisor := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)

	// Convert to float64 with proper precision
	balanceFloat := new(big.Float).SetInt(balance)
	divisorFloat := new(big.Float).SetInt(divisor)

	tokenValue := new(big.Float).Quo(balanceFloat, divisorFloat)

	result, accuracy := tokenValue.Float64()
	if accuracy != big.Exact && accuracy != big.Below {
		// Log warning if precision is lost, but continue
		// In production, you might want to return an error or use string representation
	}

	return result, nil
}

// GetTransactionReceipt fetches the receipt of a transaction by hash
func (s *EvmService) GetTransactionReceipt(txHash string) (*types.Receipt, error) {
	hash := common.HexToHash(txHash)
	receipt, err := s.ethereumClient.TransactionReceipt(context.Background(), hash)
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction receipt: %w", err)
	}
	return receipt, nil
}

// GetLatestBlockNumber fetches the latest block number from Ethereum
func (s *EvmService) GetLatestBlockNumber() (*big.Int, error) {
	header, err := s.ethereumClient.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest block header: %w", err)
	}
	return header.Number, nil
}

func (s *EvmService) GetTransactionByHash(ctx context.Context, txHash string) (*types.Transaction, error) {
	hash := common.HexToHash(txHash)
	tx, _, err := s.ethereumClient.TransactionByHash(ctx, hash)
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction by hash: %w", err)
	}
	return tx, nil
}
