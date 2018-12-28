package main

import (
	"fmt"
	"log"

	"./auth"
)

func main() {

	fmt.Println("** Create Token **")
	tokenString, expire := auth.CreateToken([]byte("nur.rozikien@gmail.com"), []byte("Nur Rozikin"))
	log.Println("token : " + tokenString)
	log.Println("expiry : " + expire)

	fmt.Println("")
	fmt.Println("** Check Token and Get Data **")
	userCredential := &auth.UserCredential{}
	token, userCredential, _ := auth.JWTValidateWithClaims(tokenString)
	fmt.Printf("** token valid ? %t\n", token.Valid)
	if token.Valid {
		fmt.Println(string(userCredential.Email))
		fmt.Println(string(userCredential.Name))
	}

}
