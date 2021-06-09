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

### Запуск
docker-compose up
порты:
* :8085 - gRPC
* :8084 - gRPC-gateway 
* :8083 - OpenAPI Swagger UI
