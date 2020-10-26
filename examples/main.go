package main

import "github.com/joseluisq/compactor"

func main() {
	compactor.CreateTarballWithChecksum(
		"./pkg",
		"./.tmp/pkg.tar.gz",
		"sha256",
		"./.tmp/pkg.CHECKSUM.tar.txt",
	)

	compactor.CreateZipballWithChecksum(
		"./pkg",
		"./.tmp/pkg.zip",
		"sha256",
		"./.tmp/pkg.CHECKSUM.zip.txt",
	)
}
