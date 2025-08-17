package main

import (
	"context"
	"embed"
	"log"

	"launchy/internal/domain/services"
	"launchy/internal/infrastructure/repositories"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed frontend/dist
var assets embed.FS

// App struct
type App struct {
	appService *services.AppService
}

// NewApp creates a new App application struct
func NewApp() *App {
	// Initialize repository
	repo := repositories.NewDesktopEntryRepository()

	// Initialize service
	appService := services.NewAppService(repo)

	return &App{
		appService: appService,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	log.Println("Launchy starting up...")

	// Refresh app database on startup
	if err := a.appService.RefreshApps(ctx); err != nil {
		log.Printf("Warning: failed to refresh apps: %v", err)
	}
}

// domReady is called after the front-end dom has been loaded
func (a *App) domReady(ctx context.Context) {
	log.Println("DOM ready")
}

// beforeClose is called when the app is about to close,
// either by clicking the window close button or calling runtime.Quit
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	log.Println("Launchy shutting down...")
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	log.Println("Launchy shutdown complete")
}

// SearchApps searches for applications
func (a *App) SearchApps(query string) ([]map[string]interface{}, error) {
	ctx := context.Background()
	apps, err := a.appService.SearchApps(ctx, query)
	if err != nil {
		return nil, err
	}

	// Convert to map for JSON serialization
	result := make([]map[string]interface{}, len(apps))
	for i, app := range apps {
		result[i] = map[string]interface{}{
			"id":          app.ID,
			"name":        app.Name,
			"displayName": app.DisplayName,
			"description": app.Description,
			"execPath":    app.ExecPath,
			"iconPath":    app.IconPath,
			"categories":  app.Categories,
			"keywords":    app.Keywords,
			"useCount":    app.UseCount,
			"isFavorite":  app.IsFavorite,
		}
	}

	return result, nil
}

// LaunchApp launches an application
func (a *App) LaunchApp(appID string) error {
	ctx := context.Background()
	return a.appService.LaunchApp(ctx, appID)
}

// GetFavorites returns favorite applications
func (a *App) GetFavorites() ([]map[string]interface{}, error) {
	ctx := context.Background()
	apps, err := a.appService.GetFavorites(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]map[string]interface{}, len(apps))
	for i, app := range apps {
		result[i] = map[string]interface{}{
			"id":          app.ID,
			"name":        app.Name,
			"displayName": app.DisplayName,
			"description": app.Description,
			"execPath":    app.ExecPath,
			"iconPath":    app.IconPath,
			"categories":  app.Categories,
			"keywords":    app.Keywords,
			"useCount":    app.UseCount,
			"isFavorite":  app.IsFavorite,
		}
	}

	return result, nil
}

// GetRecentlyUsed returns recently used applications
func (a *App) GetRecentlyUsed(limit int) ([]map[string]interface{}, error) {
	ctx := context.Background()
	apps, err := a.appService.GetRecentlyUsed(ctx, limit)
	if err != nil {
		return nil, err
	}

	result := make([]map[string]interface{}, len(apps))
	for i, app := range apps {
		result[i] = map[string]interface{}{
			"id":          app.ID,
			"name":        app.Name,
			"displayName": app.DisplayName,
			"description": app.Description,
			"execPath":    app.ExecPath,
			"iconPath":    app.IconPath,
			"categories":  app.Categories,
			"keywords":    app.Keywords,
			"useCount":    app.UseCount,
			"isFavorite":  app.IsFavorite,
		}
	}

	return result, nil
}

// ToggleFavorite toggles the favorite status of an app
func (a *App) ToggleFavorite(appID string) error {
	ctx := context.Background()
	return a.appService.ToggleFavorite(ctx, appID)
}

// GetCategories returns all available categories
func (a *App) GetCategories() ([]string, error) {
	ctx := context.Background()
	return a.appService.GetCategories(ctx)
}

// GetAppsByCategory returns apps in a specific category
func (a *App) GetAppsByCategory(category string) ([]map[string]interface{}, error) {
	ctx := context.Background()
	apps, err := a.appService.GetAppsByCategory(ctx, category)
	if err != nil {
		return nil, err
	}

	result := make([]map[string]interface{}, len(apps))
	for i, app := range apps {
		result[i] = map[string]interface{}{
			"id":          app.ID,
			"name":        app.Name,
			"displayName": app.DisplayName,
			"description": app.Description,
			"execPath":    app.ExecPath,
			"iconPath":    app.IconPath,
			"categories":  app.Categories,
			"keywords":    app.Keywords,
			"useCount":    app.UseCount,
			"isFavorite":  app.IsFavorite,
		}
	}

	return result, nil
}

// RefreshApps refreshes the application database
func (a *App) RefreshApps() error {
	ctx := context.Background()
	return a.appService.RefreshApps(ctx)
}

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:             "Launchy",
		Width:             1000,
		Height:            700,
		MinWidth:          800,
		MinHeight:         600,
		MaxWidth:          0, // No maximum width
		MaxHeight:         0, // No maximum height
		DisableResize:     false,
		Fullscreen:        false,
		Frameless:         false,
		StartHidden:       false,
		HideWindowOnClose: false,
		BackgroundColour:  &options.RGBA{R: 255, G: 255, B: 255, A: 1},
		Assets:            assets,
		Menu:              nil,
		Logger:            nil,
		LogLevel:          0,
		OnStartup:         app.startup,
		OnDomReady:        app.domReady,
		OnBeforeClose:     app.beforeClose,
		OnShutdown:        app.shutdown,
		WindowStartState:  options.Normal,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
