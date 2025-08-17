package repositories

import (
	"context"
	"launchy/internal/domain/entities"
)

// AppRepository defines the interface for app data access
type AppRepository interface {
	// GetAll returns all applications
	GetAll(ctx context.Context) ([]*entities.App, error)
	
	// GetByID returns an app by its ID
	GetByID(ctx context.Context, id string) (*entities.App, error)
	
	// Search searches for apps by query
	Search(ctx context.Context, query string) ([]*entities.App, error)
	
	// GetFavorites returns all favorite apps
	GetFavorites(ctx context.Context) ([]*entities.App, error)
	
	// GetRecentlyUsed returns recently used apps
	GetRecentlyUsed(ctx context.Context, limit int) ([]*entities.App, error)
	
	// GetByCategory returns apps in a specific category
	GetByCategory(ctx context.Context, category string) ([]*entities.App, error)
	
	// Save saves or updates an app
	Save(ctx context.Context, app *entities.App) error
	
	// Delete deletes an app by ID
	Delete(ctx context.Context, id string) error
	
	// Refresh refreshes the app database from system
	Refresh(ctx context.Context) error
	
	// GetCategories returns all available categories
	GetCategories(ctx context.Context) ([]string, error)
}
