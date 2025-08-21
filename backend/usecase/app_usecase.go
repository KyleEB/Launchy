package usecase

import (
	"fmt"
	"strings"

	"changeme/backend/domain"
	"changeme/backend/repository"
)

// AppUseCase handles application business logic
type AppUseCase struct {
	appRepo     repository.AppRepository
	appLauncher repository.AppLauncher
	discovered  bool
}

// NewAppUseCase creates a new AppUseCase instance
func NewAppUseCase(appRepo repository.AppRepository, appLauncher repository.AppLauncher) *AppUseCase {
	return &AppUseCase{
		appRepo:     appRepo,
		appLauncher: appLauncher,
	}
}

// GetAllApps returns all available applications
func (uc *AppUseCase) GetAllApps() ([]domain.AppInfo, error) {
	// Trigger app discovery on first call only
	if !uc.discovered {
		if err := uc.appRepo.DiscoverApps(); err != nil {
			return nil, fmt.Errorf("failed to discover apps: %w", err)
		}
		uc.discovered = true
	}

	apps, err := uc.appRepo.GetAllApps()
	if err != nil {
		return nil, fmt.Errorf("failed to get apps: %w", err)
	}

	// Convert domain apps to API response
	appInfos := make([]domain.AppInfo, len(apps))
	for i, app := range apps {
		appInfos[i] = app.ToAppInfo()
	}

	return appInfos, nil
}

// GetFavorites returns all favorited applications
func (uc *AppUseCase) GetFavorites() ([]domain.AppInfo, error) {
	apps, err := uc.appRepo.GetFavorites()
	if err != nil {
		return nil, fmt.Errorf("failed to get favorites: %w", err)
	}

	// Convert domain apps to API response
	appInfos := make([]domain.AppInfo, len(apps))
	for i, app := range apps {
		appInfos[i] = app.ToAppInfo()
	}

	return appInfos, nil
}

// SearchApps searches for applications based on the query
func (uc *AppUseCase) SearchApps(query string) ([]domain.AppInfo, error) {
	if strings.TrimSpace(query) == "" {
		return uc.GetAllApps()
	}

	apps, err := uc.appRepo.GetAllApps()
	if err != nil {
		return nil, fmt.Errorf("failed to get apps for search: %w", err)
	}

	query = strings.ToLower(strings.TrimSpace(query))
	var results []domain.AppInfo

	for _, app := range apps {
		if app.MatchesSearch(query) {
			results = append(results, app.ToAppInfo())
		}
	}

	return results, nil
}

// LaunchApp launches an application by its exec path
func (uc *AppUseCase) LaunchApp(execPath string) error {
	if strings.TrimSpace(execPath) == "" {
		return fmt.Errorf("exec path cannot be empty")
	}

	return uc.appLauncher.Launch(execPath)
}

// ToggleFavorite toggles the favorite status of an application
func (uc *AppUseCase) ToggleFavorite(appName string) error {
	if strings.TrimSpace(appName) == "" {
		return fmt.Errorf("app name cannot be empty")
	}

	favorites, err := uc.appRepo.LoadFavorites()
	if err != nil {
		return fmt.Errorf("failed to load favorites: %w", err)
	}

	favorites[appName] = !favorites[appName]

	if err := uc.appRepo.SaveFavorites(favorites); err != nil {
		return fmt.Errorf("failed to save favorites: %w", err)
	}

	return nil
}

// DiscoverApps triggers app discovery
func (uc *AppUseCase) DiscoverApps() error {
	return uc.appRepo.DiscoverApps()
}
