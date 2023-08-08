package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"syscall"

	"golang.org/x/term"
)

const (
	MYSQL_WINDOWS_AMD64 = "mysql-windows-amd64"
	MYSQL_LINUX_AMD64   = "mysql-linux-amd64"
)

func main() {
	mysqlOsArch := getMysqlOsArch()
	fmt.Println(mysqlOsArch)

	if isDataDirExists(mysqlOsArch) {
		runService(mysqlOsArch)
	} else {
		runSetup(mysqlOsArch)
	}
}

func runService(mysqlOsArch string) {
	switch mysqlOsArch {
	case MYSQL_WINDOWS_AMD64:
		cmd := exec.Command("cmd", "/c", "start", "cmd", "/K", MYSQL_WINDOWS_AMD64+"\\bin\\mysqld.exe --console")
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
	}
}

func runSetup(mysqlOsArch string) {
	switch mysqlOsArch {
	case MYSQL_WINDOWS_AMD64:
		fmt.Print("Database credential doesn't exists, please create your new credential\n\n")
		fmt.Print("\nPassword:")
		password := getPasswordInput()

		fmt.Print("\nPassword confirmation:")
		passwordConfirmation := getPasswordInput()

		if password != passwordConfirmation {
			fmt.Println("\n\nPassword confirmation is not same!")
			return
		}

		fmt.Println("\nPlease wait...")
		// init
		cmd := exec.Command(MYSQL_WINDOWS_AMD64+"/bin/mysqld", "--initialize-insecure")
		err := cmd.Run()
		if err != nil {
			fmt.Println(err.Error())
		}

		// run mysql service
		runService(MYSQL_WINDOWS_AMD64)
		fmt.Println("Run database server...")

		// Command and its arguments as separate elements
		args := []string{"-u", "root"}
		// Create the exec.Cmd object
		cmd3 := exec.Command(MYSQL_WINDOWS_AMD64+"/bin/mysql.exe", args...)
		fmt.Println("Creating new credential on it...")

		// Get a writable pipe to the command's standard input
		stdin, err := cmd3.StdinPipe()
		if err != nil {
			fmt.Println("Error creating stdin pipe:", err.Error())
			return
		}

		// Start the command
		err = cmd3.Start()
		if err != nil {
			fmt.Println("Error starting command:", err.Error())
			return
		}

		// set mysql root password
		_, err = stdin.Write([]byte("ALTER USER 'root'@'localhost' IDENTIFIED BY '" + string(password) + "';\n"))
		if err != nil {
			fmt.Println("Error writing to command:", err.Error())
			return
		}

		// Close the stdin pipe to indicate that we are done writing
		stdin.Close()

		// Wait for the command to finish
		err = cmd3.Wait()
		if err != nil {
			fmt.Println("Error executing command:", err.Error())
			return
		}

		fmt.Println("DONE!")
	}
}

func getMysqlOsArch() string {
	goos := runtime.GOOS
	goarch := runtime.GOARCH

	return "mysql-" + goos + "-" + goarch
}

func isDataDirExists(osArch string) bool {
	switch osArch {
	case MYSQL_WINDOWS_AMD64:
		_, err := os.Stat(MYSQL_WINDOWS_AMD64 + "\\data")
		if err != nil {
			return false
		}
		return true
	}

	return false
}

func getPasswordInput() string {
	result, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Println("Error reading password:", err.Error())
	}

	return string(result)
}
