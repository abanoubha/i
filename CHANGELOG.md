# change log for i

## v0.1.0

- barebone Go project

## v25.12.17

- feature: specify a package manager in the command (e.g. `i --brew vim`)
- feature: specify a package manager by alias or symlink.
  - For example, alias i to apt like this `alias apt=/usr/bin/i`, then use `apt info vim` to see info about vim using the apt package manager through the i alias.
  - Or create a symlink like this `ln -s /usr/bin/i /usr/local/bin/apt`, then use `apt info vim` to see info about vim using the apt package manager through the i symlink.   
- feature: verbose output (e.g. `i -v vim` or `i --verbose vim`)

## next 

- latest GitHub release binary: create a simple Bash script to get the latest binary/executable from GitHub releases and install it and add it to the PATH.
