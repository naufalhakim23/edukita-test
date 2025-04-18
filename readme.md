# Edukita LMS API

A Learning Management System API built with Go, providing functionality for course management, assignments, and submissions.

## Prerequisites

- [Go](https://golang.org/dl/) (version 1.24 or higher)
- PostgreSQL database
- Docker & Docker Compose (optional, for containerized setup)

## Setup Instructions

### Option 1: Local Setup

#### 1. Install Go

Download and install Go from the [official website](https://golang.org/dl/).

#### 2. Clone the repository

```bash
git clone https://github.com/yourusername/edukita-lms.git
cd edukita-lms
```

#### 3. Install dependencies

```bash
go mod tidy
```

#### 4. Set up the database

Create a PostgreSQL database named `edukita_lms`:

```bash
psql -U postgres -c "CREATE DATABASE edukita_lms"
```

#### 5. Configure environment variables

Create a `.env` file in the root directory using the provided template:

```bash
cp .env.example .env
```

The `.env.example` file contains the following configuration:

```
APP_NAME="edukita-teaching-grading"
APP_ENV="local"
APP_PORT="8081"
APP_SECRET="supersecretsecret"
APP_STATIC_TOKEN="supersecretsecret"
APP_SWAGGER_PATH=""
POSTGRES_NAME="edukita-teaching-grading"
POSTGRES_URL="postgres://postgres:postgresql@localhost:5432/edukita_lms?sslmode=disable"
```

Adjust the values as needed for your environment.

#### 6. Run database migrations

Apply the database migrations to set up the required tables:

```bash
make migration-up
```

#### 7. Start the application

```bash
make start
```

The API will be available at `http://localhost:8081`.

### Option 2: Docker Setup

#### 1. Clone the repository

```bash
git clone https://github.com/yourusername/edukita-lms.git
cd edukita-lms
```

#### 2. Configure environment variables

Create a `.env` file in the root directory using the provided template:

```bash
cp .env.example .env
```

Update the PostgreSQL connection string in the `.env` file to use the Docker service name:

```
POSTGRES_URL="postgres://postgres:postgresql@postgres:5432/edukita_lms?sslmode=disable"
```

#### 3. Build and run with Docker Compose

```bash
docker-compose up -d
```

This will:
- Start a PostgreSQL container
- Build and start the LMS API container
- Run the necessary migrations automatically

The API will be available at `http://localhost:8081`.

#### 4. Stop Docker containers

```bash
docker-compose down
```

To remove volumes and all data:

```bash
docker-compose down -v
```

## Dockerfile

For reference, here's the Dockerfile structure used in this project:

```dockerfile
################################################################################
# BASE
################################################################################
FROM golang:1.24-alpine AS base

RUN apk add --no-cache git openssh-client bash curl gcc g++ make libc6-compat git openssh-client ca-certificates vips vips-dev libc-dev libheif

################################################################################
# DEPENDENCY
################################################################################
FROM base AS dependency

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download -x && go mod verify

################################################################################
# BUILDER
################################################################################
FROM dependency AS builder

COPY . .

# build swagger file
RUN go install github.com/swaggo/swag/cmd/swag@v1.8.0
RUN swag init

RUN CGO_ENABLED=0 GOOS=linux go build -a --ldflags '-extldflags "static"' -tags "netgo" -installsuffix netgo -o edukita-teaching-grading -v

################################################################################
# WORKER
################################################################################
FROM alpine:latest

WORKDIR /usr/local/bin

RUN apk --update add ca-certificates

# copy docs (swagger API docs)
COPY --from=builder /build/docs /usr/local/bin/docs

COPY --from=builder /build/migrations  /usr/local/bin/migrations

COPY --from=builder /build/edukita-teaching-grading /usr/local/bin/edukita-teaching-grading

CMD ["./edukita-teaching-grading"]

```

## API Endpoints

### User Management

| Method | Endpoint | Description | Authentication |
|--------|----------|-------------|---------------|
| POST | `/api/v1/user/register` | Register a new user | No |
| POST | `/api/v1/user/login` | User login | No |
| POST | `/api/v1/user/logout` | User logout | Yes |
| GET | `/api/v1/user/me` | Get current user details | Yes |
| GET | `/api/v1/user/:id` | Get user by ID | Yes |

### Course Management

| Method | Endpoint | Description | Authentication |
|--------|----------|-------------|---------------|
| POST | `/api/v1/lms/courses` | Create a new course | Yes |
| GET | `/api/v1/lms/courses/:id` | Get course by ID | Yes |
| GET | `/api/v1/lms/courses/:code` | Get course by code | Yes |
| GET | `/api/v1/lms/courses` | Get all courses | Yes |
| PUT | `/api/v1/lms/courses/:id` | Update course by ID | Yes |

### Assignment Management

| Method | Endpoint | Description | Authentication |
|--------|----------|-------------|---------------|
| POST | `/api/v1/lms/assignments` | Create a new assignment | Yes |
| GET | `/api/v1/lms/assignments/:id` | Get assignment by ID | Yes |
| PUT | `/api/v1/lms/assignments/:id` | Update assignment by ID | Yes |

### Submission Management

| Method | Endpoint | Description | Authentication |
|--------|----------|-------------|---------------|
| POST | `/api/v1/lms/submissions` | Submit an assignment | Yes |
| GET | `/api/v1/lms/submissions/:id` | Get submission by ID | Yes |
| PUT | `/api/v1/lms/submissions/:id` | Update submission by ID | Yes |
| GET | `/api/v1/lms/submissions/course/:id` | Get all submissions for a course | Yes |
| GET | `/api/v1/lms/submissions/assignments/:id` | Get all submissions for an assignment | Yes |
| GET | `/api/v1/lms/submissions/users/:id` | Get all submissions by a user | Yes |

## Authentication

Most endpoints require authentication. Include the JWT token in the Authorization header:

```
Authorization: Bearer <your_token>
```

You can obtain a token by using the login endpoint.

## Development

### Available Make Commands

- `make start`: Start the application
- `make migration-up`: Apply database migrations
- `make migration-down`: Revert database migrations

## License

[MIT](LICENSE)
