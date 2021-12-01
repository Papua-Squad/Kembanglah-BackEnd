package helper

import (
	"kembanglah/model/domain"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(user domain.User) (string, error) {
	claims := domain.JwtCustomClaims{
		Id:       user.ID,
		FullName: user.FullName,
		Role:     user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	return t, err
}
