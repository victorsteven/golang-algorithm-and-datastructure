package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {

	fmt.Println("this is the password: ", Encrypt("123456"))
}


// Encrypt securely encrypts a Password object
func Encrypt(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

// Check compared password and returns if they match
func Check(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(password))
	return err == nil
}
