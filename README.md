# Compactor [![Build Status](https://travis-ci.com/joseluisq/compactor.svg?branch=master)](https://travis-ci.com/joseluisq/compactor)

> [Tar](https://golang.org/pkg/archive/tar/)/[Gzip](https://golang.org/pkg/compress/gzip/) and [Zip](https://golang.org/pkg/archive/zip/) archive utilities with optional [checksum](https://en.wikipedia.org/wiki/Checksum) computation.

**Status:** WIP

## Usage

### Tar/Gzip

```go
package main

import (
	"github.com/joseluisq/compactor"
)

func main() {
	compactor.CreateTarballWithChecksum(
		// 1. archive input directory or file
		"./my-dir-or-file",
		// 2. archive output file
		"~/my-archive.tar.gz",
		// 3. checksum algorithm
		"sha256",
		// 4. checksum output file
		"~/my-archive.CHECKSUM.txt",
	)

    // output files:
    //	~/my-archive.tar.gz
    //	~/my-archive.sha256.txt
}
```

### Zip

TODO

## Contributions

Unless you explicitly state otherwise, any contribution intentionally submitted for inclusion in current work by you, as defined in the Apache-2.0 license, shall be dual licensed as described below, without any additional terms or conditions.

Feel free to send some [Pull request](https://github.com/joseluisq/compactor/pulls) or [issue](https://github.com/joseluisq/compactor/issues).

## License

This work is primarily distributed under the terms of both the [MIT license](LICENSE-MIT) and the [Apache License (Version 2.0)](LICENSE-APACHE).

© 2020-present [Jose Quintana](https://git.io/joseluisq)
