package checksum

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestComputeChecksum(t *testing.T) {
	type args struct {
		r    io.Reader
		algo string
	}
	tests := []struct {
		name     string
		args     args
		wantHash string
		wantErr  bool
	}{
		{
			name: "invalid algorithm",
			args: args{
				r:    strings.NewReader("abc"),
				algo: "md55",
			},
			wantErr: true,
		},
		{
			name: "md5sum",
			args: args{
				r:    strings.NewReader("abc"),
				algo: "md5",
			},
			wantHash: "900150983cd24fb0d6963f7d28e17f72",
		},
		{
			name: "sha1sum",
			args: args{
				r:    strings.NewReader("cde"),
				algo: "sha1",
			},
			wantHash: "5af13954a67eab2973b4ade01186602dd8739787",
		},
		{
			name: "sha256sum",
			args: args{
				r:    strings.NewReader("efg"),
				algo: "sha256",
			},
			wantHash: "d4ffe8e9ee0b48eba716706123a7187f32eae3bdcb0e7763e41e533267bd8a53",
		},
		{
			name: "sha512sum",
			args: args{
				r:    strings.NewReader("ghi"),
				algo: "sha512",
			},
			wantHash: "366aead3bed29b6d1de2b8d211e791e5dc7a9611b3d4c61c9323128d746e670a69e9690ce5620efc3b36f6d1b655ce36a72a2fbed4927448b668f1e3f341c0d9",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHash, err := ComputeChecksum(tt.args.r, tt.args.algo)
			if (err != nil) != tt.wantErr {
				t.Errorf("ComputeChecksum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotHash != tt.wantHash {
				t.Errorf("ComputeChecksum() = %v, want %v", gotHash, tt.wantHash)
			}
		})
	}
}

func TestCreateChecksumFiles(t *testing.T) {
	type args struct {
		files         []string
		checksumAlgos []string
		checksumDst   string
		filesBasename bool
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "invalid source files",
			args: args{
				files:         []string{"../../LICENSE-x", "../../LICENSE-y"},
				checksumAlgos: []string{"md5", "sha1"},
				checksumDst:   "/tmp/LICENSE.CHECKSUM.txt",
				filesBasename: true,
			},
			wantErr: true,
		},
		{
			name: "invalid checksum algorithms",
			args: args{
				files:         []string{"../../LICENSE-MIT", "../../LICENSE-APACHE"},
				checksumAlgos: []string{"md55", "sha11"},
				checksumDst:   "/tmp/LICENSE.CHECKSUM.txt",
				filesBasename: true,
			},
			wantErr: true,
		},
		{
			name: "checksums of valid source file",
			args: args{
				files:         []string{"../../LICENSE-MIT", "../../LICENSE-APACHE"},
				checksumAlgos: []string{"md5", "sha1"},
				checksumDst:   "/tmp/LICENSE.CHECKSUM.txt",
				filesBasename: true,
			},
			want: []string{"/tmp/LICENSE.md5.txt", "/tmp/LICENSE.sha1.txt"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateChecksumFiles(tt.args.files, tt.args.checksumAlgos, tt.args.checksumDst, tt.args.filesBasename)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateChecksumFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateChecksumFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}
