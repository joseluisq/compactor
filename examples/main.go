package main

import "github.com/joseluisq/compactor"

func main() {
	compactor.CreateTarballWithChecksum(
		"./pkg/archive/",
		"fixtures",
		"./.tmp/pkg.tar.gz",
		"sha256",
		"./.tmp/pkg.CHECKSUM.tar.txt",
	)

	compactor.CreateZipballWithChecksum(
		"./pkg/archive/",
		"fixtures/file.txt",
		"./.tmp/pkg.zip",
		"sha256",
		"./.tmp/pkg.CHECKSUM.zip.txt",
	)
}
