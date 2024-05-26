package model

import (
	"log"
	"net/http"

	"github.com/p1xart/bufer/controller"
	"github.com/p1xart/bufer/view"
)

const(
	bearerToken = "Bearer 48ab34464a5573519725deb5865cc74c" // Bearer token как константа
	dstUrl = "https://development.kpi-drive.ru/_api/facts/save_fact" // Адресс назначения
)

func Bufer() {
	view.StartBufferFunc() // Обьявляем успешный запуск буфера
	for {
		formData, payload, empty := controller.GetRequest() // Получаем информацию о запросе
		if empty { // Проверяем, есть ли запросы в очереди
			continue
		}
		request, err := http.NewRequest("POST", dstUrl, payload) // Создаем новый запрос
		if err != nil {
			view.ErrorWriteData(err)
		}

		request.Header.Add("Authorization", bearerToken) // Передаем Bearer token в заголовок
		request.Header.Add("Content-Type", formData.FormDataContentType()) // Передаем тип контента multipart/form-data
		client := http.Client{} // Создаем клиент 
		_, err = client.Do(request) // Передаем в клиент запрос и отправляем, А ПОСЛЕ ЖДЕМ ОТВЕТА, что означает - записано в бд.
		if err != nil {
			view.ErrorWriteData(err)
		}
		log.Println("Отправлен запрос из очереди")
	}
}
