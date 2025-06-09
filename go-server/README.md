```
project-name/
│
├── cmd/                # Main applications for this project
│    └── server/        # Entry point: main.go
│
├── internal/           # Private application code
│    ├── user/          # Each domain as a separate package
│    │    ├── handler.go
│    │    ├── service.go
│    │    ├── repository.go
│    │    └── model.go
│    │
│    ├── product/       # Another domain
│    │    ├── handler.go
│    │    ├── service.go
│    │    ├── repository.go
│    │    └── model.go
│    │
│    ├── db/            # Database connection, migrations
│    └── config/        # Configuration loading (env, yaml etc.)
│
├── pkg/                # Public reusable code (utils, custom errors)
│
├── api/                # API documentation (Swagger/OpenAPI)
│
├── scripts/            # Scripts like database migrations, seeders
│
├── deployments/        # Docker, Kubernetes files
│
├── go.mod              # Go module file
└── README.md
```