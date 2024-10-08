package utils

import (
	"os/exec"
	"strings"
	"syscall"
)

// stdout, stderr, exitCode := RunCommand("/bin/sh", "-c", "whoami")
func RunCommand(shell, interpretCmdAsString, command string) (string, string, int, error) {
	if shell == "" {
		shell = "/bin/sh"
	}
	if interpretCmdAsString == "" {
		interpretCmdAsString = "-c"
	}

	// Set up the command to run via the shell
	cmd := exec.Command(shell, interpretCmdAsString, command)

	// Capture stdout and stderr
	var stdout, stderr strings.Builder
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Run the command
	err := cmd.Run()

	// Get exit status
	exitStatus := 0
	if err != nil {
		// If an error occurred, check if it's an exit error and get exit status
		if exitErr, ok := err.(*exec.ExitError); ok {
			if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
				exitStatus = status.ExitStatus()
			}
		} else {
			return "", "", 0, err
		}
	} else {
		// If no error, get the exit status from ProcessState
		if status, ok := cmd.ProcessState.Sys().(syscall.WaitStatus); ok {
			exitStatus = status.ExitStatus()
		}
	}

	// Return stdout, stderr, exit status, and error
	return stdout.String(), stderr.String(), exitStatus, nil
}
