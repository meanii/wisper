package plugins

import (
	"fmt"
)

// Help displays a list of available commands and their descriptions
func generateHelpMessage(collections []CollectionInfo) string {
	var helpMessage string
	for _, collection := range collections {
		helpMessage += fmt.Sprintf("  %s - %s\n", collection.Name, collection.Description)
	}
	return helpMessage
}

func HandlerPositionIndex(collections []CollectionInfo) []CollectionInfo {
	var sortedCollections []CollectionInfo
	for _, collection := range collections {
		if collection.PositionIndex != 0 {
			sortedCollections = append(sortedCollections, collection)
		} else {
			sortedCollections = append([]CollectionInfo{collection}, sortedCollections...)
		}
	}
	return sortedCollections
}

// Help displays a list of available commands and their descriptions
func (p *Plugins) Help() {
	collectionInfos := p.PluginsInfo()
	collectionInfos = HandlerPositionIndex(collectionInfos)

	const HelpMessage = `Wisper is an end-to-end encrypted CLI-based chatting application.

Usage:
%s

This program is written by <https://github.com/meanii> - Anil Chauhan
For more information, please refer to the Wisper documentation at https://wisper.meanii.dev/docs
`

	fmt.Printf(HelpMessage, generateHelpMessage(collectionInfos))
}

func (p *Plugins) HelpInfo() CollectionInfo {
	return CollectionInfo{
		Name:          "help",
		Description:   "Display a list of available commands and their descriptions",
		PositionIndex: 100,
	}
}
