package utils

import (
	"github.com/joho/godotenv"
)

func loadEnv(){
	if err := godotenv.Load(); err != nil{
		panic(err)
	}

}