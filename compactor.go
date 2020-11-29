// Package compactor provides Tar/Gzip and Zip archive utilities with optional checksum computation.
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

// ArchiveFormat represents the archive output format.
type ArchiveFormat uint8

const (
	// ArchiveFormatTar represents the Tar/Gzip output format.
	ArchiveFormatTar ArchiveFormat = iota
	// ArchiveFormatZip represents the Zip output format.
	ArchiveFormatZip
)

func createArchiveFile(src string, dst string, format ArchiveFormat) error {
	var ext string
	switch format {
	case ArchiveFormatTar:
		ext = "tar.gz"
		break
	case ArchiveFormatZip:
		ext = "zip"
		break
	default:
		return fmt.Errorf("archive format provided is not supported")
	}

	dst = strings.TrimSpace(dst)
	if dst == "" {
		_, src := filepath.Split(src)
		dst = src + "." + ext
	}
	if !strings.HasSuffix(dst, "."+ext) {
		dst = dst + "." + ext
	}

	err := os.MkdirAll(filepath.Dir(dst), 0755)
	if err != nil {
		return fmt.Errorf("can't create provided parent directories: %s", err)
	}

	var buf bytes.Buffer
	if format == ArchiveFormatZip {
		if err := archive.CreateZipballBytes(src, &buf); err != nil {
			return err
		}
	}
	if format == ArchiveFormatTar {
		if err := archive.CreateTarballBytes(src, &buf); err != nil {
			return err
		}
	}
	return ioutil.WriteFile(dst, buf.Bytes(), 0755)
}

// CreateTarball archives and compresses a file or folder (src) using Tar/Gzip to dst (tarball).
func CreateTarball(src string, dst string) error {
	return createArchiveFile(src, dst, ArchiveFormatTar)
}

// CreateZipball archives and compresses a file or folder (src) using Zip to dst (zipball).
func CreateZipball(src string, dst string) error {
	return createArchiveFile(src, dst, ArchiveFormatZip)
}

// CreateTarballWithChecksum archives and compresses a file or folder (src) using Tar/Gzip to dst (tarball) with checksum (`md5`, `sha1`, `sha256` or `sha512`). It returns the checksum file path or an error.
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

// CreateZipballWithChecksum archives and compresses a file or folder (src) using Zip to dst (Zipball) with checksum (`md5`, `sha1`, `sha256` or `sha512`). It returns the checksum file path or an error.
func CreateZipballWithChecksum(src string, dst string, checksumAlgo string, checksumDst string) (string, error) {
	if err := CreateZipball(src, dst); err != nil {
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
