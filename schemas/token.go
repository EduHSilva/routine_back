package schemas

import (
	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	UserID uint   `json:"userId"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

func (t *Token) Valid() error {
	return nil
}
