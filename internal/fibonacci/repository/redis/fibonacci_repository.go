package redis

import (
	"context"
	"errors"
	"github.com/Eretic431/fibonacci/internal/models"
	"github.com/gomodule/redigo/redis"
)

type FibonacciRepository struct {
	Pool *redis.Pool
}

func NewFibonacciRepository(pool *redis.Pool) *FibonacciRepository {
	return &FibonacciRepository{Pool: pool}
}

func (fr *FibonacciRepository) Get(ctx context.Context, key int) (int64, error) {
	conn, err := fr.Pool.GetContext(ctx)
	if err != nil {
		return -1, err
	}
	defer conn.Close()

	// get value with a key fibonacciNumbers from sorted set in range key:key (returns value with score key)
	value, err := redis.Int64(conn.Do("ZRANGE", "fibonacciNumbers", key-1, key-1))
	if err != nil {
		if errors.Is(err, redis.ErrNil) {
			return 0, models.ErrNoRecord
		}
		return -1, err
	}

	return value, nil
}

func (fr *FibonacciRepository) Set(ctx context.Context, key int, value int64) error {
	conn, err := fr.Pool.GetContext(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	// add value Value with key fibonacciNumbers to a sorted set with score Key(represents index number) if not exists
	err = conn.Send("ZADD", "fibonacciNumbers", "NX", key, value)

	return err
}

func (fr *FibonacciRepository) GetLastTwoNumbers(ctx context.Context) ([]int64, error) {
	conn, err := fr.Pool.GetContext(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	output := make([]int64, 4)

	// get last two numbers with index
	reply, err := redis.Values(conn.Do("ZRANGE", "fibonacciNumbers", "-2", "-1", "WITHSCORES"))
	if err != nil {
		return nil, err
	}
	if _, err := redis.Scan(reply, &output[0], &output[1], &output[2], &output[3]); err != nil {
		return nil, err
	}

	return output, nil
}
