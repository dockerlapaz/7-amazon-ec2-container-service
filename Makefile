.PHONY: build

build:
	# Obtenemos el cliente de redis para Go
	go get -d -v github.com/garyburd/redigo/redis
	# Compilamos la aplicación para Linux
	GOOS=linux go build -o app .
	# Construimos el contenedor usando docker-compose
	docker-compose build --force-rm --no-cache

run:
	# Corremos ambos contenedores con docker-compose
	# Presiona CTRL + C para detener ambos contenedores
	docker-compose up