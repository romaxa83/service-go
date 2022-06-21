# service-golang

This is a simple example to build a microservice in Go. It has an HTTP server and RabbitMQ client packages.

## Command for run

```bash
// поднять сервисы для разработки
make up
```
```bash
// собрать и запустить сервис
make build-service
```
```bash
// сборка сервиса в докер образа 
make build-docker-service
```
```bash
// запуск сервисов, где сервис на go запускаеться в контейнере
make up-service
```
```bash
// остановка сервисов
make down
```
## Accessing the API

Base route: [localhost:8000](http://localhost:8000).

To check the bitcoin variation on a period: [localhost:8000/bitcoin/startdate/YYYY-MM-DD/enddate/YYYY-MM-DD](http://localhost:8000/bitcoin/startdate/2018-11-01/enddate/2018-11-30)