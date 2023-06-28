package db

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
)

type RedisCluster struct {
	Client  *redis.ClusterClient
	address string
	rs      *redsync.Redsync
}

func (client *RedisCluster) Get(ctx context.Context, key string) (string, error) {
	val, err := client.Client.Get(ctx, key).Result()
	return val, err
}
func (client *RedisCluster) Set(ctx context.Context, key, value string, expiration time.Duration) error {
	_, err := client.Client.Set(ctx, key, value, expiration).Result()
	return err
}
