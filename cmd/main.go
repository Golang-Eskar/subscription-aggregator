package main

import (
	"fmt"
	"log"

	"github.com/Golang-Eskar/subscription-aggregator/internal/database"
)

func main() {
	fmt.Println("Запуск")

	err := database.Init()
	if err != nil {
		log.Fatal("Ошибка базы данных:", err)
	}

}
