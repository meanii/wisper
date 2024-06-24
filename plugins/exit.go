package plugins

import (
	"fmt"
	"os"
)

func (p *Plugins) Exit() {
	fmt.Println("\n\nadios!")
	os.Exit(0)
}

func (p *Plugins) ExitInfo() CollectionInfo {
	return CollectionInfo{
		Name:          "exit",
		Description:   "Exit the application",
		PositionIndex: 99,
	}
}
