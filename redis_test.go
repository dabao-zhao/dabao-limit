package limit

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type TestRedis struct {
	client *redis.Client
}

func NewTestRedis(addr string) *TestRedis {
	return &TestRedis{client: redis.NewClient(&redis.Options{
		Addr: addr,
	})}
}

func (r *TestRedis) Eval(ctx context.Context, script string, keys []string, args ...interface{}) (val interface{}, err error) {
	return r.client.Eval(ctx, script, keys, args...).Result()
}

func (r *TestRedis) Ping() bool {
	return true
}

func (r *TestRedis) Close() {
	_ = r.client.Close()
}
