package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

// AppInfo represents an application with its metadata
type AppInfo struct {
	Name       string `json:"name"`
	Exec       string `json:"exec"`
	Icon       string `json:"icon"`
	Comment    string `json:"comment"`
	Categories string `json:"categories"`
	IsFavorite bool   `json:"isFavorite"`
}

// AppRepository defines the interface for application data operations
type AppRepository interface {
	GetAllApps() ([]AppInfo, error)
	GetFavorites() ([]AppInfo, error)
	SaveFavorites(favorites map[string]bool) error
	LoadFavorites() (map[string]bool, error)
}

// AppLauncher defines the interface for launching applications
type AppLauncher interface {
	Launch(execPath string) error
}

// DesktopFileParser defines the interface for parsing desktop files
type DesktopFileParser interface {
	ParseDesktopFile(filePath string) (AppInfo, error)
}

// AppService provides application management functionality
type AppService struct {
	repository AppRepository
	launcher   AppLauncher
	parser     DesktopFileParser
}

// NewAppService creates a new instance of AppService with dependencies
func NewAppService() *AppService {
	homeDir, _ := os.UserHomeDir()
	configPath := filepath.Join(homeDir, ".config", "launchy", "favorites.json")
	
	repository := NewFileAppRepository(configPath)
	launcher := NewSystemAppLauncher()
	parser := NewDesktopFileParser()
	
	return &AppService{
		repository: repository,
		launcher:   launcher,
		parser:     parser,
	}
}

// GetApps returns all available applications
func (s *AppService) GetApps() ([]AppInfo, error) {
	return s.repository.GetAllApps()
}

// GetFavorites returns all favorited applications
func (s *AppService) GetFavorites() ([]AppInfo, error) {
	return s.repository.GetFavorites()
}

// SearchApps searches for applications based on the query
func (s *AppService) SearchApps(query string) ([]AppInfo, error) {
	if strings.TrimSpace(query) == "" {
		return s.GetApps()
	}
	
	apps, err := s.GetApps()
	if err != nil {
		return nil, fmt.Errorf("failed to get apps for search: %w", err)
	}
	
	query = strings.ToLower(strings.TrimSpace(query))
	var results []AppInfo
	
	for _, app := range apps {
		if s.matchesSearch(app, query) {
			results = append(results, app)
		}
	}
	
	return results, nil
}

// LaunchApp launches an application by its exec path
func (s *AppService) LaunchApp(execPath string) error {
	if strings.TrimSpace(execPath) == "" {
		return fmt.Errorf("exec path cannot be empty")
	}
	
	return s.launcher.Launch(execPath)
}

// ToggleFavorite toggles the favorite status of an application
func (s *AppService) ToggleFavorite(appName string) error {
	if strings.TrimSpace(appName) == "" {
		return fmt.Errorf("app name cannot be empty")
	}
	
	favorites, err := s.repository.LoadFavorites()
	if err != nil {
		return fmt.Errorf("failed to load favorites: %w", err)
	}
	
	favorites[appName] = !favorites[appName]
	
	if err := s.repository.SaveFavorites(favorites); err != nil {
		return fmt.Errorf("failed to save favorites: %w", err)
	}
	
	return nil
}

// matchesSearch checks if an app matches the search query
func (s *AppService) matchesSearch(app AppInfo, query string) bool {
	return strings.Contains(strings.ToLower(app.Name), query) ||
		   strings.Contains(strings.ToLower(app.Comment), query) ||
		   strings.Contains(strings.ToLower(app.Categories), query)
}

// FileAppRepository implements AppRepository using file system
type FileAppRepository struct {
	configPath string
	apps       []AppInfo
	favorites  map[string]bool
}

// NewFileAppRepository creates a new FileAppRepository
func NewFileAppRepository(configPath string) *FileAppRepository {
	repo := &FileAppRepository{
		configPath: configPath,
		apps:       []AppInfo{},
		favorites:  make(map[string]bool),
	}
	
	repo.loadFavorites()
	repo.discoverApps()
	return repo
}

// GetAllApps returns all discovered applications
func (r *FileAppRepository) GetAllApps() ([]AppInfo, error) {
	return r.apps, nil
}

// GetFavorites returns favorited applications
func (r *FileAppRepository) GetFavorites() ([]AppInfo, error) {
	var favorites []AppInfo
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
	for i := range r.apps {
		r.apps[i].IsFavorite = favorites[r.apps[i].Name]
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

// discoverApps scans the system for desktop applications
func (r *FileAppRepository) discoverApps() {
	desktopDirs := []string{
		"/usr/share/applications",
		"/usr/local/share/applications",
		filepath.Join(os.Getenv("HOME"), ".local/share/applications"),
	}
	
	parser := NewDesktopFileParser()
	
	for _, dir := range desktopDirs {
		r.scanDesktopDirectory(dir, parser)
	}
	
	// Sort apps by name
	sort.Slice(r.apps, func(i, j int) bool {
		return r.apps[i].Name < r.apps[j].Name
	})
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
				app.IsFavorite = r.favorites[app.Name]
				r.apps = append(r.apps, app)
			}
		}
	}
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

// SystemAppLauncher implements AppLauncher using system commands
type SystemAppLauncher struct{}

// NewSystemAppLauncher creates a new SystemAppLauncher
func NewSystemAppLauncher() *SystemAppLauncher {
	return &SystemAppLauncher{}
}

// Launch executes an application
func (l *SystemAppLauncher) Launch(execPath string) error {
	parts := strings.Fields(execPath)
	if len(parts) == 0 {
		return fmt.Errorf("invalid exec command")
	}
	
	cmd := exec.Command(parts[0], parts[1:]...)
	return cmd.Start()
}

// DesktopFileParserImpl implements DesktopFileParser
type DesktopFileParserImpl struct{}

// NewDesktopFileParser creates a new DesktopFileParser
func NewDesktopFileParser() DesktopFileParser {
	return &DesktopFileParserImpl{}
}

// ParseDesktopFile parses a desktop file and returns AppInfo
func (p *DesktopFileParserImpl) ParseDesktopFile(filePath string) (AppInfo, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return AppInfo{}, fmt.Errorf("failed to read desktop file: %w", err)
	}
	
	lines := strings.Split(string(content), "\n")
	app := AppInfo{}
	inDesktopEntry := false
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		
		// Check for section headers
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			if line == "[Desktop Entry]" {
				inDesktopEntry = true
			} else {
				inDesktopEntry = false
			}
			continue
		}
		
		// Only process lines in the [Desktop Entry] section
		if !inDesktopEntry {
			continue
		}
		
		if strings.HasPrefix(line, "Name=") {
			app.Name = strings.TrimPrefix(line, "Name=")
		} else if strings.HasPrefix(line, "Exec=") {
			app.Exec = strings.TrimPrefix(line, "Exec=")
		} else if strings.HasPrefix(line, "Icon=") {
			app.Icon = strings.TrimPrefix(line, "Icon=")
		} else if strings.HasPrefix(line, "Comment=") {
			app.Comment = strings.TrimPrefix(line, "Comment=")
		} else if strings.HasPrefix(line, "Categories=") {
			app.Categories = strings.TrimPrefix(line, "Categories=")
		}
	}
	
	// Only return apps that have a name and exec
	if app.Name != "" && app.Exec != "" {
		return app, nil
	}
	
	return AppInfo{}, fmt.Errorf("invalid desktop file: missing name or exec")
}
