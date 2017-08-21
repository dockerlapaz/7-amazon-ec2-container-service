# Docker Nights 7: Amazon EC2 Container Service

## Requisitos
* Una cuenta en [Amazon Web Services](https://aws.amazon.com/)
* [Docker](https://www.docker.com) 17.05 o superior

## Compilando la aplicación

Necesitas instalar [Go 1.8](https://golang.org) o superior para compilar la aplicación.

```bash
$ make
$ docker-compose up
```

La aplicación estará disponible en [http://localhost:8080/](http://localhost:8080).

## Creando un repositorio privado en ECS

```bash
# Para correr el comando es necesario configurar el cli aws con un access_key y secret_key
$ aws ecr create-repository --repository-name miapp

# Resultado
{
    "repository": {
        "registryId": "1234567890",
        "repositoryName": "miapp",
        "repositoryArn": "arn:aws:ecr:us-east-1:1234567890:repository/miapp",
        "createdAt": 123456789.0,
        "repositoryUri": "aws_account_id.dkr.ecr.us-east-1.amazonaws.com/miapp"
    }
}
```

Copia la url de `repositoryUri` para generar el contenedor más adelante.

Obtiene las credenciales del repositorio en ECR:

```bash
$ aws ecr get-login --no-include-email

# Resultado
docker login -u AWS -p unpasswordmuylargo https://aws_account_id.dkr.ecr.us-east-1.amazonaws.com
```
Copia y pega el resultado para ingresar al repositorio desde Docker.

## Subiendo la aplicación a un repositorio privado en ECS

```bash
# Reemplaza repositoryUri con la url del comando aws ecr create-repository
$ docker tag dockerlapaz/votacion:go repositoryUri:v1
$ docker push repositoryUri:v1
```




