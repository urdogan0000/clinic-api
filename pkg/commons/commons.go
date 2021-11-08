package commons

import (
	"github.com/dgrijalva/jwt-go"
	ent "main/internals/interfaces"
	"os"
	"time"
)

func GenerateToken(claims ent.TokenClaim) (string, error) {
	signingKey := []byte(os.Getenv("app_secret"))
	claims.Iat = time.Now().Unix()
	claims.Exp = time.Now().Local().Add(time.Second * time.Duration(3600)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(signingKey)
	return tokenString, err
}
