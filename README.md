# Dokumentasi BE
BackEnd dapat diakses di http://159.223.82.24:3000
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
