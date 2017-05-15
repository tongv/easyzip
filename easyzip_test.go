package easyzip

import (
	"testing"
)

func Test_UnCompress_1(t *testing.T) {
	UnCompress("/path/to/zip/demo.zip","/tmp/")
}
