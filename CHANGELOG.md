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

## next 

- update/upgrade all packages will run **all** found package managers.
- latest GitHub release binary: create a simple Bash script to get the latest binary/executable from GitHub releases and install it and add it to the PATH.
- TUI : type `vim` -> searching -> choosing it -> installing it.
- State Management: Tracking installed packages across all managers in a single manifest file (like a "super-package.json").
- Homebrew commands:
    ```sh
    brew search TEXT|/REGEX/
    brew info [FORMULA|CASK...]
    brew install FORMULA|CASK...
    brew update
    brew upgrade [FORMULA|CASK...]
    brew uninstall FORMULA|CASK...
    brew list [FORMULA|CASK...]
    ```
