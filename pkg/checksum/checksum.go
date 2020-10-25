package checksum

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// ComputeChecksum computes a `MD5`, `SHA1`, `SHA256` or `SHA512` message digest.
func ComputeChecksum(r io.Reader, algo string) (hash string, err error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	algo = strings.ToLower(strings.TrimSpace(algo))
	switch algo {
	case "md5":
		hash = fmt.Sprintf("%x", md5.Sum(data))
	case "sha1":
		hash = fmt.Sprintf("%x", sha1.Sum(data))
	case "sha256":
		hash = fmt.Sprintf("%x", sha256.Sum256(data))
	case "sha512":
		hash = fmt.Sprintf("%x", sha512.Sum512(data))
	default:
		return "", fmt.Errorf("hash algorithm `%s` is not supported", algo)
	}
	return hash, nil
}

// CreateChecksumFiles computes MD5, SHA1, SHA256 or SHA512 message digest and save it into a file.
// It returns checksum file paths or an error.
func CreateChecksumFiles(files []string, checksumAlgos []string, checksumDst string, filesBasename bool) ([]string, error) {
	checksums := make(map[string][]string)
	for _, algo := range checksumAlgos {
		for _, f := range files {
			r, err := os.Open(f)
			if err != nil {
				return nil, fmt.Errorf("failed to read %s file: %w", f, err)
			}
			sum, err := ComputeChecksum(r, algo)
			if err != nil {
				return nil, err
			}
			checksums[algo] = append(checksums[algo], sum, f)
		}
	}

	var outfiles []string
	for algo, values := range checksums {
		filename := strings.Replace(checksumDst, "CHECKSUM", algo, -1)
		f, err := os.Create(filename)
		if err != nil {
			return nil, err
		}
		for i := 0; i < len(values); i += 2 {
			hash := values[i]
			file := values[i+1]
			if filesBasename {
				file = filepath.Base(file)
			}
			if _, err := f.WriteString(fmt.Sprintf("%s  %s\n", hash, file)); err != nil {
				return nil, err
			}
		}
		outfiles = append(outfiles, filename)
	}
	return outfiles, nil
}
