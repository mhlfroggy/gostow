# gostow (gstow)

gostow is an improved dotfile management utility written in Go which builds on
the strong base that already exists in GNU Stow. However, GNU Stow is a very
general-purpose package which isn't designed only for dotfile management, but
also for managing packages. Gostow approaches this problem with a more suckless
philiosphy of "doing one thing and doing it well".

Licensed under the MIT license.

### CHANGELOG
See [[changelog.md]] for version release history.

### Installation
Currently, only Arch Linux and macOS are supported

Arch Linux: `paru -S gostow-bin` or `yay -Sy gostow-bin`
macOS: `brew install gostow`

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

### LLM Disclosure
The source code for this project was not written with the use of LLMs, but
packaging scripts such as for AUR, Brew, etc. were written with the assistance
of Kimi K2.7 code.
