package go_ffprobe_test

import (
	"testing"

	ffprobe "github.com/smith-30/go-ffprobe"
)

func TestGetFileInfo(t *testing.T) {
	exp := "test"
	act := ffprobe.GetFileInfo(exp)

	if exp != act {
		t.Errorf("failed.")
	}
}
