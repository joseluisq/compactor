package archive

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// CreateZipballBytes archives a file or directory (src path) using Zip.
func CreateZipballBytes(src string, outBuf io.Writer) error {
	zw := zip.NewWriter(outBuf)
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
		h.Name = filepath.ToSlash(src)
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
		break
	case fi.IsDir():
		// Traversing the directory tree on a file system
		filepath.Walk(src, func(file string, fi os.FileInfo, err error) error {
			// Create a Zip file header
			h, err := zip.FileInfoHeader(fi)
			if err != nil {
				return err
			}
			h.Name = filepath.ToSlash(file)
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
		break
	default:
		return fmt.Errorf("archive/zip: unknown file mode %v", fm)
	}
	// Write Zip content
	if err := zw.Close(); err != nil {
		return err
	}
	return nil
}