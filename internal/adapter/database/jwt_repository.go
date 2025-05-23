package database

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func (r *JwtRedisAdapter) SetKey(ctx context.Context, key, value string, ttl time.Duration) error {
	err := r.rdb.Set(ctx, key, value, ttl).Err()
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("can't set key: %v", err)
	}
	return nil
}

func (r *JwtRedisAdapter) GetKey(ctx context.Context, key string) (string, error) {
	val, err := r.rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("key does not exist")
	} else if err != nil {
		return "", fmt.Errorf("can't get key: %v", err)
	}
	return val, nil
}
