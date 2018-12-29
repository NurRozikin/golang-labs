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
	fmt.Printf("** Check Token : %s and Get Data **\n", tokenString)
	userCredential := &auth.UserCredential{}
	token, userCredential, _ := auth.JWTValidateWithClaims(tokenString)
	fmt.Printf("** token valid ? %t\n", token.Valid)
	if token.Valid {
		fmt.Println("Email : " + string(userCredential.Email))
		fmt.Println("Name : " + string(userCredential.Name))
	}

}
