# ServerWithNatsL0
1. В директории, где находится docker-compose.yaml файл, командой sudo docker-compose build создаем docker image.
2. Командой sudo docker-compose up запускаем docker container.
3. Переходим в директорию cmd и выполняем команду go run main.go для запуска сервера.
4. В директории cmd выполняем команды go run pubnats.go  и go run newpub.go для публиации данных на сервер.
