package redis

import (
	"context"
	"github.com/Eretic431/fibonacci/config"
	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
)

func NewRedisPool(c *config.Config, log *zap.SugaredLogger) (*redis.Pool, func()) {
	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", c.RedisCfg.Url)
		},
		DialContext: func(ctx context.Context) (redis.Conn, error) {
			return redis.DialContext(ctx, "tcp", c.RedisCfg.Url)
		},
	}
	return pool, func() {
		if err := pool.Close(); err != nil {
			log.Errorw("Failed closing redis pool", "err", err)
		}
	}
}
