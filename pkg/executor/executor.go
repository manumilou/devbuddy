package executor

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"

	. "github.com/logrusorgru/aurora"
)

const defaultFailedCode = 1

func makeError(cmd *exec.Cmd, err error) error {
	var exitCode int
	// if err == nil {
	// 	ws := cmd.ProcessState.Sys().(syscall.WaitStatus)
	// 	exitCode = ws.ExitStatus()
	// } else {
	// 	if exitError, ok := err.(*exec.ExitError); ok {
	// 		ws := exitError.Sys().(syscall.WaitStatus)
	// 		exitCode = ws.ExitStatus()
	// 	}
	// }

	exitCode = cmd.ProcessState.Sys().(syscall.WaitStatus).ExitStatus()

	if exitCode != 0 {
		err = fmt.Errorf("command failed with code %i", exitCode)
	}
	return err
}

func Run(program string, args ...string) error {
	fmt.Println("Running: ", Bold(Cyan(program)), Cyan(strings.Join(args, " ")))

	cmd := exec.Command(program, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	err = makeError(cmd, err)
	return err
}

func RunShell(cmdline string) error {
	return Run("sh", "-c", cmdline)
}