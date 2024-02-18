package store

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type StorageService struct {
	redisClient *redis.Client
}

var storageService = &StorageService{}
var ctx = context.Background()

func ConnectToServer(redisAddress string, redisPort string, redisPassword string, dbNo int) *StorageService {

	rClient := redis.NewClient(&redis.Options{
		Addr:     redisAddress + ":" + redisPort,
		Password: redisPassword,
		DB:       dbNo,
	})

	pong, err := rClient.Ping(ctx).Result()
	if err != nil {
		panic("Failed to connect to Redis server. Error: " + err.Error() + "\n")
	}
	fmt.Println(pong)

	storageService.redisClient = rClient
	return storageService
}

func SaveMapping(shortUrl string, originalUrl string) {
	err := storageService.redisClient.Set(ctx, shortUrl, originalUrl, 0).Err()
	if err != nil {
		panic(fmt.Sprintf("fail cannot save: %v, shortUrl: %s, originalUrl: %s\n", err, shortUrl, originalUrl))
	}
}

func RetrieveLongUrl(shortUrl string) string {
	result, err := storageService.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("cannot get long url: %v, shortUrl: %s\n", err, shortUrl))
	}
	return result
}
