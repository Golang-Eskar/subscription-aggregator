package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Golang-Eskar/subscription-aggregator/internal/database"
	"github.com/Golang-Eskar/subscription-aggregator/internal/router"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Запуск")

	err1 := godotenv.Load(".env")
	if err1 != nil {
		return
	}
	err := database.Init()
	if err != nil {
		log.Fatal("Ошибка базы данных:", err)
	}

	r := router.New()

	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", r)

}
