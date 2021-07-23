// +build !windows

package killfam

import (
	"os/exec"
	"syscall"
)

// Augment configures the process of the command such that its entire process
// tree is killable.
//
// Augment must be run after the command is created with exec.Command, but
// before the command is executed.
func Augment(cmd *exec.Cmd) {
	// Adapted from:
	// https://stackoverflow.com/a/29552044/1376127
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
}

// KillTree kills the entire process tree of the input process, assuming it was
// augmented using killfam.Augment.
func KillTree(cmd *exec.Cmd) error {
	// Adapted from:
	// https://stackoverflow.com/a/29552044/1376127
	pgid, err := syscall.Getpgid(cmd.Process.Pid)
	if err != nil {
		return err
	}

	syscall.Kill(-pgid, syscall.SIGTERM)
	return nil
}
