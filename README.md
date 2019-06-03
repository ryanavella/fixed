[![](https://godoc.org/github.com/ryanavella/fixed?status.svg)](https://godoc.org/github.com/ryanavella/fixed) [![Go Report Card](https://goreportcard.com/badge/github.com/ryanavella/fixed)](https://goreportcard.com/report/github.com/ryanavella/fixed) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/ryanavella/fixed/blob/develop/LICENSE-MIT) [![License: Unlicense](https://img.shields.io/badge/license-Unlicense-blue.svg)](https://github.com/ryanavella/fixed/blob/develop/LICENSE-UNLICENSE)

# Fixed

Fixed point integers for Go.

The following fixed point representations are implemented:

* int32_32

Fixed is free and open source software distributed under the terms of both the MIT License and the Unlicense.

## Installing

```shell
go get github.com/ryanavella/fixed
```

## Usage

```golang
package main

import (
	"fmt"

	"github.com/ryanavella/fixed/int32_32"
)

func main() {
	a := int32_32.FromInt32(-1)    // -1
	b := int32_32.FromFloat64(1.5) // +1.5
	fmt.Println(a+b, a-b, a.Mul(b))
}
```

## Scope

This package is intended for efficient and fast computations (i.e. for scientific and mathematical applications). There are no plans to support applications which require constant-time cryptographic security.

## Contributions

See [contributor guidelines](CONTRIBUTING.md).
