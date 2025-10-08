package cache

import (
	"context"
	"fmt"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(addr string, password string) *RedisClient {
	fmt.Println("addr::::", addr, password)
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	return &RedisClient{
		client: rdb,
	}
}

func (r *RedisClient) SetRedisKey(ctx context.Context, keyName string, blockNumber int) error {
	return r.client.Set(ctx, keyName, blockNumber, 0).Err()
}

func (r *RedisClient) GetRedisKey(ctx context.Context, keyName string) (uint64, error) {
	val, err := r.client.Get(ctx, keyName).Result()
	if err == redis.Nil {
		return 0, nil
	} else if err != nil {
		fmt.Println("redis err", err)
		return 0, err
	}
	return strconv.ParseUint(val, 10, 64)
}

func (r *RedisClient) RemovePendingTxnHash(ctx context.Context, chain, txHash string) error {
	key := "pending_tx:" + chain
	return r.client.SRem(ctx, key, txHash).Err()
}

func (r *RedisClient) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return r.client.HGetAll(ctx, key).Result()
}

func (r *RedisClient) HSetField(ctx context.Context, key, field, value string) error {
	return r.client.HSet(ctx, key, field, value).Err()
}

func (r *RedisClient) HDel(ctx context.Context, key string, fields ...string) (int64, error) {
	return r.client.HDel(ctx, key, fields...).Result()
}
func (r *RedisClient) SAdd(ctx context.Context, key string, fields ...string) (int64, error) {
	return r.client.SAdd(ctx, key, fields).Result()
}
