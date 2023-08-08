package main

import (
	"ahmarulabdi/gomysqlserver/m/config"
	"ahmarulabdi/gomysqlserver/m/helpers"
	"fmt"
	"os/exec"
)

func main() {
	mysqlOsArch := helpers.GetMysqlOsArch()
	fmt.Println(mysqlOsArch)

	if helpers.IsDataDirExists(mysqlOsArch) {
		runService(mysqlOsArch)
	} else {
		runSetup(mysqlOsArch)
	}
}

func runService(mysqlOsArch string) {
	switch mysqlOsArch {
	case config.MYSQL_WINDOWS_AMD64:
		cmd := exec.Command("cmd", "/c", "start", "cmd", "/K", config.MYSQL_WINDOWS_AMD64+"\\bin\\mysqld.exe --console")
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
	}
}

func runSetup(mysqlOsArch string) {
	switch mysqlOsArch {
	case config.MYSQL_WINDOWS_AMD64:
		password := helpers.SetupPassword()
		if password == "" {
			return
		}

		fmt.Println("\nPlease wait...")
		// init
		cmd := exec.Command(config.MYSQL_WINDOWS_AMD64+"/bin/mysqld", "--initialize-insecure")
		err := cmd.Run()
		if err != nil {
			fmt.Println(err.Error())
		}

		// run mysql service
		runService(config.MYSQL_WINDOWS_AMD64)
		fmt.Println("Run database server...")

		// Command and its arguments as separate elements
		args := []string{"-u", "root"}
		// Create the exec.Cmd object
		cmd3 := exec.Command(config.MYSQL_WINDOWS_AMD64+"/bin/mysql.exe", args...)
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
