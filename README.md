# Go Microservices con Traefik y Consul

Arquitectura de microservicios en Go con API Gateway (Traefik), Service Discovery (Consul) y PostgreSQL.

## 🏗️ Arquitectura

```
┌─────────────────────────────────────────────────────────────┐
│                         Cliente                             │
└────────────────────────┬────────────────────────────────────┘
                         │
                         ▼
┌─────────────────────────────────────────────────────────────┐
│                    Traefik (Puerto 80)                      │
│                      API Gateway                            │
└────────┬───────────────────────────────────┬────────────────┘
         │                                   │
         ▼                                   ▼
┌────────────────────┐            ┌─────────────────────────┐
│  oferta-service    │            │ especificacion-service  │
│    (Puerto 8082)   │            │     (Puerto 8081)       │
└────────┬───────────┘            └──────────┬──────────────┘
         │                                   │
         ▼                                   ▼
┌────────────────────┐            ┌─────────────────────────┐
│  PostgreSQL        │            │    PostgreSQL           │
│  (Puerto 5432)     │            │    (Puerto 5433)        │
└────────────────────┘            └─────────────────────────┘
         │                                   │
         └───────────────┬───────────────────┘
                         ▼
                ┌────────────────────┐
                │  Consul (8500)     │
                │ Service Discovery  │
                └────────────────────┘
```

## 🚀 Servicios

### 1. **oferta-service** (Puerto 8082)
Gestión de ofertas con CRUD completo.

### 2. **especificacion-service** (Puerto 8081)
Gestión de especificaciones con CRUD parcial según especificaciones de lo solicitado.

### 3. **Traefik** (Puerto 80)
API Gateway que enruta el tráfico a los microservicios.

### 4. **Consul** (Puerto 8500)
Service Discovery para registro y descubrimiento de servicios.

## 📋 Requisitos

- Docker
- Docker Compose
- Go 1.21+ (para desarrollo local)

## 🛠️ Instalación y Ejecución

### 1. Clonar el repositorio
```bash
git clone https://github.com/GustavoAuger/semana2.git
cd Go-Micro
```

### 2. Levantar todos los servicios
```bash
docker-compose up -d
```

### 3. Verificar que los servicios estén corriendo
```bash
docker-compose ps
```

### 4. Ver logs
```bash
# Ver logs de todos los servicios
docker-compose logs -f

# Ver logs de un servicio específico
docker-compose logs -f oferta-service
docker-compose logs -f especificacion-service
docker-compose logs -f traefik
```

## 🌐 Endpoints

### A través del API Gateway (Puerto 80) - **RECOMENDADO**

#### Oferta Service
```bash
# Health check
curl http://localhost/api/v1/ofertas/health

# Listar todas las ofertas
curl http://localhost/api/v1/ofertas

# Obtener una oferta por ID
curl http://localhost/api/v1/ofertas/1

# Crear una nueva oferta
curl -X POST http://localhost/api/v1/ofertas \
  -H "Content-Type: application/json" \
  -d '{

  }'

# Actualizar una oferta
curl -X PUT http://localhost/api/v1/ofertas/1 \
  -H "Content-Type: application/json" \
  -d '{

  }'

# Eliminar una oferta
curl -X DELETE http://localhost/api/v1/ofertas/1
```

#### Especificacion Service
```bash
# Health check
curl http://localhost/api/v1/especificaciones/health

# Listar todas las especificaciones
curl http://localhost/api/v1/especificaciones

# Obtener una especificación por ID
curl http://localhost/api/v1/especificaciones/1

# Crear una nueva especificación
curl -X POST http://localhost/api/v1/especificaciones \
  -H "Content-Type: application/json" \
  -d '{

  }'
```

### Acceso Directo a los Servicios (Solo para desarrollo)

```bash
# oferta-service (Puerto 8082)
curl http://localhost:8082/api/v1/health

# especificacion-service (Puerto 8081)
curl http://localhost:8081/api/v1/health
```

## 🎛️ Dashboards

### Traefik Dashboard
```
http://localhost:8080/dashboard/
```

### Consul UI
```
http://localhost:8500
```

## 🗄️ Base de Datos

### Conexión a PostgreSQL

