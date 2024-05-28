package main

type commands struct {
	Name          string
	Install       string
	Uninstall     string
	Upgrade       string
	Search        string
	Info          string
	UpgradeAll    string
	ListInstalled string
}

var pm_commands = map[string]commands{
	"i": commands{
		Name:          "i",
		Install:       "i install x",
		Uninstall:     "i uninstall x",
		Upgrade:       "i upgrade x",
		Search:        "i search x",
		Info:          "i info x",
		UpgradeAll:    "i upgrade",
		ListInstalled: "i list",
	},
	"apt": commands{
		Name:          "apt",
		Install:       "apt install x",
		Uninstall:     "apt remove x",
		Upgrade:       "apt install --only-upgrade x",
		Search:        "apt search x",
		Info:          "apt show x",
		UpgradeAll:    "apt upgrade",
		ListInstalled: "apt list --installed", // apt list -i
	},
	"brew": commands{
		Name:          "brew",
		Install:       "brew install x",
		Uninstall:     "brew uninstall x",
		Upgrade:       "brew upgrade x",
		Search:        "brew search x",
		Info:          "brew info x",
		UpgradeAll:    "brew upgrade",
		ListInstalled: "brew list",
	},
	"flatpak": commands{
		Name:          "flatpak",
		Install:       "flatpak install x",
		Uninstall:     "flatpak uninstall x",
		Upgrade:       "flatpak update x",
		Search:        "flatpak search x",
		Info:          "flatpak info x",
		UpgradeAll:    "flatpak update",
		ListInstalled: "flatpak list",
	},
	"snap": commands{
		Name:          "snap",
		Install:       "snap install --classic x", // --classic or not ?
		Uninstall:     "snap remove x",
		Upgrade:       "snap refresh x",
		Search:        "snap find x",
		Info:          "snap info x",
		UpgradeAll:    "snap refresh",
		ListInstalled: "snap list",
	},
	// dnf
	// pacman
	// winget
	// scoop
	// choco
	// zypper
	// yum
	// xbps
	// urpm
	// slackpkg
	// prt-get
	// pkgman

	// nix-env

	// pkg //?
	// pkg (of termux)
	// opkg
	// eopkg

	// guix
	// emerge
	// cards

	// apk
}
