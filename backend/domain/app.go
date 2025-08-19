package domain

import (
	"fmt"
	"strings"
)

// App represents an application with its metadata
type App struct {
	Name       string
	Exec       string
	Icon       string
	Comment    string
	Categories string
	IsFavorite bool
}

// NewApp creates a new App instance with validation
func NewApp(name, exec, icon, comment, categories string) (*App, error) {
	if strings.TrimSpace(name) == "" {
		return nil, fmt.Errorf("app name cannot be empty")
	}
	if strings.TrimSpace(exec) == "" {
		return nil, fmt.Errorf("app exec cannot be empty")
	}

	return &App{
		Name:       strings.TrimSpace(name),
		Exec:       strings.TrimSpace(exec),
		Icon:       strings.TrimSpace(icon),
		Comment:    strings.TrimSpace(comment),
		Categories: strings.TrimSpace(categories),
		IsFavorite: false,
	}, nil
}

// ToggleFavorite toggles the favorite status of the app
func (a *App) ToggleFavorite() {
	a.IsFavorite = !a.IsFavorite
}

// SetFavorite sets the favorite status of the app
func (a *App) SetFavorite(isFavorite bool) {
	a.IsFavorite = isFavorite
}

// MatchesSearch checks if the app matches the search query
func (a *App) MatchesSearch(query string) bool {
	query = strings.ToLower(strings.TrimSpace(query))
	if query == "" {
		return true
	}

	return strings.Contains(strings.ToLower(a.Name), query) ||
		strings.Contains(strings.ToLower(a.Comment), query) ||
		strings.Contains(strings.ToLower(a.Categories), query) ||
		strings.Contains(strings.ToLower(a.Exec), query)
}

// ToAppInfo converts the domain App to the API AppInfo struct
func (a *App) ToAppInfo() AppInfo {
	return AppInfo{
		Name:       a.Name,
		Exec:       a.Exec,
		Icon:       a.Icon,
		Comment:    a.Comment,
		Categories: a.Categories,
		IsFavorite: a.IsFavorite,
	}
}

// AppInfo represents the API response structure for an application
type AppInfo struct {
	Name       string `json:"name"`
	Exec       string `json:"exec"`
	Icon       string `json:"icon"`
	Comment    string `json:"comment"`
	Categories string `json:"categories"`
	IsFavorite bool   `json:"isFavorite"`
}
