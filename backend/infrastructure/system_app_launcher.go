package infrastructure

import (
	"fmt"
	"os/exec"
	"strings"
)

// SystemAppLauncher implements AppLauncher using system commands
type SystemAppLauncher struct{}

// NewSystemAppLauncher creates a new SystemAppLauncher
func NewSystemAppLauncher() *SystemAppLauncher {
	return &SystemAppLauncher{}
}

// Launch executes an application
func (l *SystemAppLauncher) Launch(execPath string) error {
	parts := strings.Fields(execPath)
	if len(parts) == 0 {
		return fmt.Errorf("invalid exec command")
	}

	cmd := exec.Command(parts[0], parts[1:]...)
	return cmd.Start()
}
