# pscan

Simple tcp port scanner.

## Install

```sh
$ go get github.com/jandersonmartins/pscan
```

## Usage

```go
package main

import (
	"os"

	"github.com/jandersonmartins/pscan"
)

func main() {
	pscan.Pscan("localhost", 5, os.Stdout)
}
```