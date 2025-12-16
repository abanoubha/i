package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)

var operatingSystem string

type packageManager struct {
	Name string
	Path string
}

var pm packageManager

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("i the installer v%v\nUsage:\n    i install <package-name>\n", version)
		return
	}

	// detect the OS and the PM
	os_pm()

	if pm.Name == "" {
		fmt.Println("No supported package manager found.")
		os.Exit(1)
	}

	// fmt.Println("DEBUG: Using PM:", pm.Name)

	action := os.Args[1]
	var pkgName string
	if len(os.Args) > 2 {
		pkgName = os.Args[2]
	}

	cmds, ok := pm_commands[pm.Name]
	if !ok {
		fmt.Printf("Configurations for package manager '%s' not found.\n", pm.Name)
		os.Exit(1)
	}

	switch action {
	case "version", "--version", "-v":
		fmt.Printf("i the installer v%v\n", version)
	case "help", "--help", "-h":
		fmt.Printf("i the abstraction over all package managers.\nUsage:\n  i install vim\n  i info vim\n  i search vim\n  i uninstall vim\n")
	case "info", "show":
		if pkgName == "" {
			fmt.Println("No package specified.")
			return
		}
		executeCommand(cmds.Info, pkgName)
	case "update", "upgrade", "up":
		if pkgName == "" {
			// Upgrade all
			executeCommand(cmds.UpgradeAll, "")
		} else {
			executeCommand(cmds.Upgrade, pkgName)
		}
	case "install", "add":
		if pkgName == "" {
			fmt.Println("No package specified.")
			return
		}
		executeCommand(cmds.Install, pkgName)
	case "uninstall", "remove", "rm":
		if pkgName == "" {
			fmt.Println("No package specified.")
			return
		}
		executeCommand(cmds.Uninstall, pkgName)
	case "reinstall":
		// Many PMs don't have a direct reinstall, so often it's install --reinstall or just install.
		// For now we'll just try install (many PMs handle it) or we could define Reinstall in commands.
		// Since we didn't add Reinstall to commands struct yet, let's just use Install or warn.
		// For safety, let's warn.
		fmt.Println("Reinstall not explicitly supported yet. Try install.")
	case "search", "find":
		if pkgName == "" {
			fmt.Println("No term specified to search.")
			return
		}
		executeCommand(cmds.Search, pkgName)
	case "list", "installed":
		executeCommand(cmds.ListInstalled, "")
	default:
		fmt.Printf("'%v' sub-command is not supported.\n", action)
	}
}

func os_pm() {
	operatingSystem = runtime.GOOS
	switch operatingSystem {
	case "windows":
		// TODO: scoop, choco or winget ?
		// For now, simple check like original, but we didn't populate windows commands yet except comments.
		fmt.Println("Windows support is minimal.")
	case "darwin":
		isHomebrewInstalled, path := isInstalled("brew")
		if isHomebrewInstalled {
			pm = packageManager{Name: "brew", Path: path}
		} else {
			isMacportsInstalled, path := isInstalled("port")
			if isMacportsInstalled {
				pm = packageManager{Name: "port", Path: path}
			} else {
				// Fallback or just inform
			}
		}
	case "linux":
		type OsRelease struct {
			ID   string `json:"ID"`
			Name string `json:"NAME"`
		}

		// Try to read /etc/os-release
		// The standard format is key=value, not JSON usually, but many libs parse it or we can just grep.
		// Wait, the original code used json.Unmarshal on /etc/os-release?
		// /etc/os-release is NOT JSON. It is shell-compatible assignment.
		// The previous coder made a mistake thinking it acts like JSON or maybe just assumed it.
		// I will implement a simple parser for KEY=VALUE.

		data, err := os.ReadFile("/etc/os-release")
		if err != nil {
			// Fallback: check for common PMs directly
			detectCommonLinuxPMs()
			return
		}

		content := string(data)
		var id string
		lines := strings.Split(content, "\n")
		for _, line := range lines {
			if strings.HasPrefix(line, "ID=") {
				id = strings.TrimPrefix(line, "ID=")
				id = strings.Trim(id, "\"")
				break
			}
		}

		if id != "" {
			if val, ok := distro_pm[id]; ok {
				// check if actually installed
				if okP, path := isInstalled(val); okP {
					pm = packageManager{Name: val, Path: path}
					return
				}
			}
		}

		// If os-release didn't give us a working one, try fallback
		detectCommonLinuxPMs()

	default:
		fmt.Printf("Unknown operating system: %s\n", operatingSystem)
	}
}

func detectCommonLinuxPMs() {
	// check order: apt, dnf, pacman, zypper, yum, apk...
	checks := []string{"apt", "dnf", "pacman", "zypper", "yum", "apk", "xbps-install", "emerge", "nix-env"}
	for _, p := range checks {
		wrapperName := p
		if p == "xbps-install" {
			wrapperName = "xbps"
		}

		if ok, path := isInstalled(p); ok {
			pm = packageManager{Name: wrapperName, Path: path}
			// Map xbps-install back to "xbps" for key lookup if needed
			if wrapperName == "xbps" {
				// Our commands key is "xbps"
			}
			return
		}
	}
}

func isInstalled(pkg string) (bool, string) {
	path, err := exec.LookPath(pkg)
	if errors.Is(err, exec.ErrNotFound) {
		return false, ""
	}
	return true, path
}

func executeCommand(template string, pkgName string) {
	if template == "" {
		fmt.Println("Command not defined for this package manager.")
		return
	}

	// Use regex to replace isolated "x" only (respecting word boundaries)
	// This handles "xbps-install" (no replace) and "nixpkgs.x" (replace because . is non-word)
	// and "install x" (replace).
	// \b matches at word boundary.
	re := regexp.MustCompile(`\bx\b`)

	// Use ReplaceAllStringFunc to avoid interpreting $ in pkgName
	cmdStr := re.ReplaceAllStringFunc(template, func(s string) string {
		return pkgName
	})

	parts := strings.Fields(cmdStr)
	if len(parts) == 0 {
		return
	}

	head := parts[0]
	args := parts[1:]

	cmd := exec.Command(head, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		// The command itself might have failed (exit status != 0)
		// We just exit with the same code if possible or just log
		fmt.Printf("Error execution command: %v\n", err)
		os.Exit(1)
	}
}
