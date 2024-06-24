package utils

import (
	"os"

	"golang.org/x/term"
)

func ReadPassword() (string, error) {
	password, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}
	return string(password), nil
}
