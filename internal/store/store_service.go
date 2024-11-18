package store

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type StorageService struct {
	redisClient *redis.Client
}

var (
	storeService = &StorageService{}
)

const CacheDuration = 6 * time.Hour

func InitStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}

	fmt.Printf("Connected to Redis: %v\n", pong)
	storeService.redisClient = redisClient
	return storeService
}

func SaveUrlMapping(shortUrl string, originalUrl string, userId string) {
	err := storeService.redisClient.Set(shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Error saving to Redis: %v", err))
	}
}

func RetrieveOriginalUrl(shortUrl string) string {
	res, err := storeService.redisClient.Get(shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Error retrieving from Redis: %v", err))
	}

	return res
}
