package plugins

import "fmt"

func (p *Plugins) Help() {
	const HelpMessage = `Wisper is an end-to-end encrypted CLI-based chatting application.

Usage:
  help - display this message
  about - display information about the application
  auth - log in or sign up

Plugins:
  help   Display a list of available commands and their descriptions
  about  Display general information about the Wisper application
  auth   Log in or sign up for a Wisper account

Usage examples:
  $ wisper help
  $ wisper about
  $ wisper auth

For more information, please refer to the Wisper documentation at https://wisper.meanii.online/docs`

	fmt.Println(HelpMessage)
}
