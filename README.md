# i the abstraction over all package managers

__i__ stands for __install__. You can use `i` to install any software via any package manager.

TL;DR i is an abstraction over all other package managers available on Linux/MacOS.

> [!NOTE]
> [lazyinstaller](https://github.com/abanoubha/lazyinstaller) is the TUI version of the same abstraction concept.

## Why ?

If you used to __apt__ or __brew__, and need to use __dnf__ or __swupd__. It is hard sometimes. But if you use __i__, it will be always the same.

## Install 'i'

Run the installation script in the terminal like this:

```sh
curl -fsSL https://raw.githubusercontent.com/abanoubha/i/main/scripts/install.sh | sh
```

Or use Go to install i:

```sh
go install github.com/abanoubha/i@latest
```

If you cloned/downloaded project source code, use the included installation script in `scripts/install.sh` like this:

```sh
sh scripts/install.sh
```

Or build the project from source:

```sh
# get project deps/libs, then build the binary/executable and call it 'i'
go mod tidy && go build -o i .

# run the program
./i
```

## Uninstall 'i' tool

Run the uninstall script in the terminal like this:

```sh
curl -fsSL https://raw.githubusercontent.com/abanoubha/i/main/scripts/uninstall.sh | sh
```

It will inform you that the i tool is removed.

## supported package managers

You can list all supported package managers using `i pmlist`:

```sh
$ i pmlist
Supported package managers:
- apk
- apt
- brew
- dnf
- emerge
- flatpak
- nix-env
- pacman
- pkg
- snap
- xbps
- yum
- zypper
```

| package manager    | exec | Operating Systems    | status |
|:------------------:|:----:|:--------------------:|:------:|
| brew (Homebrew)    |  1   | MacOS, Linux, BSD    |  ✅    |
| port (MacPorts)    |  1   | MacOS                |  ✅    |
| apt                |  1   | Linux (Debian-based) |  ✅    |
| apt-get            |  1   | Linux (Debian-based) |  apt   |
| dnf                |  1   | Linux (Fedora)       |  ✅    |
| nix-env            |  1   | Linux, NixOS         |  ✅    |
| pacman             |  1   | Linux                |  ✅    |
| rpm                |  1   | Linux                |  ✅    |
| emerge             |  1   | Linux                |  ✅    |
| zypper             |  1   | Linux                |  ✅    |
| apk                |  1   | Linux                |  ✅    |
| xbps               |  1   | Linux                |  ✅    |
| snap               |  2   | Linux                |  ✅    |
| flatpak            |  2   | Linux                |  ✅    |
| pkg                |  2   | Linux                |  ✅    |
| yum                |  2   | Linux                |  ✅    |
| scoop              |  2   | Linux                |  ✅    |
| pkgsrc             |  2   | Linux                |  ---   |
| urpm               |  2   | Linux                |  ✅    |
| slackpkg           |  2   | Linux                |  ✅    |
| prt-get            |  2   | Linux                |  ✅    |
| pkgman             |  2   | Linux                |  ✅    |
| opkg               |  2   | Linux                |  ✅    |
| eopkg              |  2   | Linux                |  ✅    |
| guix               |  2   | Linux                |  ✅    |
| cards              |  2   | Linux                |  ✅    |
| winget             |  2   | Windows              |  ✅    |
| choco (Chocolatey) |  2   | Windows              |  ✅    |

\* `exec` stands for __execution priority__.
\* `pm` stands for __package manager__.

## How to use `i` the abstraction over all package managers

```sh
# install a package
i install vim
# or
i add vim

# search for a package
i search vim
# or
i find vim

# show info about a package
i info vim
# or
i show vim

# uninstall a package
i uninstall vim
# or
i remove vim
# or
i rm vim
# un
i un vim

# upgrade a package
i upgrade vim
# or
i update vim
# or
i up vim

# upgrade all packages installed by all found package managers
i upgrade
# or
i update
# or
i up
```

You can add `--quiet` flag to get less verbose output with less details like this:

```sh
$ sudo i --quiet install vim
Reading package lists... Done
Building dependency tree... Done
Reading state information... Done
vim is already the newest version (2:9.1.0016-1ubuntu7.9).
0 upgraded, 0 newly installed, 0 to remove and 44 not upgraded.
```

As you can see, `--quiet` flag will show less details about the installation process.

Show all installed/found package managers on your system:

```sh
$ i pms
Available package managers:
- apt
- snap
```

### Specify a package manager to use

Force `i` to use `apt` to install `vim`:

```sh
i --apt install vim
```

Force `i` to use `brew` to get info about `vim`:

```sh
i --brew info vim
```

### Force `i` to use a specific package manager

You can force `i` to use a specific package manager by __aliasing__ `i` to the package manager name or by __symlinking__ `i` to the package manager name.

- aliasing `i` to a package manager:

```sh
alias apt=/usr/bin/i
```

- Or create a symlink:

```sh
ln -s /usr/bin/i /usr/local/bin/apt
```

So, you can use `apt install vim` to install vim using the apt package manager through the i alias/symlink.

## build executables for all operating systems / platforms

Clone the 'i' project:

```sh
git clone --depth 1 -b main https://github.com/abanoubha/i.git
```

Go inside the project directory/folder:

```sh
cd i
```

To build executables for all supported platforms, operating systems, and architectures, run this command while you are inside the cloned/downloaded project source code:

```sh
sh scripts/build-all.sh
```

The executables/binaries for all supported OSes will be inside `./dist` directory/folder.

You can specify the version/release number (if you like):

```sh
sh scripts/build-all.sh v260130
```

Or you can run manual command for each OS/arch to build its compatible executable like these commands:

```sh
# linux 64 bit (current os)
go build -o i-linux-x64 .
# linux 64 bit (if not working on Linux distro)
GOOS=linux GOARCH=amd64 go build -o i-linux-x64 .

# windows 64 bit
GOOS=windows GOARCH=amd64 go build -o i-windows-x64.exe .

# macOS M-series
GOOS=darwin GOARCH=arm64 go build -o i-macos-apple-silicon .
# macOS intel 64 bit
GOOS=darwin GOARCH=amd64 go build -o i-macos-x64 .
```

## versioning

I use a simple date as a version number for releases. If I released a version on Jan 30, 2026 , I will set the version number as `260130` and named `v260130` (prefixed with v for version).

```plain
260130
^ ^ ^
y m d

26      01       30
^       ^        ^
year    month    day
```

## Are 'i' source code anywhere else?

yes, you can find 'i' project on:

- <https://github.com/abanoubha/i>
- <https://codeberg.org/abanoubha/i>
- <https://gitlab.com/abanoubha/i>
