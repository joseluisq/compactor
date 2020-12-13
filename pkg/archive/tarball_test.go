package archive

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"
)

func TestCreateTarballBytes(t *testing.T) {
	type args struct {
		basePath string
		src      string
	}
	tests := []struct {
		name    string
		args    args
		outFile string
		wantErr bool
	}{
		{
			name: "invalid file path",
			args: args{
				src: "./fixtures/some.txt",
			},
			wantErr: true,
		},
		{
			name: "regular file archiving",
			args: args{
				src: "./fixtures/file.txt",
			},
			outFile: "./file.txt",
		},
		{
			name: "directory archiving",
			args: args{
				src: "./fixtures",
			},
			outFile: "./fixtures/file.txt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			outBuf := &bytes.Buffer{}
			err := CreateTarballBytes(tt.args.basePath, tt.args.src, outBuf)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTarballBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}

			// Create a temp directory for current test
			tmpDirPath, err := ioutil.TempDir("/tmp", "compactor-")
			if err != nil {
				t.Errorf("%v", err)
				return
			}

			// Content input file
			fTxt, err := os.Open("./fixtures/file.txt")
			defer fTxt.Close()
			if err != nil {
				t.Errorf("%v", err)
				return
			}
			inFile, err := fTxt.Stat()
			if err != nil {
				t.Errorf("%v", err)
				return
			}

			// Write tar/gz file content
			fTar, err := os.Create(tmpDirPath + "/file.tar.gz")
			defer fTar.Close()
			if err != nil {
				t.Errorf("%v", err)
				return
			}
			fTar.Write(outBuf.Bytes())
			_, err = fTar.Stat()
			if err != nil {
				t.Errorf("%v", err)
				return
			}

			// Extract tar/gz file
			var out bytes.Buffer
			cmd := exec.Command("tar", "-xvf", tmpDirPath+"/file.tar.gz", "-C", tmpDirPath)
			cmd.Stderr = os.Stderr
			cmd.Stdin = os.Stdin
			cmd.Stdout = &out
			err = cmd.Run()
			if err != nil {
				t.Errorf("%v", err)
				return
			}

			// Content extracted file
			fTxt2, err := os.Open(tmpDirPath + "/" + tt.outFile)
			defer fTxt.Close()
			if err != nil {
				t.Errorf("%v", err)
				return
			}
			outFile, err := fTxt2.Stat()
			if err != nil {
				t.Errorf("%v", err)
				return
			}

			if outFile.Size() <= 0 || outFile.Size() != inFile.Size() {
				t.Errorf(
					"CreateTarballBytes() = %v: uncompressed size %#x, want %#x",
					outFile.Name(),
					outFile.Size(),
					inFile.Size(),
				)
			}
		})
	}
}
