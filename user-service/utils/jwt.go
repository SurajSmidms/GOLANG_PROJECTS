package utils

import (
	"time"
	"user-service/models"

	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

var JwtKey = []byte("supersecretkey")

type Claims struct {
	UserID uint
	Email  string
	jwt.RegisteredClaims
}

func GenerateJWT(userID uint, email string, db *gorm.DB) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	jti := email + "-" + time.Now().Format("20060102150405")

	claims := &Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        jti,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return "", err
	}

	newToken := models.Token{
		UserID:    userID,
		JTI:       jti,
		Token:     tokenString,
		ExpiresAt: expirationTime,
	}
	db.Create(&newToken)

	return tokenString, nil
}
