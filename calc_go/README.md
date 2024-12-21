# Веб-сервис для подсчёта арифметических выражений
## Описание
Этот проект реализует веб-сервис, который вычисляет арифметические выражения, переданные пользователем через HTTP-запрос.

## Запуск 
С помощью git clone github.com/MrM2025/rpforcalc/tree/master/calc_go сделайте клон проекта. 

Далее перейдите в папку calc_go(например с помощью cd ./calc_go) и напишите в терминале go mod tidy

Запуск происходит при помощи команды go run ./cmd/main.go (!!!Важно - проверьте, что вы находитесь в папке calc_go).

Если у вас Windows соберите exeфайл с помощью команды - go build -o calc.exe 

После задайте порт - set "PORT=8087" & "calc.exe" 

## Пример запроса с использованием curl(Рекомендую использовать постман)
Для cmd :  

 curl -X POST http://localhost:8080/api/v1/calculate -H "Content-Type: application/json" -d "{"expression": "1"}" (пример корректного запроса, код:200)

git bash

curl --location 'localhost:8080/api/v1/calculate'
--header 'Content-Type: application/json'
--data '{ "expression": "2+2*2" }'

Postman:

https://identity.getpostman.com/signup?deviceId=c30fc039-7460-4f58-8cb9-b74256c4186c  

^

|

Регестрация

https://www.postman.com/downloads/

^

|

Ссылка на скачивание приложения

Инструкция по эксплуатации №1 - https://timeweb.com/ru/community/articles/kak-polzovatsya-postman

Инструкция по эксплуатации №2 - https://blog.skillfactory.ru/glossary/postman/

Инструкция по эксплуатации №3 - https://gb.ru/blog/kak-testirovat-api-postman/


## Тесты
Для тестирования перейдите в папку application_test.go и используйте команду go test или(для вывода дополнительной информации) go test -v

Для запусков всех тестом разом воспользуйтесь - go test ./...

Для запусков всех тестом разом воспользуйтесь - go test ./...
