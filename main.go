package main

import (
	"ahmarulabdi/gomysqlserver/m/config"
	"ahmarulabdi/gomysqlserver/m/helpers"
	"fmt"
	"os/exec"
	"time"
)

func main() {
	mysqlOsArch := helpers.GetMysqlOsArch()
	fmt.Println(mysqlOsArch)

	fmt.Println("Readme first!")
	fmt.Println(`
	Mysql folder each OS ARCH:
	- mysql-windows-amd64: mysql folder for mysql OS windows architecture amd64
	- mysql-linux-amd64: mysql folder for mysql OS linux architecture amd64
	`)
	fmt.Println("Press <Enter> to continue!")
	fmt.Scanln()

	if helpers.IsDataDirExists(mysqlOsArch) {
		fmt.Println("run service...")
		runService(mysqlOsArch)
	} else {
		fmt.Println("run setup...")
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
	case config.MYSQL_LINUX_AMD64:
		cmd := exec.Command("x-terminal-emulator", "-e", "bash -c './"+config.MYSQL_LINUX_AMD64+"/bin/mysqld --console "+config.LINUX_BASE_DIR_FLAG+" "+config.LINUX_DATA_DIR_FLAG+"'")
		err := cmd.Run()
		fmt.Println("run command", cmd.String())
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

	case config.MYSQL_LINUX_AMD64:
		password := helpers.SetupPassword()
		if password == "" {
			return
		}

		fmt.Println("\nPlease wait...")
		// init
		cmd := exec.Command("./"+config.MYSQL_LINUX_AMD64+"/bin/mysqld", "--initialize-insecure", config.LINUX_BASE_DIR_FLAG, config.LINUX_DATA_DIR_FLAG)
		fmt.Println("init mysql data with command:", cmd.String())
		err := cmd.Run()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// run mysql service
		runService(config.MYSQL_LINUX_AMD64)
		fmt.Println("Run database server...")

		time.Sleep(time.Second * 5)
		// Command and its arguments as separate elements
		args := []string{"-u", "root"}
		// Create the exec.Cmd object
		cmd3 := exec.Command(config.MYSQL_LINUX_AMD64+"/bin/mysql", args...)
		fmt.Println("Creating new credential on it...")

		// Get a writable pipe to the command's standard input
		stdin, err := cmd3.StdinPipe()
		if err != nil {
			fmt.Println("Error creating stdin pipe:", err.Error())
			return
		}

		// Start the command
		err = cmd3.Start()
		fmt.Println("login to mysql:", cmd3.String())
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
