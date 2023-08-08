package main

import (
	"fmt"
	"runtime"
)

const (
	MYSQL_WINDOWS_AMD64 = "mysql-windows-amd64"
	MYSQL_LINUX_AMD64   = "mysql-linux-amd64"
)

func main() {
	mysqlOsArch := getMysqlOsArch()
	fmt.Println(mysqlOsArch)
}

func getMysqlOsArch() string {
	goos := runtime.GOOS
	goarch := runtime.GOARCH

	return "mysql-" + goos + "-" + goarch
}
