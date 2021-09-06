package file

import (
	"strings"
	"testing"
)

func Test_GetSourceFilePath(t *testing.T) {
	tests := []struct {
		name     string
		fileName string
		want     string
	}{
		{
			name:     "sourceファイルパス確認",
			fileName: "aaa.png",
			want:     "source/aaa.png",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetSourceFilePath(tt.fileName)
			if err != nil {
				t.Errorf("GetSourceFilePath() Error %v", err)
			}
			if !strings.HasSuffix(got, tt.want) {
				t.Errorf("GetSourceFilePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_GetOutputFilePath(t *testing.T) {
	tests := []struct {
		name     string
		fileName string
		want     string
	}{
		{
			name:     "outputファイルパス確認",
			fileName: "bbb.png",
			want:     "output/bbb.png",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetOutputFilePath(tt.fileName)
			if err != nil {
				t.Errorf("GetOutputFilePath() Error %v", err)
			}
			if !strings.HasSuffix(got, tt.want) {
				t.Errorf("GetSourceFilePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
