package auth

import (
	"errors"
	"go-lab/09-library-api/internal/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	UserID int64
	Email string
	jwt.RegisteredClaims
}

const secret_string = "secret_string"

func GenerateToken(userId int64, email string)(string, error){
	claims := Claims{
		UserID:userId ,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24*time.Hour)),
		},
	} 
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret_string))
}

func ParseToken(tokenString string)(*Claims, error){
	token,err := jwt.ParseWithClaims(
		tokenString, 
		&Claims{}, 
		func(token *jwt.Token) (any, error) {
			return []byte(secret_string), nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}


func Compare(hash, password string) error {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password),
	)

	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return domain.ErrInvalidPassword
		}
		return err
	}

	return nil
}