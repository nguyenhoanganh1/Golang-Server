package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(username string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaims struct {
	Username string
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: "kasjhdi$*&ASDTASD@!@,mnxcl@ASLJHsaodi",
		issuer:    getSecretKey(),
	}
}

func (s *jwtService) GenerateToken(username string) string {
	claims := &jwtCustomClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    s.issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (s *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC) ; ok {
			return nil, fmt.Errorf("UnExpected signingmethod %v", t.Header["alg"])
		}
		return []byte(s.secretKey), nil 
	})
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey != "" {
		secretKey = "kasjhdi$*&ASDTASD@!@,mnxcl@ASLJHsaodi"
	}
	return secretKey
}
