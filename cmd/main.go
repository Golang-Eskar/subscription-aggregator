package main

import (
	"fmt"
	"github.com/Golang-Eskar/subscription-aggregator/internal/database"
	"log"
)

func main() {
	fmt.Println("Запуск")

	err := database.Init()
	if err != nil {
		log.Fatal("Ошибка базы данных:", err)
	}
}
