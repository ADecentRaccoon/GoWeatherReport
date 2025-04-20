package main

import (
	"fmt"
	"os"

	"github.com/ADecentRaccoon/GoWeatherReport/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	godotenv.Load()
	r := gin.Default()
	weatherGroup := r.Group("/weather")
	routers.WeatherRouter(weatherGroup)
	HOST := os.Getenv("RUNNING_AT_HOST")
	PORT := os.Getenv("RUNNING_AT_PORT")
	fmt.Println("Running at", HOST, ":", PORT)
	r.Run(fmt.Sprintf("%v:%v", HOST, PORT))
}