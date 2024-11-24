package main

import (
	"GO/foo"
	"fmt"
)

func appName() string {
	const AppName = "GOAPP"
	var Version string = "1.0"
	return AppName + " " + Version
}
func main() {
	fmt.Println(foo.MAX)
	fmt.Println(appName())

	// fmt.Println(AppName, Version)
}
