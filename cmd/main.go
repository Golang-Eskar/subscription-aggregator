package main

import (
	"fmt"
	"log"

	"github.com/Golang-Eskar/subscription-aggregator/internal/database"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Запуск")

	err1 := godotenv.Load()
	if err1 != nil {
		return
	}
	err := database.Init()
	if err != nil {
		log.Fatal("Ошибка базы данных:", err)
	}
}
