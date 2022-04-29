# go-find

[![Documentation](https://godoc.org/github.com/jaytaylor/go-find?status.svg)](https://godoc.org/github.com/jaytaylor/go-find)
[![Build Status](https://travis-ci.org/jaytaylor/go-find.svg)](https://travis-ci.org/jaytaylor/go-find)
[![Report Card](https://goreportcard.com/badge/jaytaylor/go-find)](https://goreportcard.com/report/jaytaylor/go-find)

### [go-find](https://github.com/jaytaylor/go-find) is a programmatically accessible golang implementation of the *nix `find` command.

TL:DR; This thing can help you find files on disk matching a wide variety of criteria.

The goal of this project is to achieve equivalent or better capabilities
relative to the GNU find command-line utility.

## Get it

```bash
# Install the library:
go get -u github.com/jaytaylor/go-find

# Get the `go-find' CLI:
go install github.com/jaytaylor/go-find/go-find/...
```

## Example Usage

### As a standalone command-line program

```bash
go-find . -name 'favorite-quotes.txt'
```

### From within Go

```go
package main

import (
	"fmt"
	"strings"

	"github.com/jaytaylor/go-find"
)

func main() {
	finder := find.NewFind(".").MinDepth(1).Name("favorite-quotes.txt")
	hits, _ := finder.Evaluate()
	fmt.Printf("%+v\n", strings.Join(hits, "\n"))
}
```

## The predicate feature matrix

_"Predicate tests"_ are find's terminology for what can be intuitively be thought
of as filter operations.

Even though most of the predicate tests aren't yet implemented, mostly because
many of them are quite obscure, the initial set covers all of  my "everyday"
common use-cases.

Comparison of `go-find` vs `GNU find`:

| Feature     | GNU find has it? | go-find has it? |
| ----------- | ---------------- | --------------- |
| operators   | ✅                |                   | 
| -amin       | ✅                |                   |
| -anewer     | ✅                |                   |
| -atime      | ✅                |                   |
| -cmin       | ✅                |                   |
| -cnewer     | ✅                |                   |
| -context    | ✅                |                   |
| -ctime      | ✅                |                   |
| -empty      | ✅                | ✅                 |
| -executable | ✅                |                   |
| -false      | ✅                |                   |
| -fls        | ✅                |                   |
| -fprint     | ✅                |                   |
| -fprint0    | ✅                |                   |
| -fprintf    | ✅                |                   |
| -fstype     | ✅                |                   |
| -gid        | ✅                |                   |
| -group      | ✅                |                   |
| -ilname     | ✅                |                   |
| -iname      | ✅                |                   |
| -inum       | ✅                |                   |
| -ipath      | ✅                |                   |
| -iregex     | ✅                |                   |
| -iwholename | ✅                |                   |
| -links      | ✅                |                   |
| -lname      | ✅                |                   |
| -ls         | ✅                |                   |
| -maxdepth   | ✅                | ✅                 |
| -mindepth   | ✅                | ✅                 |
| -mmin       | ✅                |                   |
| -mount      | ✅                |                   |
| -mtime      | ✅                |                   |
| -name       | ✅                | ✅                 |
| -newer      | ✅                |                   |
| -newerXY    | ✅                |                   |
| -nogroup    | ✅                |                   |
| -not        | ✅                |                   |
| -nouser     | ✅                |                   |
| -ok         | ✅                |                   |
| -okdir      | ✅                |                   |
| -path       | ✅                |                   |
| -perm       | ✅                |                   |
| -print      | ✅                |                   |
| -print0     | ✅                | ✅                 |
| -printf     | ✅                |                   |
| -prune      | ✅                |                   |
| -quit       | ✅                |                   |
| -readable   | ✅                |                   |
| -regex      | ✅                | ✅                 |
| -regextype  | ✅                |                   |
| -samefile   | ✅                |                   |
| -size       | ✅                |                   |
| -true       | ✅                |                   |
| -type       | ✅                | ✅                 |
| -uid        | ✅                |                   |
| -used       | ✅                |                   |
| -user       | ✅                |                   |
| -wholename  | ✅                | ✅                 |
| -writable   | ✅                |                   |
| -xdev       | ✅                |                   |
| -xtype      | ✅                |                   |

Note: Features involving command execution are an anti-goal and are omitted from the matrix.   If you spot an error in this list, please open an issue or PR.
