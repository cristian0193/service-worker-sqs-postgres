# Service Worker SQS - Postgres

## Tabla de contenido
1. [Contexto](#contexto)
2. [Tecnologías](#tecnologías)
3. [Arquitectura](#arquitectura)
    * [Estructura del proyecto](#estructura-del-proyecto)
5. [Despliegues](#despliegues)
    * [Local](#local)
6. [Endpoints](#endpoints)
7. [Queues](#queues)


<a name="contexto"></a>
# Contexto 📋

El objetivo de este template es proporcionar una estructura básica para crear un servicio de trabajador (worker) que establezca conexión con servicios de AWS, como SQS y RDS. Esta plantilla tiene como finalidad simplificar la construcción de servicios similares, al proporcionar una configuración sencilla y reducir el tiempo necesario para su desarrollo.

- **Estructurar un servicio worker**: Crear una estructura clara y organizada para desarrollar un servicio de trabajador que cumpla con los requisitos del proyecto.
- **Establecer conexión con servicios AWS**: Configurar la conexión con servicios de AWS, como Amazon SQS (Simple Queue Service) y RDS (Relational Database Service).
- **Simplificar la configuración**: Proporcionar una configuración sencilla que permita a los desarrolladores personalizar y adaptar el servicio según sus necesidades específicas.
- **Reducir tiempos de construcción**: Agilizar el proceso de desarrollo al brindar una plantilla predefinida y una estructura base que sirva como punto de partida para la construcción del servicio.
- **Mejorar la reutilización de código**: Fomentar la reutilización de código al proporcionar una estructura común y componentes predefinidos que puedan ser utilizados en diferentes proyectos de servicios de trabajador.


<a name="tecnologías"></a>
# Tecnologias 💻

**Dependencies** 🤝
Las siguientes dependencias se utilizan en el desarrollo para llevar a cabo depliegue de servidor http, conexiones SQS y RDS, entre otros.

* [github.com/aws/aws-sdk-go](https://github.com/aws/aws-sdk-go): SDK oficial de AWS para el lenguaje de programación Go.

* [github.com/labstack/echo/v4](https://github.com/labstack/echo): Echo es un framework de aplicaciones web Go de código abierto, extensible y centrado en el rendimiento.

* [gorm.io/gorm](https://gorm.io/): Es una libreria que permite el mapeo objeto-relacional (ORM), es una técnica que permite consultar y manipular datos de una base de datos utilizando un paradigma orientado a objetos.

* [gorm.io/driver/postgres](https://github.com/go-gorm/postgres): Controlador que permite establecer la conexion entre una base de datos postgres y un cliente.

* [go.uber.org/zap](https://pkg.go.dev/go.uber.org/zap): Permite la configuracion de logs.

**Framework**

* [Echo](https://echo.labstack.com/)
* [Gorm](https://gorm.io/)

**Servicios AWS**

* [SQS](https://aws.amazon.com/es/sqs/)

**Bases de datos**

* [Postgresql](https://www.postgresql.org/)

<a name="arquitectura"></a>
# Arquitectura 🏢

Para del proyecto se toma como base los principios de las arquitecturas limpias, utilizando en este caso gran parte del concepto de **arquitectura multicapas**, lo cual permite la independencia de frameworks, entidades externas y UI, por medio de capas con responsabilidad únicas que permite ser testeables mediante el uso de sus interfaces. Como parte de las buenas prácticas la solución cuenta en su gran mayoría con la aplicación de los principios SOLID, garantizando un código limpio, mantenible, reutilizable y escalable.

![service-template-golang](https://github.com/cristian0193/service-worker-sqs-postgres/assets/11803196/7d79f234-cfd9-4043-baf4-ea4ecbd61d99)


<a name="estructura-del-proyecto"></a>
### * **Estructura del proyecto** 🧱

- [x] `clients/`: contiene la implementacion de los clients externos
  - [ ] `awssqs/`: define el cliente para aws sqs
- [x] `cmd/`: administra los recursos de llamados al api
  - [ ] `builder/`: construye cada una de las instancias transversales
- [x] `consumer/`: define la logica para obtener los mensajes desde el consumidor
- [x] `database/`: define la logica para obtener la conexion a base de datos (postgresql)
- [x] `domain/`: administracion de los datos de manera transversal
- [x] `http/`: administra los clientes api-rest
  - [ ] `controllers/`: define los handler 
  - [ ] `services/`: define los services asociados a los repositorios
  - [ ] `repository/`: define las consultas, actializacion o inserciones a la base de datos
- [x] `processor/`: define el inicio del proceso para la lectura de mensajes desde SQS
- [x] `utils/`: define las funciones transversales

<a name="despliegues"></a>
# Despliegues 🚀

Para la fase de despliegue a nivel local, se utilizaron algunas herramientas que nos permite agilizar este proceso. Como fase se mostrara el paso a paso:

> **Nota:** Para el proceso se deben definir las variables de ambiente que nos permite establecer conexion a los diferentes servicios.

```
APPLICATION_ID=
SERVER_PORT=
LOG_LEVEL=INFO

AWS_ACCESS_KEY=
AWS_SECRET_KEY=
AWS_REGION=

AWS_SQS_URL=
AWS_SQS_MAX_MESSAGES=
AWS_SQS_VISIBILITY_TIMEOUT=

DB_PORT=
DB_HOST=
DB_NAME=
DB_USERNAME=
DB_PASSWORD=
```

<a name="local"></a>
### * **Local** 

En el proceso local podemos utilizar despliegues de contenedores con postgres RDS o local.

    1. Instalacion postgres local o aws (rds)
        - docker pull postgres
        - https://aws.amazon.com/es/rds/

    2. Creacion de SQS en AWS
        - https://aws.amazon.com/es/sqs/

    3. La automigracion a nivel de gorm a las base de datos esta activa

    4. Definir las variables de entorno definidas en (Despliegues)

    6. Ejecutar el comando 'go run main.go'

    7. Puerto :8080 run

<a name="endpoints"></a>
# Endpoints 🤖

- **GET**    http://localhost:8080/sqs/:id
```
curl --location --request GET 'http://localhost:8080/sqs/:id'
```

- **Response**
```
  {
    "id": "7a312c5a-e69e-4935-9b33-5dc33919a76f",
    "message": "Hola Mundo!!",
    "date": "2023-06-13T17:48:05-05:00"
  }
```

<a name="queues"></a>
# Queues 📨

- **URL**    https://sqs.us-east-1.amazonaws.com/XXXXXXXX/sqs-service-template-golang


- **Message**
```
    {
      "msg": "Hello World"
    }
```

# Author 🧑‍💻
```
- Christian Alexis Rodriguez Castillo
- Sr Software Engineer - Mercadolibre
- cristian010193@gmail.com
```
