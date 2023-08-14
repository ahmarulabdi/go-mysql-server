package helpers

import (
	"ahmarulabdi/gomysqlserver/m/config"
	"os"
	"runtime"
)

func GetMysqlOsArch() string {
	goos := runtime.GOOS
	goarch := runtime.GOARCH

	return "mysql-" + goos + "-" + goarch
}

func IsDataDirExists(osArch string) bool {
	switch osArch {
	case config.MYSQL_WINDOWS_AMD64:
		_, err := os.Stat(config.MYSQL_WINDOWS_AMD64 + "\\data")
		if err != nil {
			return false
		}
		return true
	case config.MYSQL_LINUX_AMD64:
		_, err := os.Stat(config.MYSQL_LINUX_AMD64 + "/data")
		if err != nil {
			return false
		}
		return true
	}

	return false
}
