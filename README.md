A continuación encontrarás un ejemplo de fichero README.md que sirve como manual de instrucciones para tu proyecto. Se incluyen detalles de instalación, configuración y despliegue con Docker/Docker Compose, así como notas de uso y buenas prácticas. Puedes adaptarlo según las necesidades de tu equipo y organización.

---

# test-service

Este proyecto es un servicio escrito en Go que utiliza el framework [GoFr](https://gofr.dev/) para levantar un servidor y conectarse a Redis. Proporciona un endpoint de ejemplo que devuelve un saludo básico y otro que interactúa con Redis.

## Índice
- [test-service](#test-service)
  - [Índice](#índice)
  - [1. Requisitos Previos](#1-requisitos-previos)
  - [2. Estructura del Proyecto](#2-estructura-del-proyecto)
  - [3. Configuración de Entornos](#3-configuración-de-entornos)
    - [Cómo se carga la configuración](#cómo-se-carga-la-configuración)
  - [4. Ejecución en Entorno Local](#4-ejecución-en-entorno-local)
  - [5. Despliegue con Docker](#5-despliegue-con-docker)
  - [6. Despliegue con Docker Compose](#6-despliegue-con-docker-compose)
    - [Pasos para levantar los contenedores](#pasos-para-levantar-los-contenedores)
  - [7. Uso de la Aplicación](#7-uso-de-la-aplicación)
    - [Ejemplo de Respuesta](#ejemplo-de-respuesta)
  - [8. Notas y Buenas Prácticas](#8-notas-y-buenas-prácticas)
  - [9. Licencia](#9-licencia)

---

## 1. Requisitos Previos

- [Go](https://golang.org/dl/) (v1.18 o superior).
- [Docker](https://www.docker.com/) instalado y configurado (versión 20.x o superior).
- [Docker Compose](https://docs.docker.com/compose/) (versión 1.29 o superior).
- (Opcional) Redis instalado localmente para pruebas rápidas, si no se ejecuta mediante contenedor.

---

## 2. Estructura del Proyecto

```
test-service/
├─ configs/
│  ├─ .local.env
│  ├─ .dev.env
│  ├─ .staging.env
│  ├─ .prod.env
├─ main.go
├─ Dockerfile
├─ docker-compose.yml
├─ go.mod
├─ go.sum
└─ README.md
```

- **main.go**: Punto de entrada de la aplicación. Define endpoints y conexión con Redis.
- **configs/**: Contiene ficheros de configuración por entorno.  
  - `.local.env`: Configuración local (por defecto si `APP_ENV` no está definido).
  - `.dev.env`, `.staging.env` y `.prod.env`: Configuraciones específicas para cada entorno.
- **Dockerfile**: Archivo para construir la imagen Docker del servicio.
- **docker-compose.yml**: Archivo para orquestar la aplicación y Redis en contenedores.
- **go.mod** y **go.sum**: Gestionan dependencias del proyecto Go.
- **README.md**: Documento principal que explica la instalación, configuración y uso.

---

## 3. Configuración de Entornos

La aplicación utiliza variables de entorno para configurar:
- **APP_NAME**: Nombre de la aplicación.
- **HTTP_PORT**: Puerto en el que se expone el servicio.
- **REDIS_HOST**: Host o IP donde se encuentra el servidor Redis.
- **REDIS_PORT**: Puerto donde escucha Redis.
- **REDIS_PASSWORD**: Contraseña para Redis (si aplica).
- **REDIS_DB**: Número de base de datos a usar en Redis.

Ejemplo de `.local.env` (utilizado cuando no se define `APP_ENV`):
```
APP_NAME=test-service
HTTP_PORT=9000
REDIS_HOST=gofr-redis
REDIS_PORT=6379
REDIS_PASSWORD=password
REDIS_DB=2
```

### Cómo se carga la configuración

1. Si existe la variable de entorno `APP_ENV`, la aplicación buscará `.dev.env`, `.staging.env` o `.prod.env` en el directorio `configs/`, según corresponda.
2. Si no se define `APP_ENV`, se toma por defecto `.local.env`.  
3. Si no encuentra el fichero correspondiente, buscará `.env`.

---

## 4. Ejecución en Entorno Local

Para ejecutar la aplicación directamente con Go:

1. Clona el repositorio y ubícate en el directorio raíz del proyecto:
   ```bash
   git clone <URL-DEL-REPO>
   cd test-service
   ```
2. Asegúrate de tener un Redis en ejecución o ajusta las variables en `.local.env` para apuntar al tuyo.
3. Ejecuta:
   ```bash
   go run main.go
   ```
4. Si deseas especificar un entorno, por ejemplo `dev`:
   ```bash
   APP_ENV=dev go run main.go
   ```
5. Por defecto, la aplicación se levantará en el puerto `9000`. Puedes verificarlo abriendo en un navegador:
   ```
   http://localhost:9000/greet
   ```

---

## 5. Despliegue con Docker

1. Construye la imagen localmente:
   ```bash
   docker build -t test-service:local .
   ```
2. Ejecuta el contenedor:
   ```bash
   docker run -p 9000:9000 --name gofr-app test-service:local
   ```
3. Verifica el endpoint principal:
   ```
   http://localhost:9000/greet
   ```

Si necesitas sobreescribir variables de entorno, puedes pasar `-e VAR=valor` al comando `docker run` o crear un fichero `.env` y usar la opción `--env-file`.

---

## 6. Despliegue con Docker Compose

El fichero `docker-compose.yml` define los servicios `hello-app` (nuestra aplicación Go) y `gofr-redis` (el contenedor de Redis).

```yml
services:
  hello-app:
    build: .
    container_name: gofr-app
    environment:
      - APP_NAME=test-service
      - HTTP_PORT=9000
      - REDIS_HOST=gofr-redis
      - REDIS_PORT=6379
      - REDIS_PASSWORD=password
      - REDIS_DB=2
    ports:
      - "9000:9000"
    depends_on:
      - gofr-redis
    networks:
      - gofr-network

  gofr-redis:
    image: redis:latest
    container_name: gofr-redis
    environment:
      - REDIS_HOST=localhost
      - REDIS_PORT=6379
      - REDIS_PASSWORD=password
      - REDIS_DB=2
    ports:
      - "6379:6379"
    command: redis-server --requirepass password
    networks:
      - gofr-network

networks:
  gofr-network:
    driver: bridge
```

### Pasos para levantar los contenedores

1. Clona el repositorio y ubícate en el directorio raíz.
2. Ejecuta:
   ```bash
   docker-compose up --build
   ```
   Esto construirá la imagen (si no existe) y levantará tanto la aplicación como Redis.
3. Para levantarlo en segundo plano (background):
   ```bash
   docker-compose up -d --build
   ```
4. Verifica que todo está funcionando:
   - Abre en el navegador:
     ```
     http://localhost:9000/greet
     ```
   - Observa los logs en tiempo real:
     ```bash
     docker-compose logs -f
     ```

---

## 7. Uso de la Aplicación

La aplicación expone dos endpoints principales:

1. **/greet**  
   - Devuelve un saludo: `"Hello, World!"`.
   - Uso (GET):
     ```
     curl http://localhost:9000/greet
     ```
2. **/redis**  
   - Conecta a Redis, setea la clave `greeting` con el valor `"Hello, Redis!"` y lo devuelve.
   - Uso (GET):
     ```
     curl http://localhost:9000/redis
     ```
   - Respuesta de ejemplo: `"Hello, Redis!"`.

### Ejemplo de Respuesta
```json
"Hello, Redis!"
```

---

## 8. Notas y Buenas Prácticas

- **Seguridad en Producción**  
  - Asegúrate de cambiar la contraseña de Redis (`REDIS_PASSWORD`) antes de exponer el servicio.
  - Limita el acceso al puerto de Redis o solo permítelo desde redes o contenedores de confianza.
- **Variables de Entorno**  
  - Usa ficheros `.env` para cada entorno y no subas estos ficheros sensibles a repositorios públicos.
  - Para un entorno de integración/producción, utiliza herramientas de gestión de secretos como HashiCorp Vault, AWS Secrets Manager, etc.
- **Logs y Monitorización**  
  - Revisa la salida de logs con `docker-compose logs -f` o tus herramientas preferidas (ej. ELK, CloudWatch).
  - Añade métricas si tu entorno de producción requiere monitorización avanzada.
- **Escalabilidad**  
  - Utiliza Docker Compose con un orquestador (ej. Kubernetes) si necesitas escalar a varios contenedores.
  - Para alta disponibilidad de Redis, considera Redis Cluster o soluciones como Amazon ElastiCache.

---

## 9. Licencia

Indica aquí la licencia del proyecto (ej. MIT, Apache 2.0, etc.) o las condiciones de uso internas si es un repositorio privado.

---

¡Listo! Con esto dispones de una guía completa de instalación, configuración y despliegue de tu servicio escrito en Go con GoFr y Redis, incluyendo instrucciones para Docker y Docker Compose. Este README.md puede crecer con más secciones según tus necesidades (contribución, testing, CI/CD, etc.).