package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type jwtClaim struct {
	jwt.StandardClaims
	UserID int64 `json:"user_id"`
	Exp    int64 `json:"exp"`
}

func CreateToken(user_id int64, Signing_key string) (string, error) {
	claims := jwtClaim{
		UserID: user_id,
		Exp:    time.Now().Add(time.Minute * 30).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(Signing_key))

	if err != nil {
		return "", err
	}

	return string(tokenString), nil
}

func VerifyToken(tokenString string, Signing_key string) (int64, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwtClaim{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Token invalide")
		}
		return []byte(Signing_key), nil
	})

	if err != nil {
		return 0, fmt.Errorf("Token invalide")
	}

	claims, ok := token.Claims.(*jwtClaim)

	if !ok {
		return 0, fmt.Errorf("Token invalide")
	}

	if claims.Exp < time.Now().Unix() {
		return 0, fmt.Errorf("Token expiré")
	}

	return claims.UserID, nil
}
