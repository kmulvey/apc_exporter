package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func getPowerStats(cmdPath string) (string, error) {

	var cmd = exec.Command(cmdPath)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	var err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("error running command, stderr: %s, go err: %w", stderr.String(), err)
	}

	return out.String(), nil
}
