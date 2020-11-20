package checksum

import (
	"io"
	"reflect"
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
