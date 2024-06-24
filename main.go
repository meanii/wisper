package main

import (
	"fmt"
	"log"

	"github.com/chzyer/readline"

	"wisper/configs"
	"wisper/plugins"
)

func main() {
	configs.Init()

	ShellName := fmt.Sprintf("%s > ", configs.Config.Name)
	commands := plugins.Plugins{}

	reader, err := readline.New(ShellName)
	if err != nil {
		log.Fatalf("failed to launch %s instance! ERROR: %s", ShellName, err)
	}

	defer func(reader *readline.Instance) {
		err := reader.Close()
		if err != nil {
			log.Fatalf("failed to close the %s instance! ERROR: %s", ShellName, err)
		}
	}(reader)

	for {
		fmt.Print(ShellName)
		userInput, err := reader.Readline()
		if err != nil {
			fmt.Printf("want to exit? type 'exit' to exit\n")
			continue
		}

		// handling all commands, including exit
		// invoke the method if the command is found
		collections := commands.Plugins()
		for _, plugin := range collections {
			if plugin.Name == userInput {
				// type assertion to check if the method is a function
				if fn, ok := plugin.Method.(func()); ok {
					fn()
				}
			}
		}

	}
}
