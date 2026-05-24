# go-login-portal

A simple Go web application demonstrating user authentication with session cookies and CSRF tokens.

## Features

- User registration with password hashing
- User login with session token cookie generation
- CSRF token issuance for protected requests
- In-memory user store for development and testing

## Run locally

```bash
go run .
```

The server listens on `http://localhost:8080`, imma add more ports later on 

## Endpoints

- `POST /register` — register a new user
- `POST /login` — authenticate and receive session cookies
- `POST /logout` — logout the current session
- `GET /protected` — example protected route
make sure you edit the csrf tokens when using it with postman (i mean you have to)

## Notes

- This project uses an in-memory store; restart the server clears users.
- The implementation is intended as a learning/demo project, not production-ready.
