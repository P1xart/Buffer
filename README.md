# BUFFER #

Тестовое задание. https://hh.ru/resume/0d251099ff0cd8444c0039ed1f757452314643

Для запуска ввести "$ go run cmd/buffer/main.go" (Или скомпилировать и запустить командой "$ go build cmd/buffer/main.go && ./cmd/buffer/main")

На http://localhost:6700/api/buffer, методом POST, типом multipart/form-data передать значения:  
period_start:2024-05-01  
period_end:2024-05-31  
period_key:month  
indicator_to_mo_id:227373  
indicator_to_mo_fact_id:0  
value:1  
fact_time:2024-05-31  
is_plan:0  
auth_user_id:40  
comment: buffer Last_name  

И Bearer Token в заголовке, иначе вернет 401!!!  
Authorization:Bearer <token>

Я бы мог обрабатывать ответы и в случае неудачи, переотправлять или еще как-то обрабатывать ошибку. Но сказано сильно не заморачиваться. Вам стоит быть внимательным. Микросервис НЕ вернет ошибку в случае отправки невалидного запроса, а просто пропустит и продолжит работать.