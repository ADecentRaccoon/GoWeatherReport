package utils

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	// "github.com/ADecentRaccoon/GoWeatherReport/"
)

var redisContext = context.Background()

func SaveResponse(location string, dataToSave any){
	var RedisClient = redis.NewClient(&redis.Options{Addr: fmt.Sprintf("%v:%v", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))})
	redisStatus := RedisClient.Set(redisContext, location, dataToSave, 5 * time.Minute)
	if redisStatus.Err() != nil{
		fmt.Println("redisError: ", redisStatus.Err().Error())
	}
}

func GetResponse(location string) (bool, []byte) {
	var RedisClient = redis.NewClient(&redis.Options{Addr: fmt.Sprintf("%v:%v", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))})
	ans, err := RedisClient.Get(redisContext, location).Result()
	if err == redis.Nil {
		return false, nil
	} else if err != nil {
		fmt.Println(err)
		return false, nil
	}


	return true, []byte(ans)
}

