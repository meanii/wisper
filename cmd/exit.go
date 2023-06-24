package cmd

import (
	"fmt"
	"os"
)

func (cmd Commands) Exit() {
	fmt.Println("adios!")
	os.Exit(0)
}
