# Fixed Point Integers for Go

Fixed point integers for scientific and mathematical applications.

The following fixed point representations are implemented:

* int32_32

Licensed under the Unlicense (https://unlicense.org/)

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

This package is intended for efficient and fast computations. There are no plans to support applications which require constant-time cryptographic security.

## Contributions

We welcome the following contributions:

* Documentation
* Unit tests and benchmarking
* Faster algorithm implementations
* Suggested additions to the API
* Other fixed integer implementations (e.g. int52_12 and int26_6)

## Contributor Guidelines

When possible, contributions to this repository should follow the git workflow:

https://nvie.com/posts/a-successful-git-branching-model/

The `master` and `develop` branches are protected 
