# Веб-сервис для подсчёта арифметических выражений
## Описание
Этот проект реализует веб-сервис, который вычисляет арифметические выражения, переданные пользователем через HTTP-запрос.

## Запуск 
С помощью git clone github.com/MrM2025/rpforcalc/tree/master/calc_go сделайте клон проекта. 

Далее перейдите в папку calc_go(например с помощью cd ./calc_go) и напишите в терминале go mod tidy

Запуск происходит при помощи команды go run ./cmd/main.go (!!!Важно - проверьте, что вы находитесь в папке calc_go).

#### При обращении к http://localhost:8080 будет возвращен readme-файл

Если у вас Windows воспользуйтесь этим примером при задании порта - set "PORT=8087" -and (go run ./cmd/main.go)

Выражение для вычисления должно передаваться в JSON-формате, в единственном поле "expression", если поле отсутствует - сервер вернет ошибку 422, "Empty expression"; если в запросе будут поля, отличные от "expression" - сервер вернет ошибку 400, "Bad request" также как и при отсуствии JSON'а в теле запроса;

## Пример запроса с использованием curl(Рекомендую использовать постман)
Для cmd:  

 curl -X POST http://localhost:8080/api/v1/calculate -H "Content-Type: application/json" -d "{"expression": "1"}" (пример корректного запроса, код:200)

Для git bash:

curl --location 'localhost:8080/api/v1/calculate'
--header 'Content-Type: application/json'
--data '{ "expression": "-1+1*2.54+41+((3/3+10)/2-(-2.5-1+(-1))*10)-1" }'
#

Postman:

https://identity.getpostman.com/signup?deviceId=c30fc039-7460-4f58-8cb9-b74256c4186c  

^

|

Регестрация

https://www.postman.com/downloads/

^

|

Ссылка на скачивание приложения    

#
Инструкция по эксплуатации №1 - https://timeweb.com/ru/community/articles/kak-polzovatsya-postman

Инструкция по эксплуатации №2 - https://blog.skillfactory.ru/glossary/postman/

Инструкция по эксплуатации №3 - https://gb.ru/blog/kak-testirovat-api-postman/

## Примеры использования

Верно заданный запрос, Status: 200

curl -X POST http://localhost:8080/api/v1/calculate -H "Content-Type: application/json" -d "{\"expression": \"20-(9+1)\"}"



Запрос с пустым выражением, Status: 422, Error: empty expression


curl -X POST http://localhost:8080/api/v1/calculate -H "Content-Type: application/json" -d "{\"expression": \"\"}"



Запрос с делением на 0, Status: 422, Error: division by zero


curl -X POST http://localhost:8080/api/v1/calculate -H "Content-Type: application/json" -d "{\"expression": \"1/0\"}"



Запрос неверным выражением, Status : 422, Error: invalid expression


curl -X POST http://localhost:8080/api/v1/calculate -H "Content-Type: application/json" -d "{\"expression\": \"1++*2\"}"


## Тесты
Для тестирования перейдите в папку application_test.go и используйте команду go test или(для вывода дополнительной информации) go test -v

Для запусков всех тестов разом воспользуйтесь - go test ./...

