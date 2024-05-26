package view

import (
	"log"
)

func StartLog() {
	log.Println("Starting Buffer...")
}

func ErrorWriteData(err error) {
	log.Println("Не удалось записать данные на целевой сервер\n", err) // Метод при любой ошибке.
}

func StartBufferFunc() {
	log.Println("Starting BufferFunc...")
}

func StartHttpHandler() {
	log.Println("Starting HttpHandler...")
}
