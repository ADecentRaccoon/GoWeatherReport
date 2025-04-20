package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	// "github.com/redis/go-redis/v9"
)

func FetchWeather(location string, date1 string, date2 string) map[string]any {
	data := make(map[string]any)
	if exist, resived := GetResponse(location); exist{
		parseError := json.Unmarshal([]byte(resived), &data)
		if parseError != nil{
			println("parseError (1): ",parseError.Error())
			return nil
		}
		data["from cache"] = true
		return data
	}else{
	url := fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s/%s/%s?key=%s", location, date1, date2, os.Getenv("APIKEY"))
	res, err := http.Get(url)
	if err != nil{
		panic(err)
	}
	body, Readerr := io.ReadAll(res.Body)
	SaveResponse(location, body)
	if Readerr != nil{
		fmt.Println("Reading error: ", err)
		data["400"] = "Invailed responce"
		return data
	}
	parseError := json.Unmarshal(body, &data)
		if parseError != nil{
			println("parseError (2): ",parseError)
		}
		data["from cache"] = false
		return data
	}
}
