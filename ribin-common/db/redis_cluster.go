package db

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
)

const (
	DEFAULT_LOCK_TOKEN = "Matrix"
)

type RedisCluster struct {
	Client  *redis.ClusterClient
	address string
	rs      *redsync.Redsync
}

func (client *RedisCluster) NewScript(src string) *redis.Script {
	return redis.NewScript(src)
}

func (client *RedisCluster) Eval(ctx context.Context, script string, keys []string, args ...interface{}) *redis.Cmd {
	return client.Client.Eval(ctx, script, keys, args)
}

func (client *RedisCluster) EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) *redis.Cmd {
	return client.Client.EvalSha(ctx, sha1, keys, args)
}

func (client *RedisCluster) ScriptExists(ctx context.Context, hashes ...string) *redis.BoolSliceCmd {
	return client.Client.ScriptExists(ctx, hashes...)
}

func (client *RedisCluster) ScriptLoad(ctx context.Context, script string) *redis.StringCmd {
	return client.Client.ScriptLoad(ctx, script)
}
func (client *RedisCluster) LPush(ctx context.Context, key string, value string) (int64, error) {
	cmd := client.Client.LPush(ctx, key, value)
	return cmd.Val(), cmd.Err()
}

func (client *RedisCluster) LRem(ctx context.Context, key string, count int64, value string) (int64, error) {
	cmd := client.Client.LRem(ctx, key, count, value)
	return cmd.Val(), cmd.Err()
}

func (client *RedisCluster) LLen(ctx context.Context, key string) (int64, error) {
	cmd := client.Client.LLen(ctx, key)
	return cmd.Val(), cmd.Err()
}

func (client *RedisCluster) Get(ctx context.Context, key string) (string, error) {
	cmd := client.Client.Get(ctx, key)
	return cmd.Val(), cmd.Err()
}

func (client *RedisCluster) Set(ctx context.Context, key, value string, expiration time.Duration) error {
	cmd := client.Client.Set(ctx, key, value, expiration)
	return cmd.Err()
}

func (client *RedisCluster) IncrBy(ctx context.Context, key string, value int64) (int64, error) {
	cmd := client.Client.IncrBy(ctx, key, value)
	return cmd.Val(), cmd.Err()
}

func (client *RedisCluster) DecrBy(ctx context.Context, key string, value int64) (int64, error) {
	cmd := client.Client.DecrBy(ctx, key, value)
	return cmd.Val(), cmd.Err()
}

func (client *RedisCluster) ZAdd(ctx context.Context, key, value string, score float64) error {
	return client.Client.ZAdd(ctx, key, &redis.Z{Member: value, Score: score}).Err()
}

func (client *RedisCluster) ZRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	zRange := client.Client.ZRange(ctx, key, start, stop)
	return zRange.Val(), zRange.Err()
}

func (client *RedisCluster) ZRangeWithScores(ctx context.Context, key string, start, stop int64) ([]redis.Z, error) {
	cmd := client.Client.ZRangeWithScores(ctx, key, start, stop)
	return cmd.Val(), cmd.Err()
}

func (client *RedisCluster) ZRevRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	cmd := client.Client.ZRevRange(ctx, key, start, stop)
	return cmd.Val(), cmd.Err()
}

func (client *RedisCluster) ZIncrBy(ctx context.Context, key string, increment float64, member string) error {
	return client.Client.ZIncrBy(ctx, key, increment, member).Err()
}

func (client *RedisCluster) ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) ([]redis.Z, error) {
	cmd := client.Client.ZRevRangeWithScores(ctx, key, start, stop)
	return cmd.Val(), cmd.Err()
}

func (client *RedisCluster) ZCard(ctx context.Context, key string) (int, error) {
	cmd := client.Client.ZCard(ctx, key)
	return int(cmd.Val()), cmd.Err()
}

func (client *RedisCluster) HGet(ctx context.Context, key string, field string) (string, error) {
	cmd := client.Client.HGet(ctx, key, field)
	if cmd.Err() != nil {
		return "", cmd.Err()
	}
	return cmd.Val(), nil
}

func (client *RedisCluster) HSet(ctx context.Context, key string, field, value string) error {
	cmd := client.Client.HSet(ctx, key, field, value)
	return cmd.Err()
}

func (client *RedisCluster) Del(ctx context.Context, key string) error {
	return client.Client.Del(ctx, key).Err()
}

func (client *RedisCluster) ExpireAt(ctx context.Context, key string, expiration time.Duration) error {
	return client.Client.Expire(ctx, key, expiration).Err()
}

func (client *RedisCluster) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	cmd := client.Client.HGetAll(ctx, key)
	return cmd.Val(), cmd.Err()
}

func (client *RedisCluster) ZRem(ctx context.Context, key string, fields ...string) error {
	return client.Client.ZRem(ctx, key, fields).Err()
}

func (client *RedisCluster) HRem(ctx context.Context, key string, fields ...string) error {
	return client.Client.HDel(ctx, key, fields...).Err()
}

func (client *RedisCluster) ZCount(ctx context.Context, key string, min, max int64) (int64, error) {
	return client.Client.ZCount(ctx, key, fmt.Sprintf("%d", min), fmt.Sprintf("%d", max)).Result()
}

// Lock
// @Description: redis distributed lock
// @param key mutex name
// @param ttl set the expiry of a mutex to the given value.
// @param retryDelay set the amount of time to wait between retries.
// @param tries set the number of times lock acquire is attempted
// @return error
func (client *RedisCluster) Lock(ctx context.Context, key string, acquireTimeOut, lockTimeOut time.Duration) error {
	var acquireDuration = acquireTimeOut
	key = "lock:" + key
	for acquireDuration > 0 {
		ok, err := client.Client.SetNX(ctx, key, DEFAULT_LOCK_TOKEN, lockTimeOut).Result()
		if err != nil {
			return err
		}
		if ok {
			return nil
		}
		time.Sleep(time.Millisecond * 100)
		acquireDuration -= time.Millisecond * 100
	}
	return fmt.Errorf("get redis lock timeout: %v\r\n", acquireTimeOut)
}

func (client *RedisCluster) UnLock(ctx context.Context, key string) error {
	const script = `if redis.call('GET', KEYS[1]) == ARGV[1] then return redis.call('DEL', KEYS[1]) else return 0 end`
	_, err := client.Client.Eval(ctx, script, []string{key}, DEFAULT_LOCK_TOKEN).Result()
	if err != nil {
		return err
	}
	return nil
}

func (client *RedisCluster) Pipeline() redis.Pipeliner {
	return client.Client.Pipeline()
}
