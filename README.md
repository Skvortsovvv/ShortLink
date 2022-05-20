 ##API Приложения для укорочения URL
 ## Для запуска
1. Поднятие базы данных:
````
make base
````
Перед первым запуском приложения:
````
make migrate
````
2. Docker:
````
docker build .
````
3. Запуск приложения:
````
docker run --rm --network="host" -e WORKMODE=<memory или db> <app>
````

