package infrastructure

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"changeme/backend/domain"
)

// FileAppRepository implements AppRepository using file system
type FileAppRepository struct {
	configPath string
	apps       []*domain.App
	favorites  map[string]bool
}

// NewFileAppRepository creates a new FileAppRepository
func NewFileAppRepository(configPath string) *FileAppRepository {
	repo := &FileAppRepository{
		configPath: configPath,
		apps:       []*domain.App{},
		favorites:  make(map[string]bool),
	}

	repo.loadFavorites()
	return repo
}

// GetAllApps returns all discovered applications
func (r *FileAppRepository) GetAllApps() ([]*domain.App, error) {
	return r.apps, nil
}

// GetFavorites returns favorited applications
func (r *FileAppRepository) GetFavorites() ([]*domain.App, error) {
	var favorites []*domain.App
	for _, app := range r.apps {
		if app.IsFavorite {
			favorites = append(favorites, app)
		}
	}
	return favorites, nil
}

// SaveFavorites saves favorites to the configuration file
func (r *FileAppRepository) SaveFavorites(favorites map[string]bool) error {
	r.favorites = favorites

	// Update apps with favorite status
	for _, app := range r.apps {
		app.SetFavorite(favorites[app.Name])
	}

	data, err := json.Marshal(favorites)
	if err != nil {
		return fmt.Errorf("failed to marshal favorites: %w", err)
	}

	// Ensure config directory exists
	configDir := filepath.Dir(r.configPath)
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	if err := os.WriteFile(r.configPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write favorites file: %w", err)
	}

	return nil
}

// LoadFavorites loads favorites from the configuration file
func (r *FileAppRepository) LoadFavorites() (map[string]bool, error) {
	return r.favorites, nil
}

// DiscoverApps scans the system for desktop applications and command line executables
func (r *FileAppRepository) DiscoverApps() error {
	// Discover desktop applications
	desktopDirs := []string{
		"/usr/share/applications",
		"/usr/local/share/applications",
		filepath.Join(os.Getenv("HOME"), ".local/share/applications"),
	}

	parser := NewDesktopFileParser()

	for _, dir := range desktopDirs {
		r.scanDesktopDirectory(dir, parser)
	}

	// Discover command line executables from PATH
	r.scanPathExecutables()

	// Sort apps by name
	sort.Slice(r.apps, func(i, j int) bool {
		return r.apps[i].Name < r.apps[j].Name
	})

	return nil
}

// loadFavorites loads favorites from the configuration file
func (r *FileAppRepository) loadFavorites() {
	data, err := os.ReadFile(r.configPath)
	if err != nil {
		return
	}

	if err := json.Unmarshal(data, &r.favorites); err != nil {
		// If unmarshaling fails, start with empty favorites
		r.favorites = make(map[string]bool)
	}
}

// scanDesktopDirectory scans a directory for desktop files
func (r *FileAppRepository) scanDesktopDirectory(dir string, parser DesktopFileParser) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return
	}

	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".desktop") {
			filePath := filepath.Join(dir, entry.Name())
			if app, err := parser.ParseDesktopFile(filePath); err == nil && app.Name != "" {
				app.SetFavorite(r.favorites[app.Name])
				r.apps = append(r.apps, app)
			}
		}
	}
}

// scanPathExecutables scans the PATH environment variable for executable files
func (r *FileAppRepository) scanPathExecutables() {
	path := os.Getenv("PATH")
	paths := strings.Split(path, ":")

	for _, dir := range paths {
		if dir == "" {
			continue
		}

		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}

		for _, entry := range entries {
			if !entry.IsDir() {
				filePath := filepath.Join(dir, entry.Name())

				// Check if file is executable
				if info, err := os.Stat(filePath); err == nil {
					if info.Mode()&0111 != 0 { // Check if executable
						app, err := domain.NewApp(
							entry.Name(),
							filePath,
							"",
							fmt.Sprintf("Command line executable from %s", dir),
							"Command Line",
						)
						if err == nil {
							app.SetFavorite(r.favorites[entry.Name()])
							r.apps = append(r.apps, app)
						}
					}
				}
			}
		}
	}
}
