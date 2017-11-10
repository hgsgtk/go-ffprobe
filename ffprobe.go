package go_ffprobe

import (
	"encoding/json"
	"os/exec"
)

var (
	ExecFunc = ExecCmd
)

func Execute(fileName string) (r Result, err error) {
	out, err := ExecFunc(fileName)

	if err != nil {
		return r, err
	}

	if err := json.Unmarshal(out, &r); err != nil {
		return r, err
	}

	return r, nil
}

func ExecCmd(fileName string) ([]byte, error) {
	return exec.Command("ffprobe",
		"-v", "quiet", "-print_format", "json", "-show_format", "-show_streams", fileName).Output()
}
