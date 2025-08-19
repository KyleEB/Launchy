package repository

// AppLauncher defines the interface for launching applications
type AppLauncher interface {
	Launch(execPath string) error
}
