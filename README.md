# Go weather API

## Description

This is a simple API built in Go that fetches weather data from a third-party API (https://www.visualcrossing.com/weather-api/), caches the response using Redis, and returns the cached data if necessary.


## SETUP
1. ``` git clone github.com/ADecentRaccoon/GoWeatherReport``` Clone the repository
2. ```touch .env``` create .env file or ```mv .env.example .env``` to edit example
2. ```nano .env``` Edit .env file 
3. Insert key from https://www.visualcrossing.com/weather-api/ after APIKEY=
4. ```make``` Build and start app. Use ```make docker``` to start through docker. or ```make build && make run``` to start app native.

project https://roadmap.sh/projects/weather-api-wrapper-service