package model

import (
	"encoding/json"
	"hangoutin/authentication/internal/constant"

	"github.com/golang-jwt/jwt/v4"
)

type JwtConfig struct {
	SigningKey string
	Iss        string
	Exp        string
}

type Claims struct {
	UserId      uint64             `json:"user_id"`
	Name        string             `json:"name"`
	Username    string             `json:"username"`
	Email       string             `json:"email"`
	PhoneNumber string             `json:"phone_number"`
	Roles       []string           `json:"roles"`
	TokenType   constant.TokenType `json:"token_type"`
	jwt.RegisteredClaims
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// MarshalBinary fulfills encoding.BinaryMarshaler implementation.
func (t Token) MarshalBinary() ([]byte, error) {
	return json.Marshal(t)
}
