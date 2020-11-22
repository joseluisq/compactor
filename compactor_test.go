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
		{
			name: "invalid archive format",
			args: args{
				src:    "pkg/archive/fixtures/file.txt",
				dst:    "/tmp/file.xz",
				format: 3,
			},
			wantErr: true,
		},
		{
			name: "archive file with default destination",
			args: args{
				src:    "pkg/archive/fixtures/file.txt",
				format: ArchiveFormatZip,
			},
		},
		{
			name: "archive file without destination extension",
			args: args{
				src:    "pkg/archive/fixtures/file.txt",
				dst:    "/tmp/file.txt",
				format: ArchiveFormatZip,
			},
		},
		{
			name: "create tar/gz file",
			args: args{
				src:    "pkg/archive/fixtures/file.txt",
				dst:    "/tmp/file.tar.gz",
				format: ArchiveFormatTar,
			},
		},
		{
			name: "create zip file",
			args: args{
				src:    "pkg/archive/fixtures/file.txt",
				dst:    "/tmp/file.zip",
				format: ArchiveFormatZip,
			},
		},
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
		{
			name: "create valid tar/gz file",
			args: args{
				src: "pkg/archive/fixtures/file.txt",
				dst: "/tmp/file.tar.gz",
			},
		},
		{
			name: "invalid tar/gz file",
			args: args{
				src: "pkg/archive/fixtures/file.abc",
				dst: "/tmp/file.tar.gz",
			},
			wantErr: true,
		},
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
		{
			name: "create valid zip file",
			args: args{
				src: "pkg/archive/fixtures/file.txt",
				dst: "/tmp/file.zip",
			},
		},
		{
			name: "invalid zip file",
			args: args{
				src: "pkg/archive/fixtures/file.abc",
				dst: "/tmp/file.zip",
			},
			wantErr: true,
		},
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
		{
			name: "create valid tar/gz file",
			args: args{
				src:          "pkg/archive/fixtures/file.txt",
				dst:          "/tmp/file.tar.gz",
				checksumAlgo: "sha1",
				checksumDst:  "/tmp/file.CHECKSUM.tar.txt",
			},
			want: "/tmp/file.sha1.tar.txt",
		},
		{
			name: "invalid tar/gz file",
			args: args{
				src:          "pkg/archive/fixtures/file.xyz",
				dst:          "/tmp/file.tar.gz",
				checksumAlgo: "sha1",
				checksumDst:  "/tmp/file.CHECKSUM.tar.txt",
			},
			wantErr: true,
		},
		{
			name: "tar/gz file with invalid algorithm",
			args: args{
				src:          "pkg/archive/fixtures/file.txt",
				dst:          "/tmp/file.tar.gz",
				checksumAlgo: "sha11",
				checksumDst:  "/tmp/file.CHECKSUM.tar.txt",
			},
			wantErr: true,
		},
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
		{
			name: "create valid zip file",
			args: args{
				src:          "pkg/archive/fixtures/file.txt",
				dst:          "/tmp/file.zip",
				checksumAlgo: "sha1",
				checksumDst:  "/tmp/file.CHECKSUM.zip.txt",
			},
			want: "/tmp/file.sha1.zip.txt",
		},
		{
			name: "invalid zip file",
			args: args{
				src:          "pkg/archive/fixtures/file.xyz",
				dst:          "/tmp/file.zip",
				checksumAlgo: "sha1",
				checksumDst:  "/tmp/file.CHECKSUM.zip.txt",
			},
			wantErr: true,
		},
		{
			name: "zip file with invalid algorithm",
			args: args{
				src:          "pkg/archive/fixtures/file.txt",
				dst:          "/tmp/file.zip",
				checksumAlgo: "sha11",
				checksumDst:  "/tmp/file.CHECKSUM.zip.txt",
			},
			wantErr: true,
		},
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
