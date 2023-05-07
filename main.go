package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"wisper/cmd"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	user := "meanii > "

	commands := cmd.Commands{}

	for {
		fmt.Print(user)
		userInput, _ := reader.ReadString('\n')

		// Remove newline character from input
		userInput = strings.TrimSuffix(userInput, "\n")

		switch userInput {
		case "exit":
			commands.Exit()
		case "auth":
			commands.Auth()
		case "help":
			commands.Help()
		default:
			commands.Help()
		}
	}
}
