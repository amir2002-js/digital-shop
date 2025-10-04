🏪 Digital Shop — Backend (Go + Fiber + PostgreSQL + Redis + Docker)

High-performance e-commerce backend built with Go (Fiber), PostgreSQL, and Redis, designed for scalability, clean architecture, and production-grade deployment via Docker Compose.

🧠 Tech Stack
Layer	Technology
Language	Go 1.22+
Framework	Fiber v2
Database	PostgreSQL 14+
Cache	Redis 7+
ORM	GORM
Validation	go-playground/validator
Security	JWT, bcrypt, bluemonday
Containerization	Docker & Docker Compose
⚙️ Architecture Overview

Clean Architecture (a.k.a. Hexagonal) — separation of concerns between:

Domain → Core business entities.

Usecase → Application logic.

Service → Business-level operations.

Repository → Data access layer (PostgreSQL & Redis).

Interface/HTTP → API layer (Fiber).

🗂️ Project Structure
.
├── cmd/
│   └── main.go                      # Entry point
│
├── db/
│   └── migrations/                  # SQL migrations
│       ├── 000001_create_users.up.sql
│       └── 000001_create_users.down.sql
│
├── internal/
│   ├── domain/                      # Entities (core models)
│   │   ├── products/
│   │   └── users/
│   │
│   ├── repository/                  # PostgreSQL & Redis repositories
│   │   ├── cache/
│   │   └── postgres/
│   │
│   ├── services/                    # Business services
│   │   ├── cache/
│   │   ├── products/
│   │   └── users/
│   │
│   ├── usecase/                     # Application usecases
│   │   ├── cache.go
│   │   └── userUseCase.go
│   │
│   └── interface/
│       └── http/                    # HTTP layer
│           ├── handler/
│           │   ├── products/
│           │   └── user/
│           ├── middleware/
│           └── util/
│               ├── jwtToken/
│               ├── password/
│               ├── returnsHandler/
│               └── whoIs/
│
├── pkg/                             # Shared utilities (dsn, migrations)
│
├── Dockerfile
├── docker-compose.yml
├── .env
├── go.mod
├── go.sum
├── LICENSE
└── README.md

🔐 Environment Variables

All configuration is handled via .env file:

# Database
POSTGRES_DB=digital_shop
POSTGRES_PORT=5432
POSTGRES_HOST=db
POSTGRES_USER=admin
POSTGRES_PASSWORD=secret

# JWT
REFRESH_TOKEN=supersecret_refresh
ACCESS_TOKEN=supersecret_access

# Admin Account (auto-detected)
USERNAME_ADMIN=admin
PASSWORD_ADMIN=admin123
EMAIL_ADMIN=admin@digital-shop.com

# Redis
REDIS_ADDR=redis:6379
REDIS_PASSWORD=redispass

🧩 Features

✅ User Authentication

Register / Login with validation

Password hashing with bcrypt

Role-based access (user, admin)

JWT Access & Refresh tokens

Input sanitization via bluemonday

✅ Products (planned)

CRUD operations

Search, pagination, filters

✅ Caching

Redis-based caching layer for performance boost

✅ Docker Ready

One command deployment with docker compose up

✅ Clean Code

No spaghetti. Fully layered, testable, and extensible.

🚀 Getting Started
1️⃣ Clone the repo
git clone https://github.com/<your_username>/digital-shop-backend.git
cd digital-shop-backend

2️⃣ Configure .env

Copy .env.example or create your own .env file (see above).

3️⃣ Start via Docker
docker compose up --build


This will start:

db (PostgreSQL)

redis (Redis cache)

app (Go Fiber API)

4️⃣ Run migrations
go run ./pkg/runMigrations.go

🧠 JWT Structure Example

Access Token (valid 5 minutes)
Refresh Token (valid 15 days)

Claims:

{
  "user_id": 123,
  "role": "admin",
  "exp": 1736012452,
  "iat": 1736012152,
  "iss": "digital-shop"
}

🧪 API Example
Register User

POST /api/v1/users/register

{
  "username": "amirreza",
  "email": "amir@example.com",
  "password": "strongpassword",
  "confirm_password": "strongpassword"
}


✅ Response:

{
  "data": {
    "id": 1,
    "username": "amirreza",
    "email": "amir@example.com",
    "role": "user"
  },
  "access_token": "<JWT_ACCESS_TOKEN>",
  "refresh_token": "<JWT_REFRESH_TOKEN>"
}

🧱 Future Plans

 Full product module

 Order & payment system

 Role-based middlewares

 Logging & metrics

 Frontend (React or Next.js)

🧑‍💻 Author

Amirreza — Digital Shop Backend

🌍 github.com/amir2002-js