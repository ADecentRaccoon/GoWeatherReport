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
	url := fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s/%s/%s?key=%s", location, date1, date2, os.Getenv("APIKEY"))
	data := make(map[string]any)
	res, err := http.Get(url)
	if err != nil{
		panic(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil{
		fmt.Println("Reading error: ", err)
		data["400"] = "Invailed responce"
		return data
	}
	err = json.Unmarshal(body, &data)
	if err != nil{
		fmt.Println("Parse error: ", err)
	}


	return data
}
