package main

import (
	"fmt"
	"os"
)

// i vim
// go run . vim
func main() {
	if len(os.Args) < 2 {
		fmt.Println("i the installer\nUsage\ni <package-name>")
		os.Exit(1)
	}

	if len(os.Args) > 2 {
		fmt.Println("i the installer can not insall more than one package/app at once")
		os.Exit(1)
	}

	fmt.Println("searching for", os.Args[1])
	os.Exit(0)
}
