package helper

import (
	"kembanglah/model/domain"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtCustomClaims struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func GenerateToken(user domain.User) (string, error) {
	claims := JwtCustomClaims{
		Id:       user.ID,
		Username: user.Username,
		Role:     user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	return t, err
}
