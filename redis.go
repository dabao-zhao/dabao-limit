package limit

import "context"

const RedisNil = RedisError("redis: nil")

type RedisError string

func (e RedisError) Error() string { return string(e) }

func (RedisError) RedisError() {}

type Redis interface {
	Eval(ctx context.Context, script string, keys []string, args ...interface{}) (val interface{}, err error)
	Ping() bool
}
