# Compactor [![Build Status](https://travis-ci.com/joseluisq/compactor.svg?branch=master)](https://travis-ci.com/joseluisq/compactor) [![codecov](https://codecov.io/gh/joseluisq/compactor/branch/master/graph/badge.svg)](https://codecov.io/gh/joseluisq/compactor) [![Go Report Card](https://goreportcard.com/badge/github.com/joseluisq/compactor)](https://goreportcard.com/report/github.com/joseluisq/compactor) [![PkgGoDev](https://pkg.go.dev/badge/github.com/joseluisq/compactor)](https://pkg.go.dev/github.com/joseluisq/compactor)

> [Tar](https://golang.org/pkg/archive/tar/)/[Gzip](https://golang.org/pkg/compress/gzip/) and [Zip](https://golang.org/pkg/archive/zip/) archive utilities with optional [checksum](https://en.wikipedia.org/wiki/Checksum) computation.

__Status:__ WIP

## Usage

### Tar/Gzip

```go
package main

import (
	"github.com/joseluisq/compactor"
)

func main() {
	compactor.CreateTarballWithChecksum(
		// 1. archive input file or directory
		"./my-file-or-dir",
		// 2. archive output file
		"~/my-archive.tar.gz",
		// 3. checksum algorithm
		"sha256",
		// 4. checksum output file
		"~/my-archive.CHECKSUM.txt",
	)

	// output files:
	//	~/my-archive.tar.gz
	//	~/my-archive.sha256.tar.txt
}
```

### Zip


```go
package main

import (
	"github.com/joseluisq/compactor"
)

func main() {
	compactor.CreateZipballWithChecksum(
		// 1. archive input file or directory
		"./my-file-or-dir",
		// 2. archive output file
		"~/my-archive.zip",
		// 3. checksum algorithm
		"sha256",
		// 4. checksum output file
		"~/my-archive.CHECKSUM.zip.txt",
	)

	// output files:
	//	~/my-archive.zip
	//	~/my-archive.sha256.zip.txt
}
```

## Contributions

Unless you explicitly state otherwise, any contribution intentionally submitted for inclusion in current work by you, as defined in the Apache-2.0 license, shall be dual licensed as described below, without any additional terms or conditions.

Feel free to send some [Pull request](https://github.com/joseluisq/compactor/pulls) or [issue](https://github.com/joseluisq/compactor/issues).

## License

This work is primarily distributed under the terms of both the [MIT license](LICENSE-MIT) and the [Apache License (Version 2.0)](LICENSE-APACHE).

© 2020-present [Jose Quintana](https://git.io/joseluisq)