#### Base de datos de ofertas
```bash
docker exec -it db_oferta psql -U postgres -d ofertas_db
```

#### Base de datos de especificaciones
```bash
docker exec -it db_especificacion psql -U postgres -d especificaciones_db
```

## 🔧 Desarrollo

### Estructura del Proyecto

```
Go-Micro/
├── oferta-service/
│   ├── cmd/
│   │   └── server/
│   │       └── main.go
│   ├── internal/
│   │   ├── handler/
│   │   ├── service/
│   │   ├── repository/
│   │   ├── model/
│   │   └── consul/
│   ├── Dockerfile
│   └── go.mod
├── especificacion-service/
│   ├── cmd/
│   │   └── server/
│   │       └── main.go
│   ├── internal/
│   │   ├── handler/
│   │   ├── service/
│   │   ├── repository/
│   │   ├── model/
│   │   └── consul/
│   ├── Dockerfile
│   └── go.mod
├── docker-compose.yml
├── traefik.yml
├── dynamic_conf.yml
└── README.md
```

### Reconstruir un servicio específico

```bash
# Reconstruir oferta-service
docker-compose up -d --build oferta-service

# Reconstruir especificacion-service
docker-compose up -d --build especificacion-service
```

### Reiniciar un servicio

```bash
docker-compose restart oferta-service
docker-compose restart especificacion-service
```

## 🧹 Limpieza

### Detener todos los servicios
```bash
docker-compose down
```

### Detener y eliminar volúmenes (⚠️ Elimina los datos de la BD)
```bash
docker-compose down -v
```

### Limpiar imágenes y caché de Docker
```bash
docker system prune -a -f --volumes
```

## 🐛 Troubleshooting

### Ver logs de un servicio
```bash
docker logs oferta-service
docker logs especificacion-service
docker logs traefik
docker logs consul
```

### Verificar la red de Docker
```bash
docker network inspect go-micro_app-network
```

### Verificar servicios registrados en Consul
```bash
curl http://localhost:8500/v1/agent/services
```

### Verificar rutas en Traefik
```bash
curl http://localhost:8080/api/http/routers
```

## 📝 Variables de Entorno

### oferta-service
- `DB_HOST`: Host de la base de datos (default: `db_oferta`)
- `DB_PORT`: Puerto de la base de datos (default: `5432`)
- `DB_USER`: Usuario de la base de datos (default: `postgres`)
- `DB_PASSWORD`: Contraseña de la base de datos (default: `postgres`)
- `DB_NAME`: Nombre de la base de datos (default: `ofertas_db`)
- `CONSUL_HTTP_ADDR`: Dirección de Consul (default: `consul:8500`)
- `SERVICE_NAME`: Nombre del servicio (default: `oferta-service`)
- `SERVICE_PORT`: Puerto del servicio (default: `8082`)

### especificacion-service
- `DB_HOST`: Host de la base de datos (default: `db_especificacion`)
- `DB_PORT`: Puerto de la base de datos (default: `5432`)
- `DB_USER`: Usuario de la base de datos (default: `postgres`)
- `DB_PASSWORD`: Contraseña de la base de datos (default: `postgres`)
- `DB_NAME`: Nombre de la base de datos (default: `especificaciones_db`)
- `CONSUL_HTTP_ADDR`: Dirección de Consul (default: `consul:8500`)
- `SERVICE_NAME`: Nombre del servicio (default: `especificacion-service`)
- `SERVICE_PORT`: Puerto del servicio (default: `8081`)

## 🔒 Seguridad

- En producción, NO expongas los puertos de los servicios directamente (8081, 8082)
- Usa HTTPS con certificados SSL/TLS
- Implementa autenticación JWT
- Configura rate limiting en Traefik
- Usa variables de entorno para credenciales sensibles

## 📚 Tecnologías Utilizadas

- **Go 1.21+**: Lenguaje de programación
- **Gin**: Framework web para Go
- **GORM**: ORM para Go
- **PostgreSQL**: Base de datos relacional
- **Traefik v2.10**: API Gateway y reverse proxy
- **Consul 1.15**: Service Discovery
- **Docker & Docker Compose**: Containerización

## 👥 Autor

Gustavo Auger
