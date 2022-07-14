# Backend Template Go

This repository contains a template for a Go backend with my own twist.
This will become my goto template when i'm making a backend for project.
I'm not sure if this is the best way to do it, but it works. Feel free to make a pull request and add your own twist.
This repository already contain JWT authentication as starter.

### Specs and library

- Framework : [Fiber](https://gofiber.io/)
- ORM: [GORM](https://gorm.io/)
- Database: [Postgres](https://www.postgresql.org/)
- Language: Go
- Architecture: [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- API: RESTful
- Logger: [Zerolog](https://github.com/rs/zerolog)

### Prerequisites
- Make (optional)
- Docker (optional)
- Docker Compose (optional)
- Go 1.18 or latest
- [Air](https://github.com/cosmtrek/air)

### Development Environment:
This repository mainly develop with docker and docker-compose environment for assuring the behaviour of the app itself.
You can run it locally by running `make docker.dev` or without docker with command `go run cmd/main.go`

### Commands:
- `make docker.dev`: Build and run docker-compose development environment
- `make docker.dev.build`: Force build and run docker-compose development environment
- `make docker.build`: Build docker image
- `make docker.build.alpine`: Build docker image with alpine base
- `make clean`: Clean temporary and build folder
- `make build`: Build app

### Folder structure:

```
.
├── cmd
│   └── main.go
├── config
│   └── config.go
├── docker
│   ├── alpine.Dockerfile
│   ├── dev.Dockerfile
│   ├── docker-compose.dev.yaml
│   └── Dockerfile
├── Dockerfile -> ./docker/Dockerfile
├── go.mod
├── go.sum
├── go.work
├── internal
│   ├── entities
│   │   ├── model
│   │   │   ├── token_model.go
│   │   │   └── user_model.go
│   │   └── web
│   │       └── response.go
│   ├── handler
│   │   ├── auth_handler.go
│   │   └── auth_handler_interface.go
│   ├── repository
│   │   ├── token
│   │   │   ├── token_repository.go
│   │   │   └── token_repository_interface.go
│   │   └── user
│   │       ├── user_repository.go
│   │       └── user_repository_interface.go
│   ├── routes
│   │   └── auth_route.go
│   ├── service
│   │   └── auth
│   │       ├── auth_service.go
│   │       └── auth_service_interface.go
│   └── validations
│       ├── auth_validation.go
│       └── validate.go
├── Makefile
├── pkg
│   ├── database
│   │   ├── connection.go
│   │   └── migration.go
│   ├── middleware
│   │   ├── fiber_middleware.go
│   │   └── jwt_middleware.go
│   └── utils
│       ├── password.go
│       ├── start_server.go
│       ├── token.go
│       └── uuid.go
├── README.md
├── tmp
│   ├── build-errors.log
│   └── main
└── tree

20 directories, 39 files
```

### License
This project licensed under the MIT license. See the LICENSE file for more information.