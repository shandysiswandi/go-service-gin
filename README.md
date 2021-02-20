# Go Service Gin
go started template for building microservice using framework gin

---

## Architecture
**Domain-Driven-Design** is the concept that the structure and language of software code should match the business domain. <br><br>
DDD has 4 layers in the architecture:

- **Interface**: This layer responsibles for the interaction with user, whether software presents information or recieves information from user.
- **Application**: This is a thin layer between interface and domain, it could call domain services to serve the application purposes.
- **Domain**: The heart of the software, this layer holds domain logic and business knowledge.
- **Infrastructure**: A supporting layer for the other layers. This layer contains supporting libraries or external services like database or UI supporting library.

---

## Project Structure
```
.
├── application
│   └── blog.go
├── cmd
│   ├── console
│   └── rest
│       └── main.go
├── config
│   ├── app.go
│   ├── app_test.go
│   ├── database.go
│   ├── database_test.go
│   ├── redis.go
│   ├── redis_test.go
│   ├── sentry.go
│   └── sentry_test.go
├── domain
│   └── blogs
│       ├── entity.go
│       ├── entity_test.go
│       ├── event.go
│       ├── payload.go
│       └── repository.go
├── infrastructure
│   ├── database
│   │   ├── database.go
│   │   ├── mysql
│   │   │   └── blog.go
│   │   └── postgres
│   ├── external
│   │   └── jsonplaceholder
│   └── library
│       ├── redis
│       │   ├── redis.go
│       │   └── redis_test.go
│       └── sentry
│           ├── sentry.go
│           └── sentry_test.go
├── interfaces
│   ├── console
│   └── rest
│       ├── blog.go
│       └── home.go
├── util
│    ├── bcrypt
│    │   ├── bcrypt.go
│    │   └── bcrypt_test.go
│    ├── logger
│    │   ├── logger.go
│    │   └── logger_test.go
│    ├── numbers
│    │   ├── random.go
│    │   └── random_test.go
│    └── stringy
│        ├── snackcase.go
│        └── snackcase_test.go
├── .env.example
├── .gitignore
├── go.mod
├── go.sum
├── LICENSE
├── Makefile
└── README.md
```