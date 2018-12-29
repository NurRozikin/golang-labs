package auth

import (
	"fmt"
	"time"

	jwt "gopkg.in/dgrijalva/jwt-go.v3"
)

type UserCredential struct {
	Email []byte `json:"email"`
	Name  []byte `json:"name"`
	jwt.StandardClaims
}

var JWTSignKey = []byte("k13N")

func CreateToken(email []byte, name []byte) (string, string) {
	expireAt := time.Now().Add(time.Hour * 1)

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS512, &UserCredential{
		Email: email,
		Name:  name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireAt.UTC().Unix(),
			IssuedAt:  time.Now().UTC().Unix(),
		},
	})

	tokenString, err := newToken.SignedString(JWTSignKey)
	if err != nil {
		fmt.Println(err.Error)
	}
	return tokenString, expireAt.Format(time.RFC3339)
}

func JWTValidateWithClaims(requestToken string) (*jwt.Token, *UserCredential, error) {
	user := &UserCredential{}
	token, err := jwt.ParseWithClaims(requestToken, user, func(token *jwt.Token) (interface{}, error) {
		return JWTSignKey, nil
	})
	claims := token.Claims.(*UserCredential)
	return token, claims, err
}
