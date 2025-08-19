package service

import (
	"fmt"
	"os"
	"path/filepath"

	"changeme/backend/domain"
	"changeme/backend/infrastructure"
	"changeme/backend/usecase"
)

// AppService provides application management functionality
type AppService struct {
	appUseCase *usecase.AppUseCase
}

// NewAppService creates a new instance of AppService with dependencies
func NewAppService() *AppService {
	fmt.Println("Initializing AppService...")

	homeDir, _ := os.UserHomeDir()
	configPath := filepath.Join(homeDir, ".config", "launchy", "favorites.json")
	fmt.Printf("Config path: %s\n", configPath)

	// Initialize infrastructure
	appRepo := infrastructure.NewFileAppRepository(configPath)
	appLauncher := infrastructure.NewSystemAppLauncher()

	// Initialize use case
	appUseCase := usecase.NewAppUseCase(appRepo, appLauncher)

	// Note: App discovery will happen on first GetApps call
	fmt.Println("AppService ready - app discovery will happen on first request")

	appService := &AppService{
		appUseCase: appUseCase,
	}

	fmt.Println("AppService initialization completed")
	return appService
}

// GetApps returns all available applications
func (s *AppService) GetApps() ([]domain.AppInfo, error) {
	apps, err := s.appUseCase.GetAllApps()
	if err != nil {
		fmt.Printf("GetApps error: %v\n", err)
		return nil, err
	}
	fmt.Printf("GetApps returned %d applications\n", len(apps))
	return apps, nil
}

// GetFavorites returns all favorited applications
func (s *AppService) GetFavorites() ([]domain.AppInfo, error) {
	favorites, err := s.appUseCase.GetFavorites()
	if err != nil {
		fmt.Printf("GetFavorites error: %v\n", err)
		return nil, err
	}
	fmt.Printf("GetFavorites returned %d favorites\n", len(favorites))
	return favorites, nil
}

// SearchApps searches for applications based on the query
func (s *AppService) SearchApps(query string) ([]domain.AppInfo, error) {
	results, err := s.appUseCase.SearchApps(query)
	if err != nil {
		fmt.Printf("SearchApps error: %v\n", err)
		return nil, err
	}
	fmt.Printf("SearchApps returned %d results for query: %s\n", len(results), query)
	return results, nil
}

// LaunchApp launches an application by its exec path
func (s *AppService) LaunchApp(execPath string) error {
	return s.appUseCase.LaunchApp(execPath)
}

// ToggleFavorite toggles the favorite status of an application
func (s *AppService) ToggleFavorite(appName string) error {
	return s.appUseCase.ToggleFavorite(appName)
}
