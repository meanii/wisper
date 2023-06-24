package main

import (
	"fmt"
	"log"
	"wisper/cmd"

	"github.com/chzyer/readline"
)

func main() {
	SHELL_NAME := "wisper > "
	commands := cmd.Commands{}

	reader, err := readline.New(SHELL_NAME)
	if err != nil {
		log.Fatalln("failed to launch wisper instance! ERROR: ", err)
	}

	defer func(reader *readline.Instance) {
		err := reader.Close()
		if err != nil {
			log.Fatalln("failed to close the wisper instance! ERROR: ", err)
		}
	}(reader)

	for {
		fmt.Print(SHELL_NAME)
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
		}
	}
}
