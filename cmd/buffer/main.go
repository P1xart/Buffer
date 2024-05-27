package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/p1xart/bufer/internal/api"
	"github.com/p1xart/bufer/internal/buffer"
)

func main() {
	go buffer.Buffer()   // Асинхронно запускаем сам буфер в отдельном потоке, дабы не мешать принимать запросы
	api.AcceptRequests() // Принимаем запросы
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env")
	}
}
