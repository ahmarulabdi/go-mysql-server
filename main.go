package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
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
