# [p]ortal

portal is a file system navigation tool.
portal record your terminal usage and fast jump to your require location.

![portal-usage](https://user-images.githubusercontent.com/5730881/36635478-12970294-1994-11e8-9bb2-ef1e0b6cfc06.gif)

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

