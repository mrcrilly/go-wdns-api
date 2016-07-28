package gowdns

import (
	"bufio"
	"fmt"
	"os/exec"
)

func executeCommand(ec *ExecutableConfig) error {
	cmd := exec.Command(ec.Executable, ec.Flags...)
	outputBuffer, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	err = cmd.Run()
	if err != nil {
		return err
	}

	if ec.Stdout != nil {
		scanner := bufio.NewScanner(outputBuffer)
		for scanner.Scan() {
			fmt.Fprintf(ec.Stdout, "%s\n", scanner.Text())
		}
	}

	return nil
}
