package buffer

import (
	"github.com/p1xart/bufer/internal/api"
	"io"
	"log"
	"net/http"
)

func Buffer() {
	log.Println("Запущен Buffer") // Обьявляем успешный запуск буфера
	for {
		contentType, payload, bearerToken, empty := api.GetRequest() // Получаем информацию о запросе
		if empty {                                                   // Проверяем, есть ли запросы в очереди
			continue
		}
		request, err := http.NewRequest("POST", "https://development.kpi-drive.ru/_api/facts/save_fact", payload) // Создаем новый запрос
		if err != nil {
			log.Println("Error: Запрос не был отправлен.")
		}

		request.Header.Add("Authorization", bearerToken) // Передаем Bearer token в заголовок
		request.Header.Add("Content-Type", contentType)  // Передаем тип контента multipart/form-data
		client := http.Client{}                          // Создаем клиент
		response, err := client.Do(request)              // Передаем в клиент запрос и отправляем, А ПОСЛЕ ЖДЕМ ОТВЕТА, что означает - записано в бд.
		if err != nil {
			log.Println("Error: Запрос не был отправлен.")
		}
		b, err := io.ReadAll(response.Body)
		if err != nil {
			log.Println("Error: Запрос не был отправлен.")
		}
		log.Println("Отправлен запрос из очереди\n", string(b))
	}
}
