# Go REST API — Clean Architecture

A RESTful API built with Go as a learning project, exploring how Clean Architecture principles translate from Java/Spring Boot to a non-object-oriented language.

## Stack

- **Go** — main language
- **Gin** — HTTP framework
- **MySQL** — database
- **Docker** — containerization
- **bcrypt** — password hashing

---

## Architecture

The project follows Clean Architecture, with a strict dependency rule: outer layers know about inner layers, but never the other way around.

```
cmd/api/
└── main.go                         # Entry point — wires all dependencies manually

internal/
├── adapter/
│   ├── handler/                    # HTTP layer (equivalent to @Controller)
│   │   └── user_handler.go
│   └── repositories/               # DB implementation (equivalent to @Repository)
│       └── mysql_user_repo.go
│
├── domain/
│   ├── entities/                   # Core business objects
│   │   └── user.go
│   └── repositories/               # Repository interfaces (dependency inversion)
│       └── user_repository.go
│
├── infrastructure/
│   ├── config/                     # Environment config
│   ├── database/                   # MySQL connection
│   └── server/                     # Gin router setup
│
└── usecases/
    └── user/                       # One file per use case
        ├── create_user.go
        ├── update_user.go
        ├── change_password.go
        ├── delete_user.go
        ├── get_user.go
        └── get_users.go

pkg/
├── apperror/                       # Centralized error handling
└── validator/                      # Validation error formatting
```

### Dependency flow

```
HTTP Request → Handler → Use Case → Repository Interface → Repository Implementation → MySQL
```

The domain layer (entities + repository interfaces) has zero dependencies on frameworks, databases, or HTTP — it is pure Go.

---

## Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/api/v1/users` | List all users |
| `GET` | `/api/v1/users/:id` | Get user by ID |
| `POST` | `/api/v1/users` | Create user |
| `PATCH` | `/api/v1/users/:id` | Update name and/or email |
| `PATCH` | `/api/v1/users/:id/change-password` | Change password |
| `DELETE` | `/api/v1/users/:id` | Delete user |

---

## Key Concepts Explored

### Pointers and memory
Unlike Java where all objects are automatically references, Go requires explicit choice. `&User{}` creates a pointer (equivalent to `new User()` in Java) while `User{}` creates a real copy in memory. Understanding this was one of the biggest learnings — it makes explicit what happens when data is passed between functions.

### Error handling without exceptions
Go has no `try/catch`. Errors are plain values returned by functions:
```go
user, err := repo.FindById(ctx, id)
if err != nil {
    return nil, err
}
```
Each layer propagates errors upward until the handler translates them into HTTP responses.

### Implicit interfaces
No `implements` keyword. If a struct has the methods defined in an interface, it satisfies it automatically. This is how `MySQLUserRepository` satisfies `UserRepository` without ever declaring it.

### Manual dependency injection
No `@Autowired` or framework magic. Everything is wired explicitly in `main.go`:
```go
userRepo     := repositories.NewMySQLUserRepository(db)
createUserUC := userusecase.NewCreateUserUseCase(userRepo)
userHandler  := handler.NewUserHandler(createUserUC, ...)
```

---

## Getting Started

### Prerequisites
- Go 1.21+
- Docker

### Setup

1. Clone the repository
```bash
git clone https://github.com/your-username/your-repo.git
cd your-repo
```

2. Create your `.env` file based on the example
```bash
cp .env.example .env
```

3. Fill in the environment variables
```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=go_api
SERVER_PORT=8080
```

4. Start the database
```bash
docker-compose up -d
```

5. Run the application
```bash
go run ./myapp/cmd/api/main.go
```

---

## Error Response Format

All errors follow a consistent structure:

```json
{
  "timestamp": "2026-04-26T16:31:53Z",
  "code": 404,
  "status": "NOT_FOUND",
  "message": "user not found"
}
```

Validation errors include field-level details:

```json
{
    "timestamp": "2026-04-27T10:46:58.1752643-03:00",
    "code": 400,
    "status": "VALIDATION",
    "message": "validation failed",
    "errors": [
        {
            "field": "Email",
            "message": "Email must be a valid email"
        },
        {
            "field": "Password",
            "message": "Password must be at least 8 characters"
        }
    ]
}
```

---

## What I Learned

This project was built because I wanted to learn Go, to test myself on a new environment and to understand Clean Architecture without "framework magic". The main takeaways were how Go handles memory explicitly through pointers, how error propagation works without exceptions, and how the same architectural principles apply regardless of the language or framework.
