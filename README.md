# Snippets Project

## Introduction

This is a website for add snippet. This project is sample project from Let's Go Book by Alex Edwards. This project is intended for learning purpose only.

Feature : 
- Add Snippet
- Login
- Signup
- Logout

Implemented Technique :
- HTML templating and Inheritance
- Caching template
- Dependency Injection
- Leveled Logging
- Centralized Error Handling
- Database Driven Response
- RESTful Routing
- Middleware
- Processing Form (Setup, Parse, Validate)
- Session
- HTTPS (kSelf Signed TLS certificate)
- Authentication (Login, Logout)
- CSRF protection
- Testing (Mocking Dependencies, Unit Test, End To End, Integration)

## How To Run This Project

### 1.) Setup Database

Run sql script ./db/01_init.sql to initialize database

### 2.) Install all dependencies

Run this command at root directory : `go mod download`

### 3.) Generate a Self-Signed TLS Certificate

create dedicated folder to store certificate and key : 

`mkdir tls`

`cd tls`

Run `generate_cert.go` tool, you'll need to know the place on your computer where the source code for the Go std library is installed.
If you are using Linux, macOS, or FreeBSD then generate_cert.go file should be located under `/usr/local/go/src/crypto/tls`

Once you know where it is located, you can then run the `generate_cert.go` tool like so :

`go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost`

It then stores the private key in `key.pem` file and generates a self-signed TLS certificate for the host `localhost` containing the public key which it stores in a `cert.pem` file

### 4.) Run

Before running the project you should decide which port and what dsn (data source name) mysql first. Default value :

`port` : 4000

`mysql dsn` : root:secret@tcp(localhost:3306)/snippetbox?parseTime=true

run app with default port and mysql dsn : `go run ./cmd/web/!(*_test).go`

for more information you can run : `go run ./cmd/web/!(*_test).go -h`
