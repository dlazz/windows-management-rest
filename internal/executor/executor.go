package executor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
)

func RunPowershellCommand(command string) (output []byte, err error) {
	StdOut := bytes.NewBufferString("")
	StdErr := bytes.NewBufferString("")
	cmd := exec.Command("powershell.exe", "-c", command)
	cmd.Stdout = StdOut
	cmd.Stderr = StdErr

	if err := cmd.Run(); err != nil {
		return output, err
	}

	if StdErr.String() != "" {
		return output, fmt.Errorf(StdErr.String())
	}

	return StdOut.Bytes(), err
}

func Wrap(command string, res interface{}) error {
	out, err := RunPowershellCommand(command)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(out, &res); err != nil {
		return err
	}
	return nil
}
