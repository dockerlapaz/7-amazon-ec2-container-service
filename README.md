# Docker Nights 7: Amazon EC2 Container Service

##Compilando la aplicación

```bash
# Necesitas GO instalado para compilar la aplicación
$ GOOS=linux go build -o app .
$ docker-compose build
$ docker-compose up
```

Abre http://localhost:8080/

## Subiendo la aplicación a un repositorio privado en ECS

```bash
# Reemplaza $ECS_REPO_URL con la url del repositorio de AWS
$ docker build -t $ECS_REPO_URL:v1 .
$ docker push $ECS_REPO_URL:v1
```