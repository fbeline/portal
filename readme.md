# [p]ortal [![Build Status](https://travis-ci.org/fbeline/portal.svg?branch=master)](https://travis-ci.org/fbeline/portal)

portal is a file system navigation tool inspired on python project [autojump](https://github.com/wting/autojump).

portal record your terminal usage and fast jump to your desired location.

![portal-demo](https://user-images.githubusercontent.com/5730881/36635707-9abe098e-1998-11e8-970c-983e22e3289d.gif)

## prerequisite

- go
- bash or ZSH

## Installation

```
go get github.com/fbeline/portal
```
copy the following code to your bashrc or zshrc file.

### bash

```bash
function p() {
	  local output="$(portal $1)"
    if [[ -d "${output}" ]]; then
        cd "${output}"
    else
        echo "portal: directory '${@}' not found"
        echo "Try \`portal --help\` for more information."
        false
    fi
}
PROMPT_COMMAND='portal s $(pwd)'
```

### zsh

```bash
function p() {
	  local output="$(portal $1)"
    if [[ -d "${output}" ]]; then
        cd "${output}"
    else
        echo "portal: directory '${@}' not found"
        echo "Try \`portal --help\` for more information."
        false
    fi
}
preexec() { portal s $(pwd) }
```

## Usage

Jump to a directory that contains `down`

```
p down
```

List saved paths and the relative scores

```
portal l
```

for more information:

```
portal --help
```

## License
MIT

