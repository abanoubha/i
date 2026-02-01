package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"sort"
	"strings"
	"time"
)

var (
	operatingSystem string
	quiet           bool = false
	forcedPM        string
	forcesh         bool = false
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
			case "--quiet", "--silent", "--compact", "-q":
				quiet = true
			case "--help", "-h":
				printUsage()
				return
			case "--version", "-v":
				fmt.Printf("i the installer v%v\n", version)
				return
			case "--forcesh", "-sh":
				forcesh = true
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

	if !quiet {
		fmt.Printf("[info] using package manager: %s\n", pm.Name)
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
		if !quiet {
			fmt.Println("[info] updating local index...")
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
				if !quiet {
					fmt.Printf("[info] upgrading packages for manager: %s\n", p.Name)
				}

				// If this is not the primary PM (which was already updated at start), update its index
				if p.Name != pm.Name && c.UpdateIndex != "" {
					if !quiet {
						fmt.Printf("[info] updating index for %s...\n", p.Name)
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
	case "uninstall", "remove", "rm", "un":
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
	case "help":
		printUsage()
		return
	case "version":
		fmt.Printf("i the installer v%v\n", version)
		return
	case "selfup", "selfupdate", "selfupgrade":
		if forcesh {
			const upgradeScript = "https://raw.githubusercontent.com/abanoubha/i/main/scripts/install.sh"
			fmt.Println("[info] Starting upgrade...")
			if err := streamToShell(upgradeScript); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			fmt.Println("[info] 'i' is upgraded successfully.")
		}
		installLatestVersion() // Go impl
	case "selfun", "selfuninstall", "selfdelete":
		if forcesh {
			const uninstallScript = "https://raw.githubusercontent.com/abanoubha/i/main/scripts/uninstall.sh"
			fmt.Println("[info] Starting self delete...")
			if err := streamToShell(uninstallScript); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			fmt.Println("[info] 'i' is deleted successfully.")
		}
		selfUninstall() // Go impl
	default:
		fmt.Printf("'%v' sub-command is not supported.\n", action)
	}
}

func printUsage() {
	fmt.Printf(`
i the abstraction over all package managers v%v
Usage:
i install vim			# install vim program
i add vim				# install vim program

i info vim				# show information about vim program
i show vim				# show information about vim program

i search vim			# search for vim program
i find vim				# search for vim program

i uninstall vim			# uninstall vim program from the system
i remove vim			# uninstall vim program from the system
i rm vim				# uninstall vim program from the system
i un vim				# uninstall vim program from the system

i install --quiet vim	# show less information while installing vim
i install --silent vim	# show less information while installing vim
i install --compact vim	# show less information while installing vim
i install -q vim		# show less information while installing vim

i uninstall --quiet vim		# show less information while removing vim
i uninstall --silent vim	# show less information while removing vim
i uninstall --compact vim	# show less information while removing vim
i uninstall -q vim			# show less information while removing vim

i install --apt vim		# use 'apt' to install vim program
i add --apt vim			# use 'apt' to install vim program

i install --snap vim	# use 'snap' to install vim program
i add --snap vim		# use 'snap' to install vim program

i install --flatpak vim	# use 'flatpak' to install vim program
i add --flatpak vim		# use 'flatpak' to install vim program

i install --pacman vim	# use 'pacman' to install vim program
i add --pacman vim		# use 'pacman' to install vim program

i --help				# show this information
i -h					# show this information

i --version				# show version number
i -v					# show version number

`, version)
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
		fmt.Println("[info] Windows support is experimental.")
		if ok, path := isInstalled("winget"); ok {
			detectedPMs = append(detectedPMs, packageManager{Name: "winget", Path: path})
		}
		if ok, path := isInstalled("choco"); ok {
			detectedPMs = append(detectedPMs, packageManager{Name: "choco", Path: path})
		}
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

	cmdStr := template
	// if template ends with ".x" or " x" remove "x" and add pkgName
	if strings.HasSuffix(template, ".x") || strings.HasSuffix(template, " x") {
		cmdStr = strings.TrimSuffix(template, "x") + pkgName
	}

	if after, ok := strings.CutPrefix(cmdStr, "sudo "); ok {
		cmdStr = after

		if !quiet {
			fmt.Printf("[info] executing: %s\n", cmdStr)
		}

		parts := strings.Fields(cmdStr)
		if len(parts) == 0 {
			return
		}

		if err := runAsSuperUser(parts...); err != nil {
			fmt.Println("[error] can not run the command because sudo/doas not found yet the command require super user privilege/permissions", err)
		}
		return
	}

	if !quiet {
		fmt.Printf("[info] executing: %s\n", cmdStr)
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
		if !quiet {
			fmt.Printf("[error] error executing command: %v\n", err)
		}
		os.Exit(1)
	}
}

// fetches a remote shell script and pipes it directly to sh.
func streamToShell(url string) error {
	client := &http.Client{
		Timeout: 60 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("failed to initiate request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned non-200 status: %d %s", resp.StatusCode, resp.Status)
	}

	cmd := exec.Command("sh")

	// pipe streams
	cmd.Stdin = resp.Body

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("script execution failed: %w", err)
	}

	return nil
}
