# i the abstraction over all package managers

__i__ stands for __install__. You can use `i` to install any software via any package manager.

THIS PROJECT IS **DISCONTINUED** BECAUSE THERE IS "[UPT](https://github.com/sigoden/upt)" AND **NO NEED TO REINVENT THE WHEEL**.

TL;DR i is an abstraction over all other package managers available on Linux/MacOS.

This software is a work in progress - WIP - a.k.a It is not done yet.

## Why ?

If you used to apt or brew, and need to use dnf or swupd it is hard sometimes. But if you use __i__, it will be always the same.

## i wizard

- search for the_term
- show search results
- choose the package/app/program
- use default package manager for the distro, or suggest another one to the user to choose (install via blah-blah or do not install blah-blah-app).
- run the needed commands, give realtime feedback to user.

## Available As

- CLI : `i vim` to istall __vim__.
- GUI : `vim` then click search.

## Plan: commands

| command  | meaning |
|:--------|:--------|
| `i search x` | search for x |
| `i install x` | install x |
| `i add x` | install x |
| `i uninstall x` | uninstall x |
| `i remove x` | uninstall x |
| `i reinstall x` | uninstall x, then install it |
| `i info x` | show info about x |
| `i upgrade x` | upgrade x to the newer version if available |
| `i update x` | upgrade x to the newer version if available |
| `i up x` | upgrade x to the newer version if available |
| `i upgrade` | upgrade all to the newer version if available |
| `i update` | upgrade all to the newer version if available |
| `i up` | upgrade all to the newer version if available |
| `i updateable` | list all upgradeable apps/programs |
| `i updatable` | list all upgradeable apps/programs |
| `i upgradeable` | list all upgradeable apps/programs |
| `i upgradable` | list all upgradeable apps/programs |

\* updating the local index of packages is always run first. No need to run it manually.

```sh
$ i search x # search for x
searching for x ...

$ i install x # install x
installing x ...

$ i add x # install x
installing x ...

$ i uninstall x # uninstall x
uninstalling x ...

$ i remove x # uninstall x
uninstalling x ...

$ i reinstall x # uninstall x, then install it
uninstalling x ... DONE
installing x ... DONE

$ i info x # show info about x
x is blah blah blah

$ i upgrade x # upgrade x to the newer version if available
upgrading x from v1.0.0 to v1.1.0

$ i update x # upgrade x to the newer version if available
upgrading x from v1.0.0 to v1.1.0

$ i up x # upgrade x to the newer version if available
upgrading x from v1.0.0 to v1.1.0

$ i upgrade # upgrade all to the newer version if available
upgrading x from v1.0.0 to v1.1.0
upgrading y from v0.1.0 to v0.6.0
upgrading z from v1.3.0 to v1.3.2

$ i update # upgrade all to the newer version if available
upgrading x from v1.0.0 to v1.1.0
upgrading y from v0.1.0 to v0.6.0
upgrading z from v1.3.0 to v1.3.2


$ i up # upgrade all to the newer version if available
upgrading x from v1.0.0 to v1.1.0
upgrading y from v0.1.0 to v0.6.0
upgrading z from v1.3.0 to v1.3.2

$ i updateable # list all upgradeable apps/programs
x v1.2.1 >> v1.2.2
y v0.5.1 >> v1.0.2
z v1.2.3 >> v2.0.0

$ i updatable # list all upgradeable apps/programs
x v1.2.1 >> v1.2.2
y v0.5.1 >> v1.0.2
z v1.2.3 >> v2.0.0

$ i upgradeable # list all upgradeable apps/programs
x v1.2.1 >> v1.2.2
y v0.5.1 >> v1.0.2
z v1.2.3 >> v2.0.0

$ i upgradable # list all upgradeable apps/programs
x v1.2.1 >> v1.2.2
y v0.5.1 >> v1.0.2
z v1.2.3 >> v2.0.0

# updating the local index of packages is always run first. No need to run it manually.
```

## Plan: Arguments

### verbose output

```sh
$ i install z
installing z ... DONE

$ i install -v z
using homebrew to install z
    $ brew install z
installing z ... DONE

$ i install --verbose z
using homebrew to install z
    $ brew install z
installing z ... DONE
```

### specify a package manager

```sh
$ i install y
installing y ... DONE

$ i install --brew y
installing y via homebrew ... DONE

$ i install --apt z
installing z via apt ... DONE
```

### no output/script

```sh
$ i install a # nothing will be returned if successful; a.k.a os.Exit(0)
$
```

## Plan: supported package managers

| package manager    | exec | Operating Systems    | status |
|:------------------:|:----:|:--------------------:|:------:|
| brew (Homebrew)    |  1   | MacOS, Linux, BSD    |  WIP   |
| port (MacPorts)    |  1   | MacOS                |  ---   |
| apt                |  1   | Linux (Debian-based) |  ---   |
| apt-get            |  1   | Linux (Debian-based) |  ---   |
| dnf                |  1   | Linux (Fedora)       |  ---   |
| nix                |  1   | Linux, NixOS         |  ---   |
| pacman             |  1   | Linux                |  ---   |
| swupd              |  1   | Linux                |  ---   |
| rpm                |  1   | Linux                |  ---   |
| snap               |  2   | Linux                |  ---   |
| flatpak            |  2   | Linux                |  ---   |
| pkgsrc             |  2   | Linux                |  ---   |
| winget             |  2   | Windows              |  ---   |
| choco (Chocolatey) |  2   | Windows              |  ---   |
| go                 |  3   | language-based pm    |  ---   |
| cargo              |  3   | language-based pm    |  ---   |
| python             |  3   | language-based pm    |  ---   |

\* `exec` stands for __execution priority__.
\* `pm` stands for __package manager__.

## Plan: homebrew

```sh
brew search TEXT|/REGEX/
brew info [FORMULA|CASK...]
brew install FORMULA|CASK...
brew update
brew upgrade [FORMULA|CASK...]
brew uninstall FORMULA|CASK...
brew list [FORMULA|CASK...]
```

## tasks

- [x] check if the package/app/program is already installed and executable/callable
- [x] support search via homebrew
