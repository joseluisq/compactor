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

func createArchiveFile(basePath string, src string, dst string, format ArchiveFormat) error {
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
		if err := archive.CreateZipballBytes(basePath, src, &buf); err != nil {
			return err
		}
	}
	if format == ArchiveFormatTar {
		if err := archive.CreateTarballBytes(basePath, src, &buf); err != nil {
			return err
		}
	}
	return ioutil.WriteFile(dst, buf.Bytes(), 0755)
}

// CreateTarball archives and compresses a file or folder (src) using Tar/Gzip to dst (tarball).
// basePath param specify the base path directory of src path which will be skipped for each archive file header.
// Otherwise if basePath param is empty then only src path will taken into account.
func CreateTarball(basePath string, src string, dst string) error {
	return createArchiveFile(basePath, src, dst, ArchiveFormatTar)
}

// CreateZipball archives and compresses a file or folder (src) using Zip to dst (zipball).
// basePath param specify the base path directory of src path which will be skipped for each archive file header.
// Otherwise if basePath param is empty then only src path will taken into account.
func CreateZipball(basePath string, src string, dst string) error {
	return createArchiveFile(basePath, src, dst, ArchiveFormatZip)
}

// CreateTarballWithChecksum archives and compresses a file or folder (src) using Tar/Gzip to dst (tarball) with checksum (`md5`, `sha1`, `sha256` or `sha512`). It returns the checksum file path or an error.
// basePath param specify the base path directory of src path which will be skipped for each archive file header.
// Otherwise if basePath param is empty then only src path will taken into account.
func CreateTarballWithChecksum(basePath string, src string, dst string, checksumAlgo string, checksumDst string) (string, error) {
	if err := CreateTarball(basePath, src, dst); err != nil {
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
// basePath param specify the base path directory of src path which will be skipped for each archive file header.
// Otherwise if basePath param is empty then only src path will taken into account.
func CreateZipballWithChecksum(basePath, src string, dst string, checksumAlgo string, checksumDst string) (string, error) {
	if err := CreateZipball(basePath, src, dst); err != nil {
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
