package main

import (
	"github.com/p1xart/bufer/controller"
	"github.com/p1xart/bufer/model"
	"github.com/p1xart/bufer/view"
)

func main() {
	view.StartLog() // Лог о успешной начале работы Bufer API
	go model.Bufer() // Асинхронно запускаем сам буфер в отдельном потоке, дабы не мешать принимать запросы
	controller.AcceptRequests() // Принимаем запросы
}