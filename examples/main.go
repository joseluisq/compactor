package main

import (
	"github.com/joseluisq/compactor"
)

func main() {
	compactor.CreateTarballWithChecksum(
		"./pkg",
		"./.tmp/pkg.tar.gz",
		"sha256",
		"./.tmp/pkg.CHECKSUM.txt",
	)
}
