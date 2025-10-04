# 🛒 Digital Shop — Backend (Go + Fiber + PostgreSQL + Redis)

**Digital Shop** is a secure and scalable e-commerce backend built with **Go (Fiber)**.  
It follows **Clean Architecture**, uses **PostgreSQL** for data persistence, **Redis** for caching, and **JWT** for authentication.

Built for performance, maintainability, and real-world deployment.

---

## 🚀 Features

- 🔐 JWT Authentication (Access & Refresh Tokens)
- 🧩 Role-Based Access (admin / user)
- 🧼 XSS Protection using bluemonday
- 🧱 Clean Architecture (Domain → Usecase → Repository → Interface)
- 🧠 Input Validation with validator
- 🧰 Secure Password Hashing (bcrypt)
- 🐘 PostgreSQL + 🧊 Redis integration
- 🐳 Docker support for local and production environments

---

## 🧭 Project Structure
```bash
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
```

---

## ⚙️ Environment Variables

| Variable | Description |
|-----------|-------------|
| `POSTGRES_DB` | Database name |
| `POSTGRES_PORT` | Database port (default: 5432) |
| `POSTGRES_HOST` | Database host (`localhost` or `db`) |
| `POSTGRES_USER` | Database username |
| `POSTGRES_PASSWORD` | Database password |
| `ACCESS_TOKEN` | Secret key for signing Access Tokens |
| `REFRESH_TOKEN` | Secret key for signing Refresh Tokens |
| `USERNAME_ADMIN` | Default admin username |
| `PASSWORD_ADMIN` | Default admin password |
| `EMAIL_ADMIN` | Default admin email |
| `REDIS_ADDR` | Redis address (e.g. `redis:6379`) |
| `REDIS_PASSWORD` | Redis password (if set) |

### 🧩 Example `.env`

```env
POSTGRES_DB=digital_shop
POSTGRES_PORT=5432
POSTGRES_HOST=db
POSTGRES_USER=admin
POSTGRES_PASSWORD=secret

ACCESS_TOKEN=your_access_secret
REFRESH_TOKEN=your_refresh_secret

USERNAME_ADMIN=admin
PASSWORD_ADMIN=supersecret
EMAIL_ADMIN=admin@shop.com

REDIS_ADDR=redis:6379
REDIS_PASSWORD=
🐳 Run with Docker
1️⃣ Build and start services
docker-compose up --build
2️⃣ Run database migrations
go run ./pkg/runMigrations.go
3️⃣ Access the API
http://localhost:8080
🔑 Authentication
Token Type	Lifetime	Description
Access Token	⏱ 5 minutes	Used for authorized API calls
Refresh Token	🕒 15 days	Used to renew access tokens

Each token contains:

exp — Expiration timestamp

iat — Issued-at timestamp

iss — Issuer (digital-shop)

📡 API Overview
Endpoint	Method	Description	Auth
/register	POST	Register a new user	❌
/login	POST	Login and receive tokens	❌
/products	GET	List all products	✅
/products/:id	GET	Product details	✅

🧱 Tech Stack
Layer	Technology
Language	Go 1.23+
Web Framework	Fiber
Database	PostgreSQL
Cache	Redis
Auth	JWT
Container	Docker
Security	bluemonday, bcrypt

🧹 Future Plans
⚛️ Frontend with React or Next.js

🛍️ Add product categories and cart system

🧪 Unit and integration tests

📘 Swagger API Documentation

⚙️ CI/CD pipeline (GitHub Actions)

🧠 Architecture Summary
Digital Shop follows Clean Architecture:

Domain — Core entities and business logic

Usecase — Application-level logic (use cases)

Repository — Data access layer (Postgres / Redis)

Interface — HTTP layer (Fiber handlers, middleware, utils)

🪪 License
MIT License
Copyright (c) 2025
See LICENSE for more details.

💬 Author
Amir 2002
📧 amir2002-js
🧠 Backend Developer — Go / Fiber / PostgreSQL / Docker

