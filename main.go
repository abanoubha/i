package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"sort"
	"strings"
)

var (
	operatingSystem string
	verbose         bool
	forcedPM        string
)

type packageManager struct {
	Name string
	Path string
}

var pm packageManager
var detectedPMs []packageManager

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	// Parse arguments
	args := os.Args[1:]
	var action string
	var pkgName string

	// Simple custom parsing to handle flags mixed with args
	for i := range args {
		arg := args[i]
		if strings.HasPrefix(arg, "-") {
			switch arg {
			case "--verbose", "-v":
				verbose = true
			case "--help", "-h":
				printUsage()
				return
			case "--version":
				fmt.Printf("i the installer v%v\n", version)
				return
			default:
				// Check for specific PM flags (e.g., --apt, --brew)
				if after, ok := strings.CutPrefix(arg, "--"); ok {
					pmName := after
					// Verify if it's a known PM
					if _, ok := pm_commands[pmName]; ok {
						forcedPM = pmName
						continue
					}
				}
				fmt.Printf("Unknown flag: %s\n", arg)
				return
			}
		} else {
			if action == "" {
				action = arg
			} else if pkgName == "" {
				pkgName = arg
			} else {
				// Multiple packages or extra args?
				// For now, let's just append to pkgName or warn.
				// The original design seemed to handle one package.
				// Let's keep it simple for now, maybe append to allow "i install x y" later?
				// But existing logic uses pkgName as a single string replacement.
				fmt.Printf("Too many arguments: %s\n", arg)
				return
			}
		}
	}

	if action == "" {
		printUsage()
		return
	}

	// Detect OS and PM
	detectPM()

	if pm.Name == "" {
		fmt.Println("No supported package manager found.")
		os.Exit(1)
	}

	if verbose {
		fmt.Printf("Using package manager: %s\n", pm.Name)
	}

	if pkgName != "" {
		if !validateInput(pkgName) {
			fmt.Printf("Invalid package name: %s\n", pkgName)
			os.Exit(1)
		}
	}

	cmds, ok := pm_commands[pm.Name]
	if !ok {
		fmt.Printf("Configurations for package manager '%s' not found.\n", pm.Name)
		os.Exit(1)
	}

	// Check if we should update the index
	updateRequiredActions := map[string]bool{
		"install": true, "add": true,
		"update": true, "upgrade": true, "up": true,
		"search": true, "find": true,
	}

	if cmds.UpdateIndex != "" && updateRequiredActions[action] {
		if verbose {
			fmt.Println("Updating local index...")
		}
		executeCommand(cmds.UpdateIndex, "")
	}

	switch action {
	case "pmlist":
		var pms []string
		for k := range pm_commands {
			if k == "i" {
				continue
			}
			pms = append(pms, k)
		}
		sort.Strings(pms)
		fmt.Println("Supported package managers:")
		for _, pm := range pms {
			fmt.Println("- " + pm)
		}
		return
	case "pms":
		fmt.Println("Available package managers:")
		for _, p := range detectedPMs {
			fmt.Println("- " + p.Name)
		}
		return
	case "info", "show":
		if pkgName == "" {
			fmt.Println("No package specified.")
			return
		}
		executeCommand(cmds.Info, pkgName)
	case "update", "upgrade", "up":
		if pkgName == "" {
			// Upgrade all packages for all detected package managers
			fmt.Println("Upgrading all packages...")
			for _, p := range detectedPMs {
				c, ok := pm_commands[p.Name]
				if !ok {
					continue
				}
				if verbose {
					fmt.Printf("Upgrading packages for manager: %s\n", p.Name)
				}

				// If this is not the primary PM (which was already updated at start), update its index
				if p.Name != pm.Name && c.UpdateIndex != "" {
					if verbose {
						fmt.Printf("Updating index for %s...\n", p.Name)
					}
					executeCommand(c.UpdateIndex, "")
				}

				executeCommand(c.UpgradeAll, "")
			}
		} else {
			executeCommand(cmds.Upgrade, pkgName)
		}
	case "install", "add":
		if pkgName == "" {
			fmt.Println("No package specified.")
			return
		}
		if ok, path := isInstalled(pkgName); ok {
			fmt.Printf("Package '%s' is already installed at %s\n", pkgName, path)
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
		// Fallback to install for now, as existing code did
		fmt.Println("Reinstall not explicitly supported yet. Try install.")
	case "search", "find":
		if pkgName == "" {
			fmt.Println("No term specified to search.")
			return
		}
		executeCommand(cmds.Search, pkgName)
	case "list", "installed":
		for i, p := range detectedPMs {
			c, ok := pm_commands[p.Name]
			if !ok {
				continue
			}
			if i > 0 {
				fmt.Println()
			}
			fmt.Printf("Listing installed packages for %s:\n", p.Name)
			executeCommand(c.ListInstalled, "")
		}
	default:
		fmt.Printf("'%v' sub-command is not supported.\n", action)
	}
}

