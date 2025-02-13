# Simple Account Recovery Docs

Golang REST API With gin framework and database postgreSQL

## [Go-Account-Recovery](https://github.com/xprasetio/go-account-recovery)

## Table of Contents

- [Getting Started](#getting-started)
- [Structures](#structures)
- [Features](#features)
  - [API Response](#api-endpoint)
  - [Sign Up response](#signup-response)
  - [Login response](#login-response)
  - [Recovery response](#recovery-response)
  - [Verify response](#verify-response)
  - [Reset Password response](#reset-password-response)
  - [Login After Reset Password response](#login-after-reset-password-response)
- [Logging Logrus ](#log-error)
  - [Email Not Found ](#email-not-found)
  - [Email Or Password Not Match ](#password-or-email-not-match)
  - [Email Not Found ](#email-not-found)
  - [Wrong Code Recovery ](#wrong-code-recovery)
  - [Invalid Token ](#invalid-token)
- [Database Schema](#database-schema)
- [Credits](#credits)
- [Copyright](#copyright)

## Getting Started

#### Requirements

- Database: `Postgres`
- Docker
- Viper
- Gin
- Go v1.23.x

#### Install & Run

Download this project:

```shell script
git clone https://github.com/xprasetio/go-account-recovery.git
```

Download project dependencies:

```shell script
go mod download && go mod tidy
```

Before run this project, you should set configs with yours.
Create & configure your `config.yaml` or with my setting config smtp with gmail smtp

Fast run with:

```shell script
make run

# running on default port 9999
```

Build with docker for the database `postgreeSQL`

```shell script
docker-compose up -d
```

make sure to double check in dbeaver or another tools after that create the database `account-recovery` db

Run with Dockerfile

```shell script
docker build . -t yourdockerhub-account/go-account-recovery:latest
```

```shell script
 docker run -p 9999:9999 -d yourdockerhub-account/go-account-recovery:latest

# running on default port 9999
```

## Structures

```
├── cmd
│   ├── main.go
├── internal
│   ├── configs
    │       ├── config.go
    │       ├── config.yaml
    │       ├── types.go
    │── constants
    │       ├── constants.go
    |── handler
    │       ├── memberships
    │           ├── handler_mock_test.go
    │           ├── handler.go
    │           ├── login.go
    │           ├── recovery.go
    │           ├── reset.go
    │           ├── signup.go
    │           ├── verify.go
    │── helpers
    │    │── logger.go
    │    │── response.go
    ├── models
    │    │── memberships
    │       │── user.go
    ├── repository
    │    │── memberships
    │       │── repository.go
    │       │── user_test.go
    │       │── user.go
    ├── service
    │    │── memberships
    │       │── login_test.go
    │       │── login.go
    │       │── recovery_test.go
    │       │── recovery.go
    │       │── reset_test.go
    │       │── reset.go
    │       │── service_mock_test.go
    │       │── service.go
    │       │── signup_test.go
    │       │── signup.go
    │       │── verify_test.go
    │       │── verify.go
├── pkg
│   │── internalsql
│       │── sql.go
│   │── jwt
│       │── jwt.go
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── Makefile
├── README.md
```

## Features

### API Endpoint

All RESTful endpoint has `prefix` and `versioning` support. Prefix format is: /`api`/`v1`/`memberships`/routes.

API endpoint design :

- `signup`
- `login`
- `recovery`
- `verify`
- `reset-password`

#### Signup Response

```shell script
curl --location 'localhost:9999/api/v1/memberships/signup' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email" : "xprasetio@gmail.com",
    "username" : "xprasetio",
    "password" : "admin789"
}'
```

will return:

```json
{
  "message": "success"
}
```

#### Login Response

```shell script
curl --location 'localhost:9999/api/v1/memberships/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email" : "xprasetio@gmail.com",
    "password" : "admin789"
}'
```

will return:

```json
{
  "message": "success",
  "data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzkzOTAzOTYsImlkIjoyMCwidXNlcm5hbWUiOiJ4cHJhc2V0aW8ifQ.ngqCIaV5pYUv9SRHWYqvBEu-os6YiLwBEy2aSnB4ODI"
}
```

#### Recovery Response

```shell script
curl --location 'localhost:9999/api/v1/memberships/recovery' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email" : "xprasetio@gmail.com"
}'
```

will return:

```json
{
  "message": "success"
}
```

![alt text](image.png)

#### Verify Response

```shell script
curl --location 'localhost:9999/api/v1/memberships/verify' \
--header 'Content-Type: application/json' \
--data '{
    "recover_code" : "357e924d8568c1ca75f24bcc95fdc7dc"
}'
```

will return:

```json
{
  "message": "success",
  "data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzkzOTA1OTgsImlkIjoyMCwidXNlcm5hbWUiOiJ4cHJhc2V0aW8ifQ.jwadDIYKpYi_LS7PoIuCmAhB-3hDnZ48EZ31X04LkG4"
}
```

#### Reset Password Response

```shell script
curl --location 'localhost:9999/api/v1/memberships/reset-password' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzkzOTA1OTgsImlkIjoyMCwidXNlcm5hbWUiOiJ4cHJhc2V0aW8ifQ.jwadDIYKpYi_LS7PoIuCmAhB-3hDnZ48EZ31X04LkG4' \
--header 'Content-Type: application/json' \
--data '{
    "password" : "AfterResetPassword78910"
}'
```

will return:

```json
{
  "message": "success"
}
```

#### Login After Reset Password Response

```shell script
curl --location 'localhost:9999/api/v1/memberships/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email" : "xprasetio@gmail.com",
    "password" : "AfterResetPassword78910"
}'
```

will return:

```json
{
  "message": "success",
  "data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzkzOTA3NTAsImlkIjoyMCwidXNlcm5hbWUiOiJ4cHJhc2V0aW8ifQ.Oy1BZQVptxqdU_Rcuotiko8eVUSVSr4CZ28ZfGajT6A"
}
```

## Logging

### Email Not Found

```shell script
curl --location 'localhost:9999/api/v1/memberships/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email" : "xprasetio1@gmail.com",
    "password" : "AfterResetPassword78910"
}'
```

logrus will return:

```json
{
  "error": "email not found",
  "level": "error",
  "msg": "failed to login: ",
  "time": "2025-02-12T20:12:18Z"
}
```

### password or email not match

```shell script
curl --location 'localhost:9999/api/v1/memberships/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email" : "xprasetio@gmail.com",
    "password" : "AfterResetPassword7891"
}'
```

logrus will return:

```json
{
  "error": "password or email not match",
  "level": "error",
  "msg": "failed to login: ",
  "time": "2025-02-12T20:12:18Z"
}
```

### email not found

```shell script
curl --location 'localhost:9999/api/v1/memberships/recovery' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email" : "xprasetio1@gmail.com"
}'
```

logrus will return:

```json
{
  "error": "email not found",
  "level": "error",
  "msg": "failed to initiate recovery:",
  "time": "2025-02-12T20:12:18Z"
}
```

### wrong code recovery

```shell script
curl --location 'localhost:9999/api/v1/memberships/verify' \
--header 'Content-Type: application/json' \
--data '{
    "recover_code" : "357e924d8568c1ca75f24bcc95fdc7dct"
}'
```

logrus will return:

```json
{
  "error": "record not found",
  "level": "error",
  "msg": "failed to verify recovery code: ",
  "time": "2025-02-12T20:12:18Z"
}
```

### invalid token

```shell script
curl --location 'localhost:9999/api/v1/memberships/reset-password' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzkzOTA1OTgsImlkIjoyMCwidXNlcm5hbWUiOiJ4cHJhc2V0aW8ifQ.jwadDIYKpYi_LS7PoIuCmAhB-3hDnZ48EZ31X04LkG41' \
--header 'Content-Type: application/json' \
--data '{
    "password" : "AfterResetPassword78910"
}'
```

logrus will return:

```json
{
  "error": "signature is invalid",
  "level": "error",
  "msg": "failed to reset password: ",
  "time": "2025-02-12T20:12:18Z"
}
```

## Database Schema

| Column name  | Data Type | Not Null | Key |
| ------------ | --------- | -------- | --- |
| id           | BigSerial | ✅       | PK  |
| email        | String    | ✅       |     |
| username     | String    | ✅       |     |
| password     | String    | ✅       |     |
| recover_code | String    | ✅       |     |
| created_by   | String    | ✅       |     |
| updated_by   | String    | ✅       |     |
| created_at   | Timestamp |          |     |
| updated_at   | Timestamp |          |     |
| deleted_at   | Timestamp |          |     |

## Credits

- [Go](https://github.com/golang/go) - The Go Programming Language
- [gin](https://github.com/gin-gonic/gin) - Gin is HTTP web framework written in Go (Golang)
- [gorm](https://github.com/go-gorm/gorm) - The fantastic ORM library for Golang
- [viper](https://github.com/spf13/viper) - Complete configuration solution for Golang
- [docker](https://www.docker.com/products/docker-hub/) - Cloud Native a software development
- [gomock](https://github.com/uber-go/mock) - Mocking Framework Golang

## Copyright

Copyright (c) 2025 Eko Prasetio.
