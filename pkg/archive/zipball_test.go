package archive

import (
	"bytes"
	"testing"
)

func TestCreateZipballBytes(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name       string
		args       args
		wantOutBuf string
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			outBuf := &bytes.Buffer{}
			if err := CreateZipballBytes(tt.args.src, outBuf); (err != nil) != tt.wantErr {
				t.Errorf("CreateZipballBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOutBuf := outBuf.String(); gotOutBuf != tt.wantOutBuf {
				t.Errorf("CreateZipballBytes() = %v, want %v", gotOutBuf, tt.wantOutBuf)
			}
		})
	}
}
