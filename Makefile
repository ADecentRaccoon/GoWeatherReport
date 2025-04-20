APP_NAME=GoWeatherReport

default: build

build:
	go build -o $(APP_NAME) cmd/main.go
