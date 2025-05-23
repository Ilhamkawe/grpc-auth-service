package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseAdapter struct {
	db *gorm.DB
}

func NewDatabaseAdapter(conn *sql.DB) (*DatabaseAdapter, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: conn,
	}), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("Can't connect database (gorm) : %v", err)
	}
	return &DatabaseAdapter{
		db: db,
	}, nil
}

type JwtRedisAdapter struct {
	rdb *redis.Client
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
		rdb: rdb,
	}, nil
}
