package rediska

import (
	"context"
	"stalcraftbot/internal/logs"

	"github.com/redis/go-redis/v9"
)

var (
	client = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	ctx = context.Background()
)

func SaveLastEmissionDataToRedis(data string) error {

	err := client.Set(ctx, "data", data, 0).Err()
	if err != nil {
		return err
	}
	logs.Logger.Info().Msg("Save data to Redis - OK")
	return nil
}
func LoadEmDataFromRedis() (string, error) {

	val, err := client.Get(ctx, "data").Result()
	if err != nil {
		return "", err
	}
	logs.Logger.Debug().Msg("Load Data from Redis is OK")
	return val, nil
}
