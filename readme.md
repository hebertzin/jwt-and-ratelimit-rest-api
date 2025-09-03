# jwt-and-ratelimit-rest-api

A RESTful API built with Go, implementing authentication with JWT and request rate limiting following industry best practices.

## Technology

* [Go](https://golang.org/) 1.22
* [PostgreSQL](https://www.postgresql.org/)
* [Docker](https://www.docker.com/)
* [Chi](https://go-chi.io/#/)

## Project structure

```
.
├── .github/            # Actions, files of github
├── build/              # Docker, docker-compose
├── cmd/                # App Entrypoint
│   └── main.go         # HTTP Server Startup
├── docs/               # App documentations
├── migrations/         # Database migrations
├── packages/           # Packages for the application to run
│   ├── domain/         # Entities and interfaces
│   ├── handler/        # HTTP requests
│   ├── infra/          # Technical implementations (DB, hash, jwt, etc.)
│   ├── middlewares/    # Application middlewares
│   ├── router/         # Application route
│   ├── services/       # Application logic
│   └── utils/          # Validations, helpers, etc.
└── tools/              # Tools to automate some application processes
```

## Requirements

*  Go 1.22+
*  Docker and Docker Compose
*  Make (optional, but recommended)

## Instalation

Clone repository:

```bash
git clone https://github.com/hebertzin/jwt-and-ratelimit-rest-api.git
cd jwt-and-ratelimit-rest-api
```

After setting the environment variables as shown in the .env.example file, run the application locally:

```bash
make run
```

Running on: [http://localhost:8080](http://localhost:8080)

## API Documentation

After running the project, the Swagger documentation will be available at:  [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

## Architecture

* Communication via **REST API**
* Persistence in **PostgreSQL**
* Authentication via **JWT**

## Licence

Distributed under the MIT License. See `LICENSE` for more information.
