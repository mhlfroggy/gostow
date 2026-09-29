# gostow (gstow)

gostow is an improved dotfile management utility written in Go which builds on
the strong base that already exists in GNU Stow. However, GNU Stow is a very
general-purpose package which isn't designed only for dotfile management, but
also for managing packages. Gostow approaches this problem with a more suckless
philiosphy of "doing one thing and doing it well".

Licensed under the MIT license.

### CHANGELOG
See changelog.md for version release history.

### Installation
From source:
```aiignore bash
git clone https://github.com/mattheenan/gostow.git
cd gostow
go build
go install
```

Additionally, packages exist for the main Linux distributions as well as MacOS
- Debian/Ubuntu: `sudo apt install gostow`
- Fedora/RHEL: `sudo dnf install gostow`
- Arch Linux: `paru -S gostow`
- Homebrew: `brew install gostow`

### Documentation

### Why should I use gostow instead of GNU Stow?
- gostow is a single binary (under 200loc), meaning that there are fewer
  dependenceies
- Platform agnostic: Gostow can run on Windows, MacOS, Linux, and BSD-like
  systems
- Better safety handles: Gostow offers a more robust `--dry-run` flag which will
  not affect your filesystem, and automatically aborts any attempted linking if
  files already exist in the target directory unless the `--override` flag is
  passed in
- Better output to show what is going on