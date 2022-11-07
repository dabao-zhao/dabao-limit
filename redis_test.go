package limit

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type testRedis struct {
	client *redis.Client
}

func newTestRedis(addr string) *testRedis {
	return &testRedis{client: redis.NewClient(&redis.Options{
		Addr: addr,
	})}
}

func (r *testRedis) Eval(ctx context.Context, script string, keys []string, args ...interface{}) (val interface{}, err error) {
	return r.client.Eval(ctx, script, keys, args...).Result()
}

func (r *testRedis) Ping() bool {
	return true
}

func (r *testRedis) Close() {
	_ = r.client.Close()
}
