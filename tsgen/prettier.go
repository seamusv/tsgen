package tsgen

import (
	"bytes"
	"os/exec"
)

func Prettier(b []byte) ([]byte, error) {
	var out bytes.Buffer
	cmd := exec.Command("prettier", "--parser", "typescript", "--single-quote", "--tab-width", "4")
	cmd.Stdin = bytes.NewBuffer(b)
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}