func printUsage() {
	fmt.Printf("i the abstraction over all package managers v%v\nUsage:\n  i install vim\n  i install --verbose vim\n  i install --apt vim\n  i info vim\n  i search vim\n  i uninstall vim\n", version)
}

func validateInput(input string) bool {
	// Allow a-z, A-Z, 0-9, _, -, @, ., +
	// Some packages have dots (e.g. python3.8) or plus (g++)
	match, _ := regexp.MatchString(`^[a-zA-Z0-9_\-@.+]+$`, input)
	return match
}

// detectCommonLinuxPMs appends all found supported package managers to detectedPMs.
func detectCommonLinuxPMs() {
	checks := []string{"apt", "dnf", "pacman", "snap", "flatpak", "zypper", "yum", "apk", "xbps-install", "emerge", "nix-env", "brew", "port", "winget", "choco", "scoop"}
	for _, p := range checks {
		wrapperName := p
		if p == "xbps-install" {
			wrapperName = "xbps"
		}
		if ok, path := isInstalled(p); ok {
			detectedPMs = append(detectedPMs, packageManager{Name: wrapperName, Path: path})
		}
	}
}

func detectPM() {
	if forcedPM != "" {
		pm = packageManager{Name: forcedPM, Path: ""}
		detectedPMs = append(detectedPMs, pm)
		return
	}

	// Check if binary name acts as an alias
	binName := filepath.Base(os.Args[0])
	if _, ok := pm_commands[binName]; ok && binName != "i" {
		pm = packageManager{Name: binName, Path: ""}
		detectedPMs = append(detectedPMs, pm)
		return
	}

	operatingSystem = runtime.GOOS
	switch operatingSystem {
	case "windows":
		fmt.Println("Windows support is minimal.")
		// Potential windows logic could be added here similar to linux/mac
	case "darwin":
		if ok, path := isInstalled("brew"); ok {
			detectedPMs = append(detectedPMs, packageManager{Name: "brew", Path: path})
		}
		if ok, path := isInstalled("port"); ok {
			detectedPMs = append(detectedPMs, packageManager{Name: "port", Path: path})
		}
	case "linux":
		// Try parsing /etc/os-release for ID
		id := getOSReleaseID()
		if id != "" {
			if val, ok := distro_pm[id]; ok {
				if okP, path := isInstalled(val); okP {
					detectedPMs = append(detectedPMs, packageManager{Name: val, Path: path})
					// Don't return, continue to check common ones for co-existing PMs (e.g. invalidating assumption that we only have one)
					// Actually, distro_pm often maps to the MAIN system PM.
					// We should probably check common ones too, but deduplicate?
					// For now, let's keep the logic simple: verify distro PM, then check common ones.
				}
			}
		}
		detectCommonLinuxPMs()

	default:
		fmt.Printf("Unknown operating system: %s\n", operatingSystem)
	}

	// Deduplicate detectedPMs based on Name
	uniquePMs := make([]packageManager, 0, len(detectedPMs))
	seen := make(map[string]bool)
	for _, p := range detectedPMs {
		if !seen[p.Name] {
			seen[p.Name] = true
			uniquePMs = append(uniquePMs, p)
		}
	}
	detectedPMs = uniquePMs

	if len(detectedPMs) > 0 {
		pm = detectedPMs[0]
	}
}

func getOSReleaseID() string {
	data, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return ""
	}
	content := string(data)
	lines := strings.SplitSeq(content, "\n")
	for line := range lines {
		if after, ok := strings.CutPrefix(line, "ID="); ok {
			id := after
			return strings.Trim(id, "\"")
		}
	}
	return ""
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

	re := regexp.MustCompile(`\bx\b`)
	cmdStr := re.ReplaceAllStringFunc(template, func(s string) string {
		return pkgName
	})

	if verbose {
		fmt.Printf("Executing: %s\n", cmdStr)
	}

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
		if verbose {
			fmt.Printf("Error executing command: %v\n", err)
		}
		os.Exit(1)
	}
}
