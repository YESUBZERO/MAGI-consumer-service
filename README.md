# 🚀 MAGI Consumer Service

**MAGI Consumer Service** es un microservicio desarrollado en **Golang** que consume mensajes de **Apache Kafka**, procesa datos AIS y los almacena en **PostgreSQL**.

---

## 📌 **Arquitectura del Servicio**
Este servicio forma parte del ecosistema de **MAGI Backend** y tiene la función de:
- 📥 **Consumir datos AIS** desde Kafka (`static-message` y `enriched-message`).
- 🔄 **Determinar si un barco ya está en la base de datos**.
- 📡 **Publicar mensajes en `scrape-message` si faltan datos**.
- 💾 **Guardar barcos enriquecidos en PostgreSQL** cuando llegan desde `enriched-message`.

📍 **Flujo de Datos en `MAGI Consumer Service`**:
```
            +-------------------------------------+
            |       Apache Kafka                 |
            | - static-message                   |
            | - enriched-message                 |
            | - scrape-message                   |
            +-------------------------------------+
                       |
                       v
             +----------------+
             | Consumer Service|
             | (Golang)        |
             +--------+--------+
                       |
                       v
            +-------------------------------------+
            |         PostgreSQL                 |
            | (Almacenamiento de barcos)        |
            +-------------------------------------+
```
---

## 🚀 **Instalación y Configuración**

### 🛠️ **1️⃣ Requisitos Previos**
Antes de iniciar, asegúrate de tener:
- Docker & Docker Compose
- Golang `>= 1.18`
- PostgreSQL `>= 13`
- Apache Kafka `>= 2.8`

### 🐳 **2️⃣ Instalación con Docker**
Este servicio se ejecuta con **Docker Compose**. Clona el repositorio y ejecuta:

```sh
git clone https://github.com/YESUBZERO/MAGI-consumer-service.git
cd MAGI-consumer-service
docker-compose up --build
```

---

## Estructura del proyecto
```
MAGI-consumer-service/
│── internal/
│   ├── config/         # Configuración del servicio (Kafka, PostgreSQL)
│   ├── kafka/          # Lógica de consumo y producción de Kafka
│   ├── models/         # Definición de estructuras de datos (Ship, AISMessage)
│   ├── repository/     # Lógica para almacenar en PostgreSQL
│── main.go             # Punto de entrada del servicio
│── Dockerfile          # Definición del contenedor de Docker
│── docker-compose.yml  # Orquestación de servicios
│── README.md           # Documentación
```

## Configuracion de Variables de Entorno
```
# Configuración de PostgreSQL
POSTGRES_USER="TU USUARIO"
POSTGRES_PASSWORD="TU PASSWORD"
POSTGRES_DB=magi
DATABASE_DSN=postgres:"TU DIRECCION A LA BASE DE DATOS POSTGRES SQL"

# Configuración de Kafka
KAFKA_BROKERS="DIRECCION DEL SERVIDOR KAFKA":"PUERTO"
KAFKA_STATIC_TOPIC=static-message
KAFKA_ENRICHED_TOPIC=enriched-message
KAFKA_SCRAPE_TOPIC=scrape-message
```
