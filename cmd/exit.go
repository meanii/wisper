package cmd

import (
	"fmt"
	"os"
)

func (cmd Commands) Exit() {
	fmt.Println("leaving...")
	os.Exit(0)
}
