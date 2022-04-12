package tsgen

import (
	"bytes"
	"fmt"
	"os/exec"
)

func Prettier(b []byte) ([]byte, error) {
	var out bytes.Buffer
	var derr bytes.Buffer
	cmd := exec.Command("prettier", "--parser", "typescript", "--single-quote", "--tab-width", "4", "--print-width", "2000")
	cmd.Stdin = bytes.NewBuffer(b)
	cmd.Stdout = &out
	cmd.Stderr = &derr
	err := cmd.Run()
	if err != nil {
		fmt.Println(derr.String())
		return nil, err
	}
	return out.Bytes(), nil
}
