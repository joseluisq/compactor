package archive

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// CreateZipballBytes archives a file or directory (src path) using Zip.
// basePath param specify the base path directory of src path which will be skipped for each archive file header.
// Otherwise if basePath param is empty then only src path will taken into account.
func CreateZipballBytes(basePath string, src string, outBuf io.Writer) error {
	zw := zip.NewWriter(outBuf)
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
		// Get Zip source file header
		h, err := zip.FileInfoHeader(fi)
		if err != nil {
			return err
		}
		// Write Zip source file header
		hw, err := zw.CreateHeader(h)
		if err != nil {
			return err
		}
		h.Method = zip.Deflate
		// Get source file content
		f, err := os.Open(src)
		defer f.Close()
		if err != nil {
			return err
		}
		if _, err := io.Copy(hw, f); err != nil {
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
			// Create a Zip file header
			h, err := zip.FileInfoHeader(fi)
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
			h.Name = filepath.ToSlash(fileName)
			if fi.IsDir() {
				h.Name += "/"
			} else {
				h.Method = zip.Deflate
			}
			// Write Zip header
			hw, err := zw.CreateHeader(h)
			if err != nil {
				return err
			}
			// If it's not a directory, write file content instead
			if !fi.IsDir() {
				f, err := os.Open(file)
				defer f.Close()
				if err != nil {
					return err
				}
				if _, err := io.Copy(hw, f); err != nil {
					return err
				}
			}
			return nil
		})
	default:
		return fmt.Errorf("archive/zip: unknown file mode %v", fm)
	}
	// Write Zip content
	if err := zw.Close(); err != nil {
		return err
	}
	return nil
}
