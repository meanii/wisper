package main

import (
	"fmt"
	"github.com/chzyer/readline"
	"os"
	"wisper/cmd"
)

func main() {
	user := "meanii > "
	commands := cmd.Commands{}

	reader, err := readline.New(user)
	if err != nil {
		fmt.Println("failed to launch wisper instance! ERROR: ", err)
		os.Exit(1)
	}

	defer func(reader *readline.Instance) {
		err := reader.Close()
		if err != nil {

		}
	}(reader)

	for {
		fmt.Print(user)
		userInput, err := reader.Readline()

		if err != nil {
			fmt.Println("failed to read the line! ERROR: ", err)
			continue
		}

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
