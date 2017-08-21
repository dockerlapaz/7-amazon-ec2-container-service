# Docker Nights 7: Amazon EC2 Container Service

## Requisitos
* Una cuenta en [Amazon Web Services](https://aws.amazon.com/)
* El [CLI de AWS](https://aws.amazon.com/cli) instalado y configurado
* Instalar el [CLI de ECS](https://aws.amazon.com/ecs)
* [Docker](https://www.docker.com)

## Compilando la aplicación

```bash
# Necesitas GO instalado para compilar la aplicación
$ GOOS=linux go build -o app .
$ docker-compose build
$ docker-compose up
```

Abre [http://localhost:8080/](http://localhost:8080/)

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
$ docker build -t repositoryUri:v1 .
$ docker push repositoryUri:v1
```

## Generando un cluster para ECS

Lanza un cluster usando este link: [Lanzar cluster en ECS](https://console.aws.amazon.com/ecs/home?region=us-east-1#/clusters/create/new)

* **Cluster Name:** miappCluster
* **EC2 instance type:** t2.small
* **Number of instances:** 2

**Importante:** Un keypair (llave SSH) debe existir para poder ingresar a los servidores que correrán los contenedores. [Crear llave SSH](https://console.aws.amazon.com/ec2/v2/home?region=us-east-1#KeyPairs:sort=keyName)

_El resto de las opciones puede dejarse por defecto._

## Listar las instancias de un cluster

```bash
$ aws ecs list-container-instances --cluster miappCluster
```

## Crear un dominio interno en Route 53
Usa este link para ir a Route 53: [Panel Route 53](https://console.aws.amazon.com/route53/home?region=us-east-1#hosted-zones:)

Click en `Create Hosted Zone`:

* **Domain name:** miapp.internal
* **Command:** Dominio interno para miapp
* **Type**: A private hosted zone for Amazon VPC
* **VPC ID**: _Encontrar el VPC de US East (N. Virginia)_

Click en `Create` para terminar la creación del dominio interno.

## Crear task definitions

Redis:

```json
$ aws ecs register-task-definition --cli-input-json file://redis.json
```
App:

```bash
aws ecs register-task-definition --cli-input-json file://app.json
```





