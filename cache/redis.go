package cache

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client
var Ctx = context.Background()

func ConnectRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),     // e.g. "127.0.0.1:6379"
		Password: os.Getenv("REDIS_PASSWORD"), // if set in redis.conf
		DB:       0,
	})

	_, err := Rdb.Ping(Ctx).Result()
	if err != nil {
		panic("❌ Failed to connect to Redis: " + err.Error())
	}

	fmt.Println("✅ Connected to Redis")
}
