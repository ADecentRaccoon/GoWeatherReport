package main

import(
	"github.com/ADecentRaccoon/GoWeatherReport/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func main() {
	godotenv.Load()
	r := gin.Default()
	weatherGroup := r.Group("/weather")
	routers.WeatherRouter(weatherGroup)

	r.Run("localhost:8083")
}