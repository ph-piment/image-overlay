package image

import (
	"reflect"
	"strings"
	"testing"
)

func Test_GetPngImage(t *testing.T) {
	tests := []struct {
		name     string
		fileName string
		want     string
	}{
		{
			name:     "png画像取得確認",
			fileName: "../source/background.png",
			want:     "image.", // image.Imageのはずなんだけど*image.NRGBAが返ってくる。そのためimage.までで比較する。
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPngImage(tt.fileName)
			if err != nil {
				t.Errorf("GetPngImage() Error %v, %v", err, tt.fileName)
			}
			if !strings.Contains(reflect.TypeOf(got).String(), tt.want) {
				t.Errorf("GetPngImage() = %v, want %v", reflect.TypeOf(got).String(), tt.want)
			}
		})
	}
}
