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
	fmt.Printf("GetAllApps called - returning %d discovered applications\n", len(r.apps))

	if len(r.apps) == 0 {
		fmt.Println("WARNING: No applications found! This might indicate an issue with app discovery.")
		fmt.Println("Checking if app discovery completed properly...")
	}

	// Log first few apps for debugging
	for i, app := range r.apps {
		if i < 5 { // Only log first 5 apps to avoid spam
			fmt.Printf("App %d: %s (%s)\n", i+1, app.Name, app.Exec)
		}
	}

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
	fmt.Println("Starting app discovery...")

	// Discover desktop applications
	desktopDirs := []string{
		"/usr/share/applications",
		"/usr/local/share/applications",
		filepath.Join(os.Getenv("HOME"), ".local/share/applications"),
	}

	parser := NewDesktopFileParser()

	for _, dir := range desktopDirs {
		fmt.Printf("Scanning desktop directory: %s\n", dir)
		r.scanDesktopDirectory(dir, parser)
	}

	// Discover command line executables from PATH
	fmt.Println("Scanning PATH for executables...")
	r.scanPathExecutables()

	// Sort apps by name
	sort.Slice(r.apps, func(i, j int) bool {
		return r.apps[i].Name < r.apps[j].Name
	})

	fmt.Printf("App discovery completed. Found %d applications\n", len(r.apps))
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
		fmt.Printf("Error reading directory %s: %v\n", dir, err)
		return
	}

	fileCount := 0
	appCount := 0
	
	fmt.Printf("Scanning %s for .desktop files...\n", dir)
	
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".desktop") {
			fileCount++
			filePath := filepath.Join(dir, entry.Name())
			if app, err := parser.ParseDesktopFile(filePath); err == nil && app.Name != "" {
				app.SetFavorite(r.favorites[app.Name])
				r.apps = append(r.apps, app)
				appCount++
				if appCount <= 3 { // Log first 3 successful parses
					fmt.Printf("  ✓ Parsed: %s -> %s\n", entry.Name(), app.Name)
				}
			} else {
				if fileCount <= 5 { // Log first 5 failed parses
					fmt.Printf("  ✗ Failed to parse: %s (%v)\n", entry.Name(), err)
				}
			}
		}
	}
	
	fmt.Printf("Scanned %s: found %d .desktop files, parsed %d applications\n", dir, fileCount, appCount)
}

// scanPathExecutables scans the PATH environment variable for executable files
func (r *FileAppRepository) scanPathExecutables() {
	path := os.Getenv("PATH")
	paths := strings.Split(path, ":")
	execCount := 0

	fmt.Printf("Scanning PATH for executables...\n")
	fmt.Printf("PATH directories: %v\n", paths)

	for _, dir := range paths {
		if dir == "" {
			continue
		}
		
		entries, err := os.ReadDir(dir)
		if err != nil {
			fmt.Printf("Error reading PATH directory %s: %v\n", dir, err)
			continue
		}

		dirExecCount := 0
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
							execCount++
							dirExecCount++
							if execCount <= 5 { // Log first 5 executables
								fmt.Printf("  ✓ Found executable: %s (%s)\n", entry.Name(), dir)
							}
						}
					}
				}
			}
		}
		
		if dirExecCount > 0 {
			fmt.Printf("  Directory %s: found %d executables\n", dir, dirExecCount)
		}
	}
	
	fmt.Printf("Found %d command line executables total\n", execCount)
}
