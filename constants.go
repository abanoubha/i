package main

const version = "25.12.20"

var distro_pm = map[string]string{
	"clearlinux": "swupd",

	// Debian-based => apt, apt-get
	"ubuntu":     "apt",
	"debian":     "apt",
	"linuxmint":  "apt",
	"pop":        "apt",
	"deepin":     "apt",
	"elementary": "apt",
	"raspbian":   "apt",
	"kali":       "apt", // security
	"parrot":     "apt", // security
	"aosc":       "apt",
	"zorin":      "apt",
	"devuan":     "apt",
	"bodhi":      "apt",
	"lxle":       "apt",
	"sparky":     "apt",
	"armbian":    "apt",
	"antix":      "apt",
	"lite":       "apt", // Linux Lite
	"linuxfx":    "apt",
	"endless":    "flatpak", // immutable, rely on flatpak to get apps
	// Redhat => dnf
	"fedora":    "dnf", // "yum",
	"redhat":    "dnf", // "yum",
	"rhel":      "dnf", // "yum",
	"centos":    "dnf", // "yum",
	"rocky":     "dnf", // "yum",
	"amzn":      "dnf", // "yum",
	"ol":        "dnf", // "yum",
	"almalinux": "dnf", // "yum",
	"qubes":     "dnf", // "yum",
	"eurolinux": "dnf", // "yum",
	"oubes":     "dnf", // "yum",

	"oracle":   "rpm",
	"sailfish": "rpm",

	// Arch-based => pacman
	"arch":        "pacman",
	"manjaro":     "pacman",
	"endeavouros": "pacman",
	"arcolinux":   "pacman",
	"garuda":      "pacman",
	"antergos":    "pacman",
	"kaos":        "pacman",
	"archbang":    "pacman",
	"artix":       "pacman", // Artix Linux
	// apk
	"alpine":     "apk",
	"postmarket": "apk",
	// zypper
	"opensuse":            "zypper",
	"opensuse-leap":       "zypper",
	"opensuse-tumbleweed": "zypper",
	// nix
	"nixos": "nix-env",
	// emerge
	"gentoo": "emerge",
	"funtoo": "emerge",
	// xps
	"void": "xbps",
	// urpm
	"mageia": "urpm",

	// slackpkg
	"slackware": "slackpkg",
	// eopkg
	"solus": "eopkg",
	// opkg
	"openwrt": "opkg",
	// cards
	"nutyx": "cards",
	// prt-get
	"crux": "prt-get",
	// pkg
	"freebsd":  "pkg",
	"ghostbsd": "pkg",

	"android": "pkg", // in termux

	"haiku": "pkgman",
}
