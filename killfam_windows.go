// +build windows

package killfam

import (
	"os/exec"
	"strconv"
)

// Augment configures the process of the command such that its entire process
// tree is killable.
//
// Augment must be run after the command is created with exec.Command, but
// before the command is executed.
func Augment(cmd *exec.Cmd) {
	// Does nothing on Windows
}

// KillTree kills the entire process tree of the input process, assuming it was
// augmented using killfam.Augment.
func KillTree(cmd *exec.Cmd) error {
	// Adapted from:
	// https://stackoverflow.com/a/44551450/1376127
	kill := exec.Command("TASKKILL", "/T", "/F", "/PID", strconv.Itoa(cmd.Process.Pid))
	return kill.Run()
}
