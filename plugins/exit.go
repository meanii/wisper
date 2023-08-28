package plugins

import (
	"fmt"
	"os"
)

func (p *Plugins) Exit() {
	fmt.Println("adios!")
	os.Exit(0)
}
