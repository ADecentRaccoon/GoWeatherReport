package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	// "github.com/ADecentRaccoon/GoWeatherReport/"
)

var RedisClient = redis.NewClient(&redis.Options{Addr: fmt.Sprintf("%v:%v", os.Getenv("DB_HOST"), os.Getenv("REDIS_PORT"))})
var redisContext = context.Background()

func SaveResponse(location string, dataToSave any){
	RedisClient.Set(redisContext, location, dataToSave, 3600 * time.Second)
}

func GetResponse(location string) (bool, any) {
	ans, err := RedisClient.Get(redisContext, location).Result()
	if err == redis.Nil {
		return false, nil
	} else if err != nil {
		return false, nil
	}

	var data any
	json.Unmarshal([]byte(ans), &data) // Десериализация JSON
	return true, data
}

