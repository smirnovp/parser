# Моё решение тестового задания

Необходимо сделать gRPC обёртку над сайтом https://www.rusprofile.ru/

## API

Сервис должен реализовывать один метод, принимающий на вход ИНН компании, ищущий компанию на rusprofile, и возвращающий её ИНН, КПП, название, ФИО руководителя

## Технологии

* сервис должен быть написан на Go с использованием Go Modules
* предоставлять API через [gRPC](https://grpc.io/docs/languages/go/quickstart/)
* предоставлять API через HTTP с помощью [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway)
* предоставлять Swagger UI с документацией, сгенерированной из .proto файла с помощью protoc-gen-swagger
* быть упакован в docker контейнер

## Решено
Запуск: `docker-compose up`

порты:
* `localhost:8085` - gRPC
* `localhost:8084` - gRPC-gateway 
* `localhost:8083` - OpenAPI Swagger UI документация

## Примеры
* `curl -X 'GET' 'http://localhost:8084/inn/7721679536'` - проверка gRPC сервиса
* `http://localhost:8084/inn/7721679536`  - проверка gRPC-gateway сервиса
* `http://localhost:8083` - запуск Swagger UI
