package main

import (
	"regexp"
	"testing"
)

func TestCommandGeneration(t *testing.T) {
	pkgName := "testpkg"

	tests := []struct {
		pmName           string
		action           string // Install, Uninstall, Search, etc.
		expectedTemplate string
		expectedCmd      string
	}{
		{"dnf", "install", "dnf install -y x", "dnf install -y testpkg"},
		{"pacman", "install", "pacman -S --noconfirm x", "pacman -S --noconfirm testpkg"},
		{"yum", "search", "yum search x", "yum search testpkg"},
		{"zypper", "info", "zypper info x", "zypper info testpkg"},
		{"apk", "upgrade", "apk add --upgrade x", "apk add --upgrade testpkg"},
		{"xbps", "uninstall", "xbps-remove -y x", "xbps-remove -y testpkg"},
		{"nix-env", "install", "nix-env -iA nixpkgs.x", "nix-env -iA nixpkgs.testpkg"},
		// Add checks for multi-word commands or flags
		{"apt", "list", "apt list --installed", "apt list --installed"},
	}

	for _, tt := range tests {
		cmds, ok := pm_commands[tt.pmName]
		if !ok {
			t.Errorf("Package manager %s not found in pm_commands", tt.pmName)
			continue
		}

		var template string
		switch tt.action {
		case "install":
			template = cmds.Install
		case "uninstall":
			template = cmds.Uninstall
		case "search":
			template = cmds.Search
		case "info":
			template = cmds.Info
		case "upgrade":
			template = cmds.Upgrade
		case "list":
			template = cmds.ListInstalled
		}

		if template != tt.expectedTemplate {
			t.Errorf("[%s] %s template mismatch: got %q, want %q", tt.pmName, tt.action, template, tt.expectedTemplate)
		}

		re := regexp.MustCompile(`\bx\b`)
		// Use ReplaceAllStringFunc to avoid interpreting $ in pkgName
		cmdStr := re.ReplaceAllStringFunc(template, func(s string) string {
			return pkgName
		})

		// Fix expectation for list if we passed empty string
		expected := tt.expectedCmd
		if tt.action == "list" {
			expected = "apt list --installed" // no change
		}

		if cmdStr != expected {
			t.Errorf("[%s] %s command mismatch: got %q, want %q", tt.pmName, tt.action, cmdStr, expected)
		}
	}
}
