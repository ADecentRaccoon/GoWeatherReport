package routers

import (
	"fmt"
	"net/http"

	"github.com/ADecentRaccoon/GoWeatherReport/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Query struct{
	Date1 string `binding:"required" validate:"datetime=2006-01-02"`
	Date2 string `binding:"required" validate:"datetime=2006-01-02"`
	Location string `binding:"gte=2,lte=20"`
}


func WeatherRouter(r *gin.RouterGroup){
	r.POST("/get", func(ctx *gin.Context) {
		var request Query
	if validateErr := ctx.ShouldBindJSON(&request); validateErr != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"message": validateErr})
		return
	}

	validate := validator.New()
	if valError := validate.Struct(&request); valError != nil{
		fmt.Println(valError)
		return
	}
	fmt.Println()
	if exist, data := utils.GetResponse(request.Location); exist{ // Getting value from Redis
		ctx.JSON(203, data)
		println("from Redis")
		return
	}
	weather := utils.FetchWeather(request.Location, request.Date1, request.Date2)
	utils.SaveResponse(request.Location, weather)
	ctx.JSON(200, weather)
	})
}