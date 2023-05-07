package cmd

import "fmt"

func (cmd *Commands) Auth() {
	var username string
	var password string

	fmt.Println("Username: ")
	fmt.Scanln(&username)

	fmt.Println("Password: ")
	fmt.Scanln(&password)

	fmt.Printf("Username: %s\n", username)
}
