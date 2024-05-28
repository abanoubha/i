package main

import (
	"fmt"
	"os"
)

const version = "0.1.0"

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("i the installer v%v\nUsage:\n    i <package-name>", version)
		return
	}

	if len(os.Args) == 2 {
		switch os.Args[1] {
		case "version", "--version", "-v":
			fmt.Printf("i the installer v%v", version)
			return
		case "help", "--help", "-h":
			fmt.Printf("i the abstraction over all package managers.\nUsage:\n  i install vim\n  i info vim\n  i search vim\n  i uninstall vim")
			return
		case "update", "upgrade", "up":
			fmt.Println("upgrading all installed apps...")
			return
		case "install", "add":
			fmt.Println("No package/app specified.\nUsage:\n  i install vim\n  or\n  i add vim")
			return
		case "search", "find":
			fmt.Println("No package/app specified to search for.\nUsage:\n  i search vim\n  or\n  i find vim")
			return
		default:
			fmt.Printf("'%v' sub-command is not supported in 'i'.\ntry one of these commands:\n  i install vim\n  i info vim\n  i search vim\n  i uninstall vim", os.Args[1])
			return
		}
	}

	// // check if the app already installed
	// path, err := exec.LookPath(t)
	// // if errors.Is(err, exec.ErrDot) {
	// // 	err = nil
	// // }
	// // if err != nil {
	// // 	log.Fatal(err)
	// // }
	// if errors.Is(err, exec.ErrNotFound) {
	// 	searchForApp(t)
	// 	os.Exit(0)
	// }
	// fmt.Println("the app you are looking for is already installed in this path", path)
	// os.Exit(0)
}

// func searchForApp(t string) {
// 	// brew
// 	_, err := exec.LookPath("brew")
// 	if errors.Is(err, exec.ErrNotFound) {
// 		fmt.Println("HomeBrew (brew) is not installed")
// 	} else {
// 		cmd := exec.Command("brew", "search", t)
// 		// cmd.Stdin = strings.NewReader("some input")
// 		var out strings.Builder
// 		cmd.Stdout = &out
// 		err := cmd.Run()
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		results := strings.Split(out.String(), "\n")

// 		fmt.Printf("%q\n", cmd)
// 		for _, r := range results {
// 			if r == "" {
// 				continue
// 			}
// 			fmt.Printf("%q\n", r)
// 		}
// 	}

// 	// apt
// }
