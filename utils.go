package main

import (
	"os/exec"
)

const (
	PACMAN string = "/usr/bin/pacman"
)

// Run a shell command and return result as any error raised
func ShellExec(command string, args ...string) (string, error) {
	output, err := exec.Command(command, args...).Output()
	return string(output), err

}
