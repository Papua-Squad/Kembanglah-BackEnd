# Dokumentasi BE
BackEnd dapat diakses di http://159.223.82.24:3000

# Requirements
1. [Golang](https://duckduckgo.com).

# How to run?

## No Docker
1. Git Clone this repository
2. run main.go

## With Docker
1. docker-compose up
2. for more [details](https://docs.docker.com/compose/gettingstarted/)

# Endpoint

### Login 
Header
```http
POST /login HTTP/1.1
Content-Type: application/json
```
Request Body:
```json
{
    "username": "mitra",
    "password": "mitra"
}
```

Response Body:
```json
{
    "auth_token": "jwt token"
}
```

### Register 

Header
```http
POST /login HTTP/1.1
Content-Type: application/json
```
Request Body:
```json
{
    "full_name": "nama anda",
    "email": "email@email.com",
    "username": "mitra",
    "password": "mitra",
    "role": "mitra"
}
```



Response Body:
```json
{
    "id": "1",
    "full_name": "nama anda",
    "email": "email@email.com",
    "username": "mitra"
}
```
