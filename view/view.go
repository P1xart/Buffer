package view

import (
	"log"
)

func StartLog() {
	log.Println("Starting Bufer...")
}

func ErrorWriteData(err error) {
	log.Println("Не удалось записать данные на целевой сервер\n", err) // Метод при любой ошибке.
}

func StartBuferFunc() {
	log.Println("Starting BuferFunc...")
}

func StartHttpHandler() {
	log.Println("Starting http handler...")
}
