package repository

import "changeme/backend/domain"

// AppRepository defines the interface for application data operations
type AppRepository interface {
	GetAllApps() ([]*domain.App, error)
	GetFavorites() ([]*domain.App, error)
	SaveFavorites(favorites map[string]bool) error
	LoadFavorites() (map[string]bool, error)
	DiscoverApps() error
}
