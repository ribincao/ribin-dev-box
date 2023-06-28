package db

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	Client *redis.Client
}

func (client *RedisClient) NewScript(src string) *redis.Script {
	return redis.NewScript(src)
}

func (client *RedisClient) Eval(ctx context.Context, script string, keys []string, args ...interface{}) *redis.Cmd {
	return client.Client.Eval(ctx, script, keys, args)
}

func (client *RedisClient) EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) *redis.Cmd {
	return client.Client.EvalSha(ctx, sha1, keys, args)
}

func (client *RedisClient) ScriptExists(ctx context.Context, hashes ...string) *redis.BoolSliceCmd {
	return client.Client.ScriptExists(ctx, hashes...)
}

func (client *RedisClient) ScriptLoad(ctx context.Context, script string) *redis.StringCmd {
	return client.Client.ScriptLoad(ctx, script)
}

func (client *RedisClient) LPush(ctx context.Context, key string, value string) (int64, error) {
	cmd := client.Client.LPush(ctx, key, value)
	return cmd.Val(), cmd.Err()
}

func (client *RedisClient) LRem(ctx context.Context, key string, count int64, value string) (int64, error) {
	cmd := client.Client.LRem(ctx, key, count, value)
	return cmd.Val(), cmd.Err()
}

func (client *RedisClient) LLen(ctx context.Context, key string) (int64, error) {
	cmd := client.Client.LLen(ctx, key)
	return cmd.Val(), cmd.Err()
}

func (client *RedisClient) Get(ctx context.Context, key string) (string, error) {
	cmd := client.Client.Get(ctx, key)
	return cmd.Val(), cmd.Err()
}

func (client *RedisClient) Set(ctx context.Context, key, value string, expiration time.Duration) error {
	return client.Client.Set(ctx, key, value, expiration).Err()
}

func (client *RedisClient) IncrBy(ctx context.Context, key string, value int64) (int64, error) {
	cmd := client.Client.IncrBy(ctx, key, value)
	return cmd.Val(), cmd.Err()
}

func (client *RedisClient) DecrBy(ctx context.Context, key string, value int64) (int64, error) {
	cmd := client.Client.DecrBy(ctx, key, value)
	return cmd.Val(), cmd.Err()
}

func (client *RedisClient) ZAdd(ctx context.Context, key, value string, score float64) error {
	return client.Client.ZAdd(ctx, key, &redis.Z{Member: value, Score: score}).Err()
}

func (client *RedisClient) ZRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	cmd := client.Client.ZRange(ctx, key, start, stop)
	return cmd.Val(), cmd.Err()
}

func (client *RedisClient) ZRangeWithScores(ctx context.Context, key string, start, stop int64) ([]redis.Z, error) {
	cmd := client.Client.ZRangeWithScores(ctx, key, start, stop)
	return cmd.Val(), cmd.Err()
}

func (client *RedisClient) ZRem(ctx context.Context, key string, fields ...string) error {
	return client.Client.ZRem(ctx, key, fields).Err()
}

func (client *RedisClient) ZRevRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	cmd := client.Client.ZRevRange(ctx, key, start, stop)
	return cmd.Val(), cmd.Err()
}

func (client *RedisClient) ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) ([]redis.Z, error) {
	cmd := client.Client.ZRevRangeWithScores(ctx, key, start, stop)
	return cmd.Val(), cmd.Err()
}

func (client *RedisClient) ZIncrBy(ctx context.Context, key string, increment float64, member string) error {
	cmd := client.Client.ZIncrBy(ctx, key, increment, member)
	return cmd.Err()
}

func (client *RedisClient) ZCard(ctx context.Context, key string) (int, error) {
	cmd := client.Client.ZCard(ctx, key)
	return int(cmd.Val()), cmd.Err()
}

func (client *RedisClient) HSet(ctx context.Context, key string, field, value string) error {
	return client.Client.HSet(ctx, key, field, value).Err()
}

func (client *RedisClient) HGet(ctx context.Context, key string, field string) (string, error) {
	cmd := client.Client.HGet(ctx, key, field)
	return cmd.Val(), cmd.Err()
}

func (client *RedisClient) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	cmd := client.Client.HGetAll(ctx, key)
	return cmd.Val(), cmd.Err()
}

func (client *RedisClient) HRem(ctx context.Context, key string, fields ...string) error {
	return client.Client.HDel(ctx, key, fields...).Err()
}

func (client *RedisClient) Del(ctx context.Context, key string) error {
	return client.Client.Del(ctx, key).Err()
}

func (client *RedisClient) ExpireAt(ctx context.Context, key string, expiration time.Duration) error {
	return client.Client.Expire(ctx, key, expiration).Err()
}

func (client *RedisClient) Lock(ctx context.Context, key string, acquireTimeout time.Duration, lockTimeout time.Duration) error {
	return nil
}

func (client *RedisClient) UnLock(ctx context.Context, key string) error {
	return nil
}

func (client *RedisClient) ZCount(ctx context.Context, key string, min, max int64) (int64, error) {
	return client.Client.ZCount(ctx, key, fmt.Sprintf("%d", min), fmt.Sprintf("%d", max)).Result()
}

func (client *RedisClient) Pipeline() redis.Pipeliner {
	return client.Client.Pipeline()
}
