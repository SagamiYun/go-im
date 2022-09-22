package test

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"testing"
	"time"
)

var ctx = context.Background()

func TestSet(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "key", "value", time.Second*30).Err()
	if err != nil {
		panic(err)
	}
}

func TestGet(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	r, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(r)
}
