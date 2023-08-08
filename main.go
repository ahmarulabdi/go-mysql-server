package main

import (
	"fmt"
	"runtime"
)

const (
	WINDOWS_AMD64 = "windows-amd64"
	LINUX_AMD64   = "linux-amd64"
)

func main() {
	osArch := getOsArch()
	fmt.Println(osArch)
}

func getOsArch() string {
	goos := runtime.GOOS
	goarch := runtime.GOARCH

	return goos + "-" + goarch
}
