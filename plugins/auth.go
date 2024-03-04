package plugins

import (
	"fmt"
	"log"

	"wisper/utils"
)

func (p *Plugins) Auth() {
	var username string
	var password string

	fmt.Println("Username: ")
	_, err := fmt.Scanln(&username)
	if err != nil {
		log.Fatalf("username couldn't read! ERROR: %s", err)
	}

	fmt.Println("Password: ")
	password, err = utils.ReadPassword()
	if err != nil {
		log.Fatalf("password couldn't read! ERROR: %s", err)
	}

	fmt.Printf("Username: %s\nPassword: %s\n", username, password)
}

func (p *Plugins) AuthInfo() CollectionInfo {
	return CollectionInfo{
		Name:        "auth",
		Description: "Log in or sign up for a Wisper account",
	}
}
