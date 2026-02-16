# change log for i

## v25.12.18

- feature: specify a package manager in the command (e.g. `i --brew vim`)
- feature: specify a package manager by alias or symlink.
  - For example, alias i to apt like this `alias apt=/usr/bin/i`, then use `apt info vim` to see info about vim using the apt package manager through the i alias.
  - Or create a symlink like this `ln -s /usr/bin/i /usr/local/bin/apt`, then use `apt info vim` to see info about vim using the apt package manager through the i symlink.
- feature: verbose output (e.g. `i -v vim` or `i --verbose vim`)
- check if the input args are valid (a-zA-Z0-9_-@)
- Supported package managers: apk, apt, brew, cards, choco, dnf, emerge, eopkg, flatpak, guix, nix-env, opkg, pacman, pkg, pkgman, prt-get, scoop, slackpkg, snap, urpm, winget, xbps, yum, zypper

## v25.12.19

- feature: support rpm, port (macports)
- refactor: efficient and more idiomatic Go code
- feature: check if the package/app/program is already installed and executable or callable. This is to prevent installing a package that is already installed via another package manager. For example, prevent installing vim by apt if vim is already installed by snap. (Conflict Resolution: Detecting if a tool is already installed by another manager)
- feature: update the local index of packages before attempting to install or update any package.
- feature: update/upgrade all packages will run **all** found package managers.
- feature: show all installed PMs using `pms` subcommand.
- feature: list all packages installed by all found package managers.

## v25.12.20

- fix: add `sudo` to some commands which need super user permissions.
- feature(`performance_test.go`): benchmarking performance of RegExp replacement vs string manipulation replacement.
- fix: use string manipulation instead of RegExp replacement for performance.

## v26.01.29

- better comprehensive help screen
- feature: use '--quiet' flag instead of '--verbose', so it will be verbose by default
- feature: use '-v' as a short flag for '--version', rewrite help screen for detailed help info
- feature: + use 'un' as a short for 'uninstall'

## v26.01.30

- experimental support for Microsoft Windows via winget and Choco(latey).

## v260201

- fix: make 'help' as proper subcommands for consistency
- fix: make 'version' as proper subcommands for consistency
- feature: upgrade the previously installed `i` program
- feature: delete/uninstall the previously installed `i` program
- feature: support `doas` along with `sudo`

## v260203

- feature: support upgrading windows version release
- fix: only download executable (i) if the new version is not installed
- ux: when upgrading, show old version (installed) and new version
- latest GitHub release binary: create a simple Bash script to get the latest binary/executable from GitHub releases and install it and add it to the PATH.

## working on

- upgrade to Go 1.26
- feature: add '--names-only' flag to APT search command

## next

- use `i i vim` as a fast command for `i install vim`
- TUI : type `vim` -> searching -> choosing it -> installing it.
