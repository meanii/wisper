package utils

import (
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

func ReadPassword() (string, error) {
	password, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}
	return string(password), nil
}
