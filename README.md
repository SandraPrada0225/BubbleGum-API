# Bubblegum API

Bubblegum API es una API REST desarrollada para gestionar una plataforma de dulcería. El proyecto permite administrar productos, categorías, carritos de compra y ventas, implementando una arquitectura por capas que facilita el mantenimiento, la escalabilidad y la separación de responsabilidades.

## Objetivo

El propósito del proyecto es construir una solución backend que permita:

* Gestionar el catálogo de dulces.
* Consultar detalles de productos.
* Administrar carritos de compra.
* Procesar compras.
* Consultar historial de compras por usuario.
* Aplicar buenas prácticas de arquitectura y pruebas.

---

## Tecnologías utilizadas

* Go
* Gin
* GORM
* MySQL
* Procedimientos almacenados
* Git / Git Flow
* UML (Diagramas de secuencia)
* Testify (Pruebas unitarias)

---

## Arquitectura

El proyecto implementa una arquitectura por capas:

```text
Cliente
   ↓
Handler
   ↓
UseCase
   ↓
Provider / Repository
   ↓
Base de Datos
```

### Responsabilidades

#### Handler

* Recibe peticiones HTTP.
* Valida parámetros.
* Construye respuestas.

#### UseCase

* Contiene la lógica de negocio.
* Coordina operaciones del sistema.

#### Provider / Repository

* Gestiona acceso a datos.
* Ejecuta consultas y procedimientos almacenados.

#### Base de datos

* Almacena y administra la información del sistema.

---

## Funcionalidades implementadas

### Dulces

* Obtener detalle de un dulce por código.
* Consultar categorías asociadas.

### Carritos

* Consultar carrito.
* Actualizar carrito.
* Obtener detalle del carrito.
* Comprar carrito.

### Compras

* Registrar ventas.
* Consultar historial de compras por usuario.

---

## 🗄 Base de datos

La aplicación utiliza MySQL con tablas relacionadas mediante llaves foráneas:

* Usuarios
* Dulces
* Categorías
* Marcas
* Carritos
* Ventas
* Medios de pago
* Estados del carrito

---

## Diagramas UML

La documentación UML se encuentra en:

```text
docs/uml/
```

Incluye diagramas de secuencia para los principales endpoints implementados.

---

## Cómo ejecutar el proyecto

Clonar repositorio:

```bash
git clone <url-del-repositorio>
```

Entrar al proyecto:

```bash
cd bubblegum-api
```

Instalar dependencias:

```bash
go mod tidy
```

Ejecutar servidor:

```bash
go run cmd/server/main.go
```

La API estará disponible en:

```text
http://localhost:8080
```
