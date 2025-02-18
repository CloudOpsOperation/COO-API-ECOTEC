# API en Golang - COO-API-ECOTEC

Este repositorio contiene una API desarrollada en Golang para gestionar datos relacionados con *COO-API-ECOTEC*. La API está estructurada en varias capas para mantener un código modular y fácil de mantener.

## 📁 Estructura del Proyecto

```
/cmd
  ├── main.go       # Archivo principal que inicia el programa
/config
  ├── ...          # Manejador de variables de entorno y configuración de la DB
/db
  ├── connection.go # Lógica para la conexión a la base de datos
/service/treeInfo
  ├── ...          # Implementación de los endpoints y llamadas a procedimientos almacenados
/types
  ├── ...          # Define estructuras de datos y tipos de retorno
/utils
  ├── ...          # Funciones auxiliares y utilitarias
```

## 📌 Configuración

Para ejecutar la API localmente, es necesario un archivo `.env` con las variables de entorno. Asegúrate de crearlo en la raíz del proyecto con el siguiente contenido:

```
PUBLIC_HOST=localhost
API_PORT=8080
DB_USER=root
DB_PWD=root
DB_NAME=ecotec
DB_ADDR=localhost:3306
```

Estas variables se manejan en `config/`, donde la API obtiene la configuración según el entorno en el que se ejecute.

---

## 🚀 Ejecución de la API

La API puede ejecutarse de dos formas:

### 🔹 1. Levantar la API Localmente
1. Asegúrate de tener **Golang** y **MySQL** instalados en tu sistema.
2. Instala las dependencias:
   ```bash
   go mod tidy
   ```
3. Crea el archivo `.env` en la raíz del proyecto con las variables de entorno (ver sección anterior).
4. Ejecuta la API con:
   ```bash
   go run cmd/main.go
   ```

La API se expondrá en el puerto configurado en `API_PORT` (por defecto `8080`).

---

### 🔹 2. Levantar la API con Docker
Si prefieres ejecutar la API en un contenedor Docker, sigue estos pasos:

#### 📌 Construcción y ejecución con Docker
1. Asegúrate de tener **Docker** instalado en tu sistema.
2. Construye la imagen del contenedor:
   ```bash
   docker build -t eco-api .
   ```
3. Ejecuta el contenedor:
   ```bash
   docker run -p 5000:5000 eco-api
   ```

🔹 **Nota**: En la ejecución con Docker, las variables de entorno se establecen en tiempo de ejecución y no se usa el archivo `.env`. Puedes definirlas con `-e` en el comando `docker run`, por ejemplo:

```bash
docker run -p 5000:5000 -e PUBLIC_HOST=localhost -e API_PORT=5000 -e DB_USER=root -e DB_PWD=root -e DB_NAME=ecotec -e DB_ADDR=localhost:3306 eco-api
```

---

## 📝 Endpoints

Los endpoints están definidos en `service/treeInfo/`, donde cada uno realiza llamadas a los procedimientos almacenados correspondientes en la base de datos.

| Método | Ruta           | Descripción                       |
|--------|---------------|-----------------------------------|
| GET    | `/trees`      | Obtiene datos de los árboles     |
| GET    | `/tree/{id}`  | Obtiene datos de un árbol por ID |
| POST   | `/tree`       | Inserta un nuevo árbol          |

*(Para más detalles, revisar `service/treeInfo/` y `db/`)*

---

## 📌 Contacto
Para cualquier duda o problema con la API, abre un *issue* en este repositorio o contacta al desarrollador principal.

---
🚀 **Happy Coding!**
