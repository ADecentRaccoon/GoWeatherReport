package utils

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
	// "github.com/ADecentRaccoon/GoWeatherReport/"
)

var redisContext = context.Background()

func SaveResponse(location string, dataToSave []byte){
	redisClient := redis.NewClient(&redis.Options{Addr: fmt.Sprintf("%v:%v", os.Getenv("DB_HOST"), os.Getenv("REDIS_PORT"))})
	redisClient.Set(redisContext, location, dataToSave)
}