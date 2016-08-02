package gowdns

import (
	"io/ioutil"
	"os/exec"
)

func executeCommand(ec *ExecutableConfig) (string, string, error) {
	cmd := exec.Command(ec.Executable, ec.Flags...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "<stdout: error>", "<stderr: nil>", err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return "<stdout: nil>", "<stderr: error>", err
	}

	err = cmd.Start()
	if err != nil {
		return "<stdout: nil>", "<stderr: nil>", err
	}

	stdoutToString, err := ioutil.ReadAll(stdout)
	if err != nil {
		return "<stdout: " + err.Error() + ">", "<stderr: nil>", err
	}

	stderrToString, err := ioutil.ReadAll(stderr)
	if err != nil {
		return string(stdoutToString), "<stderr: " + err.Error() + ">", err
	}

	err = cmd.Wait()
	return string(stdoutToString), string(stderrToString), err
}
