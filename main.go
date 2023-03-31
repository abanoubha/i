package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
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

	t := os.Args[1]

	fmt.Println("searching for", t)

	// check if the app already installed
	path, err := exec.LookPath(t)
	// if errors.Is(err, exec.ErrDot) {
	// 	err = nil
	// }
	// if err != nil {
	// 	log.Fatal(err)
	// }
	if errors.Is(err, exec.ErrNotFound) {
		searchForApp(t)
		os.Exit(0)
	}
	fmt.Println("the app you are looking for is already installed in this path", path)
	os.Exit(0)
}

func searchForApp(t string) {
	// brew
	_, err := exec.LookPath("brew")
	if errors.Is(err, exec.ErrNotFound) {
		fmt.Println("HomeBrew (brew) is not installed")
	} else {
		cmd := exec.Command("brew", "search", t)
		// cmd.Stdin = strings.NewReader("some input")
		var out strings.Builder
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		results := strings.Split(out.String(), "\n")

		fmt.Printf("%q\n", cmd)
		for _, r := range results {
			if r == "" {
				continue
			}
			fmt.Printf("%q\n", r)
		}
	}

	// apt
}
