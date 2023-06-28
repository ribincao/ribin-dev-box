package db

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type IRedis interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key, value string, expiration time.Duration) error
	HSet(ctx context.Context, key string, field, value string) error
	HGet(ctx context.Context, key string, field string) (string, error)
	HGetAll(ctx context.Context, key string) (map[string]string, error)
	HRem(ctx context.Context, key string, fields ...string) error
	ZAdd(ctx context.Context, key, value string, score float64) error
	ZRange(ctx context.Context, key string, start, stop int64) ([]string, error)
	ZRem(ctx context.Context, key string, fields ...string) error
	ZRevRange(ctx context.Context, key string, start, stop int64) ([]string, error)
	ZCard(ctx context.Context, key string) (int, error)
	ZIncrBy(ctx context.Context, key string, increment float64, member string) error
	ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) ([]redis.Z, error)
	ZRangeWithScores(ctx context.Context, key string, start, stop int64) ([]redis.Z, error)
	ZCount(ctx context.Context, key string, min, max int64) (int64, error)
	LPush(ctx context.Context, key string, value string) (int64, error)
	LRem(ctx context.Context, key string, count int64, value string) (int64, error)
	LLen(ctx context.Context, key string) (int64, error)
	Del(ctx context.Context, key string) error
	ExpireAt(ctx context.Context, key string, expiration time.Duration) error
	Lock(ctx context.Context, key string, acquireTimeout time.Duration, lockTimeout time.Duration) error
	UnLock(ctx context.Context, key string) error
	IncrBy(ctx context.Context, key string, value int64) (int64, error)
	DecrBy(ctx context.Context, key string, value int64) (int64, error)
	Pipeline() redis.Pipeliner
	Eval(ctx context.Context, script string, keys []string, args ...interface{}) *redis.Cmd
	EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) *redis.Cmd
	ScriptExists(ctx context.Context, hashes ...string) *redis.BoolSliceCmd
	ScriptLoad(ctx context.Context, script string) *redis.StringCmd
	NewScript(src string) *redis.Script
}
