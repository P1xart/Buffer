package main

import (
	"github.com/p1xart/bufer/internal/buffer"
	"github.com/p1xart/bufer/internal/api"
)

func main() {
	go buffer.Buffer() // Асинхронно запускаем сам буфер в отдельном потоке, дабы не мешать принимать запросы
	api.AcceptRequests() // Принимаем запросы
}