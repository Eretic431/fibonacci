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

func NewFibonacciRepository(pool *redis.Pool) (*FibonacciRepository, error) {
	fr := &FibonacciRepository{Pool: pool}
	ctx := context.Background()
	return fr, fr.initialSetup(ctx)
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

func (fr *FibonacciRepository) GetInterval(ctx context.Context, from, to int) ([]int64, error) {
	conn, err := fr.Pool.GetContext(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()


	reply, err := redis.Int64s(conn.Do("ZRANGE", "fibonacciNumbers", from-1, to-1))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (fr *FibonacciRepository) initialSetup(ctx context.Context) error {
	conn, err := fr.Pool.GetContext(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	err = conn.Send("ZADD", "fibonacciNumbers", "NX", 1, 0)
	if err != nil {
		return err
	}

	err = conn.Send("ZADD", "fibonacciNumbers", "NX", 2, "01")

	return err
}
