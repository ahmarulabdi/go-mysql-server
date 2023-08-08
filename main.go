package main

import (
	"fmt"
	"os"
	"runtime"
)

const (
	MYSQL_WINDOWS_AMD64 = "mysql-windows-amd64"
	MYSQL_LINUX_AMD64   = "mysql-linux-amd64"
)

func main() {
	mysqlOsArch := getMysqlOsArch()
	fmt.Println(mysqlOsArch)

	fmt.Println(isDataDirExists(mysqlOsArch))
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
