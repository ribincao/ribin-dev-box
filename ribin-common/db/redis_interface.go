package db

import (
	"context"
	"time"
)

type IRedis interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key, value string, expiration time.Duration) error
}
