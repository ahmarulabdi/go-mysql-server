package main

import (
	"fmt"
	"runtime"
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
