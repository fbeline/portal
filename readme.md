# [p]ortal [![Build Status](https://travis-ci.org/fbeline/portal.svg?branch=master)](https://travis-ci.org/fbeline/portal)

portal is a file system navigation tool inspired by [autojump](https://github.com/wting/autojump).

portal record your terminal usage and fast jump to your desired location.

![portal-demo](https://user-images.githubusercontent.com/5730881/36635707-9abe098e-1998-11e8-970c-983e22e3289d.gif)

## prerequisite

- go
- bash or ZSH

## Installation

```
go get github.com/fbeline/portal
```
copy the following code into your .bashrc or .zshrc file.

- [bash](portal.sh)
- [zsh](portal.zsh)

## Usage

*Directories must be visited first before you are able to jump.*

Jump to a directory that contains `doc`.

```
p doc
```

Jump to a Child directory that contains `test`.

```
pc test
```

Return the matched directory path without jumping to it.

```
portal doc
```

List saved paths and the relative scores.

```
portal l
```

for more information.

```
portal -h
```

## License
MIT

