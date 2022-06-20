package main

import (
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("no .env file found")
	}
}

func main() {
	//_, cancel := start()
	//defer shutdown(cancel)
	//servicemanager.WaitShutdown()
}
