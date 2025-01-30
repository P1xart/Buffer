# BUFFER #
Тестовое задание. Не актуально, но как интересный опыт, имеет место быть.

## Запуск ##
Создайте файл виртуального окружения, поместив туда ваш bearer токен: `echo "TOKEN=Bearer CHANGE_ME" >> .env`

Ручная сборка:
```
go mod download
go run cmd/buffer/main.go
```

Docker:
```
docker build -t buffer .
docker run -p 6700:6700 buffer
```

## Документация ##
```
POST:/api/buffer
Content-Type: multipart/form-data
Security: Bearer
```
Keys and values of content:

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
