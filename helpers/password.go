package helpers

import (
	"fmt"
	"syscall"

	"golang.org/x/term"
)

func SetupPassword() string {
	fmt.Print("Database credential doesn't exists, please create your new credential\n\n")
	fmt.Print("\nPassword:")
	password := GetPasswordInput()

	fmt.Print("\nPassword confirmation:")
	passwordConfirmation := GetPasswordInput()

	if password != passwordConfirmation {
		fmt.Println("\n\nPassword confirmation is not same!")
		return ""
	}

	return password
}

func GetPasswordInput() string {
	result, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Println("Error reading password:", err.Error())
	}

	return string(result)
}
