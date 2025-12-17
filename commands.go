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
	"i": {
		Name:          "i",
		Install:       "i install x",
		Uninstall:     "i uninstall x",
		Upgrade:       "i upgrade x",
		Search:        "i search x",
		Info:          "i info x",
		UpgradeAll:    "i upgrade",
		ListInstalled: "i list",
	},
	"apt": {
		Name:          "apt",
		Install:       "apt install x",
		Uninstall:     "apt remove x",
		Upgrade:       "apt install --only-upgrade x",
		Search:        "apt search x",
		Info:          "apt show x",
		UpgradeAll:    "apt upgrade",
		ListInstalled: "apt list --installed", // apt list -i
	},
	"brew": {
		Name:          "brew",
		Install:       "brew install x",
		Uninstall:     "brew uninstall x",
		Upgrade:       "brew upgrade x",
		Search:        "brew search x",
		Info:          "brew info x",
		UpgradeAll:    "brew upgrade",
		ListInstalled: "brew list",
	},
	"flatpak": {
		Name:          "flatpak",
		Install:       "flatpak install x",
		Uninstall:     "flatpak uninstall x",
		Upgrade:       "flatpak update x",
		Search:        "flatpak search x",
		Info:          "flatpak info x",
		UpgradeAll:    "flatpak update",
		ListInstalled: "flatpak list",
	},
	"snap": {
		Name:          "snap",
		Install:       "snap install --classic x", // --classic or not ?
		Uninstall:     "snap remove x",
		Upgrade:       "snap refresh x",
		Search:        "snap find x",
		Info:          "snap info x",
		UpgradeAll:    "snap refresh",
		ListInstalled: "snap list",
	},
	"dnf": {
		Name:          "dnf",
		Install:       "dnf install -y x",
		Uninstall:     "dnf remove -y x",
		Upgrade:       "dnf upgrade -y x",
		Search:        "dnf search x",
		Info:          "dnf info x",
		UpgradeAll:    "dnf upgrade -y",
		ListInstalled: "dnf list installed",
	},
	"pacman": {
		Name:          "pacman",
		Install:       "pacman -S --noconfirm x",
		Uninstall:     "pacman -Rs --noconfirm x",
		Upgrade:       "pacman -Syu --noconfirm x", // Upgrade specific pkg and system? Usually just -S to reinstall/upgrade specific
		Search:        "pacman -Ss x",
		Info:          "pacman -Qi x",
		UpgradeAll:    "pacman -Syu --noconfirm",
		ListInstalled: "pacman -Q",
	},
	"yum": {
		Name:          "yum",
		Install:       "yum install -y x",
		Uninstall:     "yum remove -y x",
		Upgrade:       "yum update -y x",
		Search:        "yum search x",
		Info:          "yum info x",
		UpgradeAll:    "yum update -y",
		ListInstalled: "yum list installed",
	},
	"zypper": {
		Name:          "zypper",
		Install:       "zypper install -n x",
		Uninstall:     "zypper remove -n x",
		Upgrade:       "zypper update -n x",
		Search:        "zypper search x",
		Info:          "zypper info x",
		UpgradeAll:    "zypper update -n",
		ListInstalled: "zypper se --installed-only",
	},
	"apk": {
		Name:          "apk",
		Install:       "apk add x",
		Uninstall:     "apk del x",
		Upgrade:       "apk add --upgrade x",
		Search:        "apk search x",
		Info:          "apk info x",
		UpgradeAll:    "apk upgrade",
		ListInstalled: "apk info",
	},
	"xbps": {
		Name:          "xbps",
		Install:       "xbps-install -y x",
		Uninstall:     "xbps-remove -y x",
		Upgrade:       "xbps-install -u x",
		Search:        "xbps-query -Rs x",
		Info:          "xbps-query -R x", // Remote info? or local -f? assuming remote
		UpgradeAll:    "xbps-install -Suy",
		ListInstalled: "xbps-query -l",
	},
	"emerge": {
		Name:          "emerge",
		Install:       "emerge x",
		Uninstall:     "emerge -C x",
		Upgrade:       "emerge -u x",
		Search:        "emerge -s x",
		Info:          "emerge -S x",
		UpgradeAll:    "emerge -uDN @world",
		ListInstalled: "qlist -I", // needs portage-utils potentially
	},
	"nix-env": {
		Name:          "nix-env",
		Install:       "nix-env -iA nixpkgs.x",
		Uninstall:     "nix-env -e x",
		Upgrade:       "nix-env -u x",
		Search:        "nix-env -qaP x",
		Info:          "nix-env -qa --description x",
		UpgradeAll:    "nix-env -u",
		ListInstalled: "nix-env -q",
	},
	"pkg": {
		Name:          "pkg",
		Install:       "pkg install -y x",
		Uninstall:     "pkg delete -y x",
		Upgrade:       "pkg upgrade -y x",
		Search:        "pkg search x",
		Info:          "pkg info x",
		UpgradeAll:    "pkg upgrade -y",
		ListInstalled: "pkg info",
	},
	"winget": {
		Name:          "winget",
		Install:       "winget install x",
		Uninstall:     "winget uninstall x",
		Upgrade:       "winget upgrade x",
		Search:        "winget search x",
		Info:          "winget show x",
		UpgradeAll:    "winget upgrade",
		ListInstalled: "winget list",
	},
	"scoop": {
		Name:          "scoop",
		Install:       "scoop install x",
		Uninstall:     "scoop uninstall x",
		Upgrade:       "scoop update x",
		Search:        "scoop search x",
		Info:          "scoop info x",
		UpgradeAll:    "scoop update",
		ListInstalled: "scoop list",
	},
	"choco": {
		Name:          "choco",
		Install:       "choco install x",
		Uninstall:     "choco uninstall x",
		Upgrade:       "choco upgrade x",
		Search:        "choco search x",
		Info:          "choco info x",
		UpgradeAll:    "choco upgrade",
		ListInstalled: "choco list",
	},
	// urpm
	// slackpkg
	// prt-get
	// pkgman

	// pkg (of termux)
	// opkg
	// eopkg

	// guix
	// cards
}
