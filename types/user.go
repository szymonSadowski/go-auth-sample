package types

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Name         string `json:"name"`
	ID           string `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
}

type CreateUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUser(createUser CreateUser) (*User, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(createUser.Password), 14)
	if err != nil {
		return nil, err
	}
	return &User{
		Email:        createUser.Email,
		PasswordHash: string(hashPassword),
	}, nil
}

func ValidatePassword(hashPassword string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password)) == nil
}

func CreateToken(user User) string {
	now := time.Now()
	validUntil := now.Add(time.Hour * 4).Unix()

	claims := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   validUntil,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")

	tokenStr, err := token.SignedString([]byte(secret))

	if err != nil {
		fmt.Printf("Failed to sign token: %v", err)
	}
	return tokenStr
}
