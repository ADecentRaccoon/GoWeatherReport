package utils

import (
	"fmt"
	"net/http"
	"os"
	"encoding/json"
	"io"
)

func FetchWeather(location string, date1 string, date2 string){
	url := fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s/%s/%s?key=%s", location, date1, date2, os.Getenv("APIKEY"))
	res, err := http.Get(url)
	if err != nil{
		panic(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil{
		fmt.Println("Reading error: ", err)
		return
	}
	var data map[string]any

	err = json.Unmarshal(body, &data)
	if err != nil{
		fmt.Println("Parse error: ", err)
	}


	fmt.Println(data)
}
