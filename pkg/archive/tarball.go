package archive

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// CreateTarballBytes archives a file or directory src using Tar and Gzip compression.
func CreateTarballBytes(src string, outBuf io.Writer) error {
	zw := gzip.NewWriter(outBuf)
	tw := tar.NewWriter(zw)
	fi, err := os.Stat(src)
	if err != nil {
		return err
	}
	fm := fi.Mode()
	switch {
	case fm.IsRegular():
		// Get Tar source file header
		h, err := tar.FileInfoHeader(fi, src)
		if err != nil {
			return err
		}
		// Write Tar source file header
		if err := tw.WriteHeader(h); err != nil {
			return err
		}
		// Get source file content
		f, err := os.Open(src)
		defer f.Close()
		if err != nil {
			return err
		}
		if _, err := io.Copy(tw, f); err != nil {
			return err
		}
		break
	case fi.IsDir():
		// Traversing the directory tree on file system
		filepath.Walk(src, func(file string, fi os.FileInfo, err error) error {
			// Create a Tar file header
			h, err := tar.FileInfoHeader(fi, file)
			if err != nil {
				return err
			}
			// Since os.FileInfo's Name method only returns the base name of
			// the file it describes, it may be necessary to modify Header.Name
			// to provide the full path name of the file.
			// https://golang.org/src/archive/tar/common.go?#L626
			h.Name = filepath.ToSlash(file)
			// Write Tar header
			if err := tw.WriteHeader(h); err != nil {
				return err
			}
			// If it's not a directory, write file content instead
			if !fi.IsDir() {
				f, err := os.Open(file)
				defer f.Close()
				if err != nil {
					return err
				}
				if _, err := io.Copy(tw, f); err != nil {
					return err
				}
			}
			return nil
		})
		break
	default:
		return fmt.Errorf("archive/tar: unknown file mode %v", fm)
	}
	// Write Tar content
	if err := tw.Close(); err != nil {
		return err
	}
	// Write Gzip content
	if err := zw.Close(); err != nil {
		return err
	}
	return nil
}
