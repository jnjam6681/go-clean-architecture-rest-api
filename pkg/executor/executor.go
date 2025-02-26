package executor

import (
	"bytes"
	"os/exec"
	"strings"
)

type Executor interface {
	RunCommand(useSudo bool, cmd string, args ...string) (bytes.Buffer, error)
}

type executor struct{}

func NewExecutor() Executor {
	return &executor{}
}

func (r *executor) RunCommand(useSudo bool, cmd string, args ...string) (bytes.Buffer, error) {
	if useSudo {
		args = append([]string{cmd}, args...)
		cmd = "sudo"
	}
	execCmd := exec.Command(cmd, args...)

	var stdout, stderr bytes.Buffer
	execCmd.Stdout = &stdout
	execCmd.Stderr = &stderr

	execute := strings.Join([]string{execCmd.Path, strings.Join(execCmd.Args, " ")}, " ")

	if err := execCmd.Run(); err != nil {
		return stdout, &Error{
			Err:     err,
			Execute: execute,
			Stderr:  stderr.String(),
		}
	}
	return stdout, nil
}
