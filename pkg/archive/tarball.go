// Package archive provides archiving and files compressing using Tar-GZ or Zip formmat.
package archive

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// CreateTarballBytes archives a file or directory src using Tar and Gzip compression.
// basePath param specify the base path directory of src path which will be skipped for each archive file header.
// Otherwise if basePath param is empty then only src path will taken into account.
func CreateTarballBytes(basePath string, src string, outBuf io.Writer) error {
	zw := gzip.NewWriter(outBuf)
	tw := tar.NewWriter(zw)
	src = strings.TrimSpace(src)
	basePath = strings.TrimSpace(basePath)
	if basePath != "" {
		p, err := filepath.Abs(filepath.Join(basePath, src))
		if err != nil {
			return err
		}
		src = p
	}
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
	case fi.IsDir():
		basePathAbs := ""
		if basePath != "" {
			basePathAbs, err = filepath.Abs(basePath)
			if err != nil {
				return err
			}
		}
		// Traversing the directory tree on a file system
		filepath.Walk(src, func(file string, fi os.FileInfo, err error) error {
			// Create a Tar file header
			h, err := tar.FileInfoHeader(fi, file)
			if err != nil {
				return err
			}
			// Base path file support
			fileName := file
			if basePathAbs != "" {
				p := strings.TrimSpace(strings.ReplaceAll(file, basePathAbs, ""))
				if p == "" {
					return nil
				}
				if strings.HasPrefix(p, "/") {
					p = p[1:]
				}
				fileName = p
			}
			// Since os.FileInfo's Name method only returns the base name of
			// the file it describes, it may be necessary to modify Header.Name
			// to provide the full path name of the file.
			// https://golang.org/src/archive/tar/common.go?#L626
			h.Name = filepath.ToSlash(fileName)
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
