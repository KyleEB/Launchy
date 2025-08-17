package repositories

import (
	"bufio"
	"context"
	"fmt"
	"launchy/internal/domain/entities"
	"launchy/internal/domain/repositories"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// DesktopEntryRepository implements AppRepository using desktop files
type DesktopEntryRepository struct {
	apps map[string]*entities.App
}

// Ensure DesktopEntryRepository implements AppRepository interface
var _ repositories.AppRepository = (*DesktopEntryRepository)(nil)

// NewDesktopEntryRepository creates a new repository instance
func NewDesktopEntryRepository() *DesktopEntryRepository {
	return &DesktopEntryRepository{
		apps: make(map[string]*entities.App),
	}
}

// GetAll returns all applications
func (r *DesktopEntryRepository) GetAll(ctx context.Context) ([]*entities.App, error) {
	if err := r.ensureLoaded(ctx); err != nil {
		return nil, err
	}

	apps := make([]*entities.App, 0, len(r.apps))
	for _, app := range r.apps {
		apps = append(apps, app)
	}

	sort.Slice(apps, func(i, j int) bool {
		return apps[i].DisplayName < apps[j].DisplayName
	})

	return apps, nil
}

// GetByID returns an app by its ID
func (r *DesktopEntryRepository) GetByID(ctx context.Context, id string) (*entities.App, error) {
	if err := r.ensureLoaded(ctx); err != nil {
		return nil, err
	}

	app, exists := r.apps[id]
	if !exists {
		return nil, fmt.Errorf("app not found: %s", id)
	}

	return app, nil
}

// Search searches for apps by query
func (r *DesktopEntryRepository) Search(ctx context.Context, query string) ([]*entities.App, error) {
	if err := r.ensureLoaded(ctx); err != nil {
		return nil, err
	}

	query = strings.ToLower(query)
	var results []*entities.App

	for _, app := range r.apps {
		if r.matchesQuery(app, query) {
			results = append(results, app)
		}
	}

	return results, nil
}

// GetFavorites returns all favorite apps
func (r *DesktopEntryRepository) GetFavorites(ctx context.Context) ([]*entities.App, error) {
	if err := r.ensureLoaded(ctx); err != nil {
		return nil, err
	}

	var favorites []*entities.App
	for _, app := range r.apps {
		if app.IsFavorite {
			favorites = append(favorites, app)
		}
	}

	sort.Slice(favorites, func(i, j int) bool {
		return favorites[i].DisplayName < favorites[j].DisplayName
	})

	return favorites, nil
}

// GetRecentlyUsed returns recently used apps
func (r *DesktopEntryRepository) GetRecentlyUsed(ctx context.Context, limit int) ([]*entities.App, error) {
	if err := r.ensureLoaded(ctx); err != nil {
		return nil, err
	}

	var recent []*entities.App
	for _, app := range r.apps {
		if !app.LastUsed.IsZero() {
			recent = append(recent, app)
		}
	}

	sort.Slice(recent, func(i, j int) bool {
		return recent[i].LastUsed.After(recent[j].LastUsed)
	})

	if limit > 0 && limit < len(recent) {
		recent = recent[:limit]
	}

	return recent, nil
}

// GetByCategory returns apps in a specific category
func (r *DesktopEntryRepository) GetByCategory(ctx context.Context, category string) ([]*entities.App, error) {
	if err := r.ensureLoaded(ctx); err != nil {
		return nil, err
	}

	var results []*entities.App
	for _, app := range r.apps {
		for _, cat := range app.Categories {
			if strings.EqualFold(cat, category) {
				results = append(results, app)
				break
			}
		}
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].DisplayName < results[j].DisplayName
	})

	return results, nil
}

// Save saves or updates an app
func (r *DesktopEntryRepository) Save(ctx context.Context, app *entities.App) error {
	r.apps[app.ID] = app
	return nil
}

// Delete deletes an app by ID
func (r *DesktopEntryRepository) Delete(ctx context.Context, id string) error {
	delete(r.apps, id)
	return nil
}

// Refresh refreshes the app database from system
func (r *DesktopEntryRepository) Refresh(ctx context.Context) error {
	r.apps = make(map[string]*entities.App)
	return r.loadDesktopEntries(ctx)
}

// GetCategories returns all available categories
func (r *DesktopEntryRepository) GetCategories(ctx context.Context) ([]string, error) {
	if err := r.ensureLoaded(ctx); err != nil {
		return nil, err
	}

	categorySet := make(map[string]bool)
	for _, app := range r.apps {
		for _, category := range app.Categories {
			categorySet[category] = true
		}
	}

	categories := make([]string, 0, len(categorySet))
	for category := range categorySet {
		categories = append(categories, category)
	}

	sort.Strings(categories)
	return categories, nil
}

// ensureLoaded ensures the apps are loaded
func (r *DesktopEntryRepository) ensureLoaded(ctx context.Context) error {
	if len(r.apps) == 0 {
		return r.loadDesktopEntries(ctx)
	}
	return nil
}

// loadDesktopEntries loads desktop entries from system directories
func (r *DesktopEntryRepository) loadDesktopEntries(ctx context.Context) error {
	desktopDirs := []string{
		"/usr/share/applications",
		"/usr/local/share/applications",
		filepath.Join(os.Getenv("HOME"), ".local/share/applications"),
	}

	for _, dir := range desktopDirs {
		if err := r.scanDirectory(ctx, dir); err != nil {
			// Log error but continue with other directories
			fmt.Printf("Warning: failed to scan %s: %v\n", dir, err)
		}
	}

	return nil
}

