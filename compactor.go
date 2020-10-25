package compactor

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/joseluisq/compactor/pkg/archive"
	"github.com/joseluisq/compactor/pkg/checksum"
)

const (
	tarGzExt = "tar.gz"
)

// CreateTarball archives and compresses a file or folder (src) using Tar/Gzip to dst (tarball).
func CreateTarball(src string, dst string) error {
	dst = strings.TrimSpace(dst)
	if dst == "" {
		_, src := filepath.Split(src)
		dst = src + "." + tarGzExt
	}
	if !strings.HasSuffix(dst, "."+tarGzExt) {
		dst = dst + "." + tarGzExt
	}
	err := os.MkdirAll(filepath.Dir(dst), 0755)
	if err != nil {
		return fmt.Errorf("can't make provided parent directories: %s", err)
	}
	var buf bytes.Buffer
	if err := archive.CreateTarballBytes(src, &buf); err != nil {
		return err
	}
	if err := ioutil.WriteFile(dst, buf.Bytes(), 0755); err != nil {
		return err
	}
	return nil
}

// CreateTarballWithChecksum archives and compresses a file or folder (src) using Tar/Gzip to dst (tarball) with checksum. It returns the checksum file path or an error.
func CreateTarballWithChecksum(src string, dst string, checksumAlgo string, checksumDst string) (string, error) {
	if err := CreateTarball(src, dst); err != nil {
		return "", err
	}
	files, err := checksum.CreateChecksumFiles(
		[]string{dst},
		[]string{checksumAlgo},
		checksumDst,
		true,
	)
	if err != nil {
		return "", fmt.Errorf("can't create checksum(s): %s", err)
	}
	return files[0], nil
}
