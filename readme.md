# i the abstraction over all package managers

__i__ stands for __install__. You can use `i` to install any software via any package manager.

TL;DR i is an abstraction over all other package managers available on Linux/MacOS.

This software is a work in progress - WIP - a.k.a It is not done yet.

## Why ?

If you used to **apt** or **brew**, and need to use **dnf** or **swupd**. It is hard sometimes. But if you use __i__, it will be always the same.

## Install 'i'

Use Go to install i:

```sh
go install github.com/abanoubha/i@latest
```

Build the project from source:

```sh
# get project deps/libs, then build the binary/executable and call it 'i'
go mod tidy && go build -o i .

# run the program
./i
```

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
| port (MacPorts)    |  1   | MacOS                |  ---   |
| apt                |  1   | Linux (Debian-based) |  ✅    |
| apt-get            |  1   | Linux (Debian-based) |  apt   |
| dnf                |  1   | Linux (Fedora)       |  ✅    |
| nix-env            |  1   | Linux, NixOS         |  ✅    |
| pacman             |  1   | Linux                |  ✅    |
| swupd              |  1   | Linux                |  ---   |
| rpm                |  1   | Linux                |  ---   |
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
| winget             |  2   | Windows              |  ✅    |
| choco (Chocolatey) |  2   | Windows              |  ---   |

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

# upgrade a package
i upgrade vim
# or
i update vim
# or
i up vim

# upgrade all packages
i upgrade
# or
i update
# or
i up
```

You can add `--verbose` flag to get verbose output with more details like this:

```sh
$ sudo i --verbose install vim
Using package manager: apt
Executing: apt install vim
Reading package lists... Done
Building dependency tree... Done
Reading state information... Done
vim is already the newest version (2:9.1.0016-1ubuntu7.9).
0 upgraded, 0 newly installed, 0 to remove and 44 not upgraded.
```

As you can see, `--verbose` flag will show more details about the installation process like these two lines in the output above.

```sh
Using package manager: apt
Executing: apt install vim
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

You can force `i` to use a specific package manager by **aliasing** `i` to the package manager name or by **symlinking** `i` to the package manager name. 

- aliasing `i` to a package manager:

```sh
alias apt=/usr/bin/i
```

- Or create a symlink:

```sh
ln -s /usr/bin/i /usr/local/bin/apt
```

So, you can use `apt install vim` to install vim using the apt package manager through the i alias/symlink.
