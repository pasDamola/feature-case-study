package initializers

import (
	"os"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func ConnectToRedis() {
	redisUrl := os.Getenv("REDIS_URL")
	RedisClient = redis.NewClient(&redis.Options{
		Addr:redisUrl,
		Password: "",
        DB:       0,
	 })

}