// Package Compactor provides Tar/Gzip and Zip archive utilities with optional checksum computation.

package compactor

import (
	"testing"
)

func Test_createArchiveFile(t *testing.T) {
	type args struct {
		src    string
		dst    string
		format ArchiveFormat
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := createArchiveFile(tt.args.src, tt.args.dst, tt.args.format); (err != nil) != tt.wantErr {
				t.Errorf("createArchiveFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateTarball(t *testing.T) {
	type args struct {
		src string
		dst string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateTarball(tt.args.src, tt.args.dst); (err != nil) != tt.wantErr {
				t.Errorf("CreateTarball() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateZipball(t *testing.T) {
	type args struct {
		src string
		dst string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateZipball(tt.args.src, tt.args.dst); (err != nil) != tt.wantErr {
				t.Errorf("CreateZipball() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateTarballWithChecksum(t *testing.T) {
	type args struct {
		src          string
		dst          string
		checksumAlgo string
		checksumDst  string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateTarballWithChecksum(tt.args.src, tt.args.dst, tt.args.checksumAlgo, tt.args.checksumDst)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTarballWithChecksum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateTarballWithChecksum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateZipballWithChecksum(t *testing.T) {
	type args struct {
		src          string
		dst          string
		checksumAlgo string
		checksumDst  string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateZipballWithChecksum(tt.args.src, tt.args.dst, tt.args.checksumAlgo, tt.args.checksumDst)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateZipballWithChecksum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateZipballWithChecksum() = %v, want %v", got, tt.want)
			}
		})
	}
}
