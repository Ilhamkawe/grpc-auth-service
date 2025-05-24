package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type JwtRedisAdapter struct {
	RDB *redis.Client
}

func NewJwtRedisAdapter(addr, username, password string, db int) (*JwtRedisAdapter, error) {

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: username,
		Password: password,
		DB:       db,
	})

	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("can't connect to Redis: %v", err)
	}

	return &JwtRedisAdapter{
		RDB: rdb,
	}, nil
}
