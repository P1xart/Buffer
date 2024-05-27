package api

import (
	"bytes"
	"log"
	"mime/multipart"
	"net/http"
	"sync"
	"os"
	
)

type request struct { // Стуктура запроса на save_fact
	bearerToken             string
	period_start            string
	period_end              string
	period_key              string
	indicator_to_mo_id      string
	indicator_to_mo_fact_id string
	value                   string
	fact_time               string
	is_plan                 string
	auth_user_id            string
	comment                 string
}

var mutex sync.Mutex = sync.Mutex{} // Создадим мьютекс для избежания одновременной записи в очередь запроса от двух горутин
var requests []request              // Массив экземпляров структур запроса. Массив очереди запросов.

func AcceptRequests() { // API /api/buffer
	log.Println("Запущен API")
	go http.HandleFunc("/api/buffer", sendHandler) // Вызываем обработчик асинхронно для одновременной обработки нескольких входных запросов

	err := http.ListenAndServe(":6700", nil) // Запуск буфера на localhost:6700
	if err != nil {
		log.Fatal(err)
	}
}

func sendHandler(w http.ResponseWriter, r *http.Request) { // Обработчик каждого запроса
	if r.Method == "POST" { // Если метод POST, в остальных случаях игнорируем
		data := request{ // Получаем запрос в структуру и добавляем в очередь запросов
			bearerToken:             r.Header.Get("Authorization"),
			period_start:            r.FormValue("period_start"),
			period_end:              r.FormValue("period_end"),
			period_key:              r.FormValue("period_key"),
			indicator_to_mo_id:      r.FormValue("indicator_to_mo_id"),
			indicator_to_mo_fact_id: r.FormValue("indicator_to_mo_fact_id"),
			value:                   r.FormValue("value"),
			fact_time:               r.FormValue("fact_time"),
			is_plan:                 r.FormValue("is_plan"),
			auth_user_id:            r.FormValue("auth_user_id"),
			comment:                 r.FormValue("comment"),
		}
		bearToken, exists := os.LookupEnv("TOKEN")
		if !exists {
			log.Fatal("Error: Не получается найти переменную TOKEN в среде окружения")
		}
		if data.bearerToken != bearToken {
			w.WriteHeader(401) // Юзер не авторизован, запрос отклоняется для записи.
			return
		}
		mutex.Lock() // Защищаем массив от одновременной записи
		defer mutex.Unlock()
		requests = append(requests, data) // Запись в массив запросов
		w.WriteHeader(201) // Успешно обработано, но ничего не вернем в ответе
	} else {
		w.WriteHeader(405) // Все запросы кроме POST - method not allowed
	}
}

func GetRequest() (string, *bytes.Buffer, string, bool) { // Метод получения запроса из массива "ожидания"
	if len(requests) == 0 { // Проверяем, есть ли запросы в очереди. Если нет - вернем пустой Content-Type, пустое тело, токен и булево true (Да, пуст)
		return "", &bytes.Buffer{}, "", true
	}

	payload := bytes.Buffer{}               // Создадим экземпляр полезной нагрузки для ее заполнения (form-data)
	writer := multipart.NewWriter(&payload) // Экземпляр функции, которая и будет писать тело в payload
	request_body := requests[0]             // Берем первый запрос из очереди
	writer.WriteField("period_start", request_body.period_start)
	writer.WriteField("period_end", request_body.period_end)
	writer.WriteField("period_key", request_body.period_key)
	writer.WriteField("indicator_to_mo_id", request_body.indicator_to_mo_id)
	writer.WriteField("indicator_to_mo_fact_id", request_body.indicator_to_mo_fact_id)
	writer.WriteField("value", request_body.value)
	writer.WriteField("fact_time", request_body.fact_time)
	writer.WriteField("is_plan", request_body.is_plan)
	writer.WriteField("auth_user_id", request_body.auth_user_id)
	writer.WriteField("comment", request_body.comment)
	err := writer.Close()
	if err != nil {
		log.Println("Error: Запрос не был отправлен.")
	}
	mutex.Lock() // Защищаемся от одновременного удаления запроса из очереди ожидания
	defer mutex.Unlock()
	copy(requests[0:], requests[0+1:]) // Удаляем текущий запрос, так как он уже отправляется в целевой микросервис
	requests[len(requests)-1] = request{}
	requests = requests[:len(requests)-1]

	return writer.FormDataContentType(), &payload, request_body.bearerToken, false // Возвращаем Content-Type, тело, bearer token и булево false (нет, не пуст)
}
