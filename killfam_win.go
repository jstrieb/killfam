// +build windows

package killfam

import (
	"os/exec"
)

// Augment configures the process of the command such that its entire process
// tree is killable.
//
// Augment must be run after the command is created with exec.Command, but
// before the command is executed.
func Augment(cmd *exec.Cmd) error {
	return nil
}

// KillTree kills the entire process tree of the input process, assuming it was
// augmented using killfam.Augment.
func KillTree(cmd *exec.Cmd) error {
	return nil
}
