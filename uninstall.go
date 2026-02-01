package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func selfUninstall() {
	installName := os.Getenv("INSTALL_NAME")
	if installName == "" {
		installName = "i"
	}
	installDir := os.Getenv("INSTALL_DIR")

	var target string

	if installDir != "" {
		target = filepath.Join(installDir, installName)
	} else {
		foundPath, err := exec.LookPath(installName)
		if err == nil {
			target = foundPath
		} else {
			fallbacks := []string{"/usr/local/bin", "/usr/bin"}
			for _, dir := range fallbacks {
				path := filepath.Join(dir, installName)
				if info, err := os.Stat(path); err == nil && !info.IsDir() {
					target = path
					break
				}
			}
		}
	}

	if target == "" {
		fmt.Fprintf(os.Stderr, "%s: not found in PATH and no INSTALL_DIR given.\n", installName)
		os.Exit(1)
	}

	target = filepath.Clean(target)
	if _, err := os.Lstat(target); os.IsNotExist(err) {
		fmt.Printf("Nothing to remove: %s does not exist.\n", target)
		os.Exit(0)
	}

	fmt.Printf("This will remove:\n  %s\n", target)
	fmt.Print("Proceed? [y/N] ")

	reader := bufio.NewReader(os.Stdin)
	ans, _ := reader.ReadString('\n')
	ans = strings.ToLower(strings.TrimSpace(ans))

	if ans != "y" && ans != "yes" {
		fmt.Println("Aborted.")
		os.Exit(1)
	}

	err := os.Remove(target)
	if err != nil {
		if os.IsPermission(err) {
			if err := runAsSuperUser("rm", "-f", target); err != nil {
				fmt.Fprintf(os.Stderr, "Failed to remove %s (even with sudo/doas).\n", target)
				os.Exit(1)
			}
			fmt.Fprintf(os.Stderr, "Cannot remove %s: permission denied and sudo/doas not found.\n", target)
			os.Exit(1)
		} else {
			fmt.Fprintf(os.Stderr, "Error removing file: %v\n", err)
			os.Exit(1)
		}
	}

	fmt.Printf("[info] uninstalled %s from %s\n", installName, target)
}
