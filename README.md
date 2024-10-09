# BUFFER #

Для запуска:  
В корне создать файл .env и поместить туда Bearer token как "TOKEN=Bearer your_token", например "TOKEN=Bearer 124f31d1a7ad7f89qgdxg8782bcv192"  
Запустить вручную
1) Загрузить зависимости с помощью "$ go mod download"  
2) Ввети команду "$ go run cmd/buffer/main.go"   
ИЛИ использовать docker  
1) Собрать образ "$ docker build -t buffer ."  
2) Запустить контейнер "$ docker run -p 6700:6700 buffer"

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
Authorization:Bearer t0k3n

Передаваемые данные не валидируются из-за недостатка информации. Только Bearer Token