// scanDirectory scans a directory for .desktop files
func (r *DesktopEntryRepository) scanDirectory(ctx context.Context, dir string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".desktop") {
			continue
		}

		filePath := filepath.Join(dir, entry.Name())
		if app := r.parseDesktopFile(filePath); app != nil {
			r.apps[app.ID] = app
		}
	}

	return nil
}

// isGPSApplication checks if the application is GPS-related
func (r *DesktopEntryRepository) isGPSApplication(app *entities.App) bool {
	// Check if the application name contains GPS-related terms
	gpsTerms := []string{"gps", "gpsd", "xgps", "xgpsspeed", "gpsbabel", "gpscorrelate", "gpsprune", "gpsdrive"}

	appNameLower := strings.ToLower(app.Name)
	displayNameLower := strings.ToLower(app.DisplayName)
	descriptionLower := strings.ToLower(app.Description)

	for _, term := range gpsTerms {
		if strings.Contains(appNameLower, term) ||
			strings.Contains(displayNameLower, term) ||
			strings.Contains(descriptionLower, term) {
			return true
		}
	}

	// Check if the executable path contains GPS-related terms
	execPathLower := strings.ToLower(app.ExecPath)
	for _, term := range gpsTerms {
		if strings.Contains(execPathLower, term) {
			return true
		}
	}

	// Check if the application is in GPS-related categories
	gpsCategories := []string{"gps", "navigation", "maps", "geography"}
	for _, category := range app.Categories {
		categoryLower := strings.ToLower(category)
		for _, gpsCategory := range gpsCategories {
			if strings.Contains(categoryLower, gpsCategory) {
				return true
			}
		}
	}

	return false
}

// parseDesktopFile parses a .desktop file and returns an App
func (r *DesktopEntryRepository) parseDesktopFile(filePath string) *entities.App {
	file, err := os.Open(filePath)
	if err != nil {
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	app := &entities.App{
		ID:          filePath,
		Categories:  []string{},
		Keywords:    []string{},
		InstalledAt: time.Now(),
		UseCount:    0,
		IsFavorite:  false,
	}

	inDesktopEntry := false
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "[Desktop Entry]" {
			inDesktopEntry = true
			continue
		}

		if !inDesktopEntry || line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		if strings.HasPrefix(line, "[") {
			break // End of Desktop Entry section
		}

		if strings.Contains(line, "=") {
			parts := strings.SplitN(line, "=", 2)
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])

			switch key {
			case "Name":
				app.DisplayName = value
				if app.Name == "" {
					app.Name = value
				}
			case "Comment":
				app.Description = value
			case "Exec":
				app.ExecPath = r.cleanExecPath(value)
			case "Icon":
				app.IconPath = value
			case "Categories":
				// Split categories and filter out empty strings
				categories := strings.Split(value, ";")
				var filteredCategories []string
				for _, cat := range categories {
					cat = strings.TrimSpace(cat)
					if cat != "" {
						filteredCategories = append(filteredCategories, cat)
					}
				}
				app.Categories = filteredCategories
			case "Keywords":
				// Split keywords and filter out empty strings
				keywords := strings.Split(value, ";")
				var filteredKeywords []string
				for _, kw := range keywords {
					kw = strings.TrimSpace(kw)
					if kw != "" {
						filteredKeywords = append(filteredKeywords, kw)
					}
				}
				app.Keywords = filteredKeywords
			}
		}
	}

	// Only return valid apps
	if app.Name != "" && app.ExecPath != "" {
		// Skip GPS applications to avoid GPSD connection issues
		if r.isGPSApplication(app) {
			fmt.Printf("Skipping GPS application: %s (to avoid GPSD connection issues)\n", app.Name)
			return nil
		}
		return app
	}

	return nil
}

// cleanExecPath cleans the exec path from desktop file format
func (r *DesktopEntryRepository) cleanExecPath(execPath string) string {
	// Remove %U, %u, %F, %f, etc.
	execPath = strings.ReplaceAll(execPath, "%U", "")
	execPath = strings.ReplaceAll(execPath, "%u", "")
	execPath = strings.ReplaceAll(execPath, "%F", "")
	execPath = strings.ReplaceAll(execPath, "%f", "")
	execPath = strings.ReplaceAll(execPath, "%k", "")

	// Get the first part (before any arguments)
	parts := strings.Fields(execPath)
	if len(parts) > 0 {
		return parts[0]
	}

	return execPath
}

// matchesQuery checks if an app matches the search query
func (r *DesktopEntryRepository) matchesQuery(app *entities.App, query string) bool {
	query = strings.ToLower(query)

	// Check name
	if strings.Contains(strings.ToLower(app.Name), query) {
		return true
	}

	// Check display name
	if strings.Contains(strings.ToLower(app.DisplayName), query) {
		return true
	}

	// Check description
	if strings.Contains(strings.ToLower(app.Description), query) {
		return true
	}

	// Check keywords
	for _, keyword := range app.Keywords {
		if strings.Contains(strings.ToLower(keyword), query) {
			return true
		}
	}

	// Check categories
	for _, category := range app.Categories {
		if strings.Contains(strings.ToLower(category), query) {
			return true
		}
	}

	return false
}
