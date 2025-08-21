package service

import (
	"os"
	"path/filepath"

	"changeme/backend/domain"
	"changeme/backend/infrastructure"
	"changeme/backend/usecase"

	"github.com/wailsapp/wails/v3/pkg/application"
)

// AppService provides application management functionality
type AppService struct {
	appUseCase *usecase.AppUseCase
	app        *application.App
}

// NewAppService creates a new instance of AppService with dependencies
func NewAppService() *AppService {
	homeDir, _ := os.UserHomeDir()
	configPath := filepath.Join(homeDir, ".config", "launchy", "favorites.json")

	// Initialize infrastructure
	appRepo := infrastructure.NewFileAppRepository(configPath)
	appLauncher := infrastructure.NewSystemAppLauncher()

	// Initialize use case
	appUseCase := usecase.NewAppUseCase(appRepo, appLauncher)

	appService := &AppService{
		appUseCase: appUseCase,
	}

	return appService
}

// SetApp sets the application reference for window management
func (s *AppService) SetApp(app *application.App) {
	s.app = app
}

// GetApps returns all available applications
func (s *AppService) GetApps() ([]domain.AppInfo, error) {
	apps, err := s.appUseCase.GetAllApps()
	if err != nil {
		return nil, err
	}
	return apps, nil
}

// GetFavorites returns all favorited applications
func (s *AppService) GetFavorites() ([]domain.AppInfo, error) {
	favorites, err := s.appUseCase.GetFavorites()
	if err != nil {
		return nil, err
	}
	return favorites, nil
}

// SearchApps searches for applications based on the query
func (s *AppService) SearchApps(query string) ([]domain.AppInfo, error) {
	results, err := s.appUseCase.SearchApps(query)
	if err != nil {
		return nil, err
	}
	return results, nil
}

// LaunchApp launches an application by its exec path and closes Launchy
func (s *AppService) LaunchApp(execPath string) error {
	// Launch the application
	err := s.appUseCase.LaunchApp(execPath)
	if err != nil {
		return err
	}

	// Close Launchy after successfully launching the app
	if s.app != nil {
		s.app.Quit()
	}

	return nil
}

// ToggleFavorite toggles the favorite status of an application
func (s *AppService) ToggleFavorite(appName string) error {
	return s.appUseCase.ToggleFavorite(appName)
}
