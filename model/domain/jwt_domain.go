package domain

import "github.com/golang-jwt/jwt"

type JwtCustomClaims struct {
	Id       uint   `json:"id"`
	FullName string `json:"full_name"`
	Role     string `json:"role"`
	jwt.StandardClaims
}
