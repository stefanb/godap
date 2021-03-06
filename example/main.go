package main

import (
	"godap"
	"log"
)

func main() {

	//create a new client
	client := godap.New(godap.Options{
		BaseDN:   "CN=Users,DC=Company",
		Filter:   "sAMAccountName",
		Password: "username",
		Username: "password",
		URL:      "ldap.directory.com:389",
	})

	//username and password to authenticate
	username := "jdoe"
	password := "pass1234"

	//attempt the authentication
	err := client.Authenticate(username, password)

	//see the results
	if err != nil {
		log.Print(err)
	} else {
		log.Print("user successfully authenticated!")
	}
}
