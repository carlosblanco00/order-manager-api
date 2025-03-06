# order-manager-api


## Descripción

El objetivo es construir un sistema básico de gestión de pedidos utilizando **Golang**, **MySQL** para el almacenamiento de productos y pedidos, y **Redis** para implementar idempotencia en uno de los endpoints, ademas se trabajo bajo Clean Architecture.

### Requisitos Funcionales

- **GET /products**: Devuelve una lista de todos los productos almacenados en la base de datos.
- **POST /orders**: Permite crear un pedido, validando el stock, calculando el total y reduciendo el stock de los productos.
- **GET /orders/:id**: Devuelve los detalles de un pedido específico, incluidas las líneas de pedido.
- **PUT /products/:id/stock**: Permite actualizar el stock de un producto.

## Instrucciones para Ejecutar el Proyecto

Sigue estos pasos para ejecutar el proyecto en tu entorno local:

### 1. Clona el Repositorio

Primero, clona este repositorio en tu máquina local usando Git:

```bash
git clone https://github.com/tu-usuario/order-manager-api.git
cd order-manager-api
```
### 3. Levanta los Contenedores con Docker Compose

Levanta los contenedores necesarios ejecutando el siguiente comando:

```bash
docker-compose up --build
```
Este comando realiza lo siguiente:

Construye las imágenes necesarias.
Levanta los contenedores de MySQL, Redis, y la aplicación Go.
Expondrá la aplicación en el puerto 8080.

### 4. Accede a los Endpoints
Una vez que los contenedores estén corriendo, puedes acceder a los siguientes endpoints:

- GET /products: Devuelve una lista de todos los productos almacenados en la base de datos.
- POST /orders: Crea un nuevo pedido (con validación de stock, cálculo del total y reducción del stock). 
- GET /orders/:id: Obtiene los detalles de un pedido específico, incluidos los artículos del pedido.
- PUT /products/:id/stock: Actualiza el stock de un producto.

### Requisitos Técnicos

- **Base de datos**: MySQL para gestionar las tablas de productos, pedidos y detalles.
- **Idempotencia**: Redis se utiliza para manejar las claves de idempotencia en el endpoint de creación de pedidos.
- **Docker**: Docker Compose para levantar los servicios de MySQL, Redis y la aplicación Go.



