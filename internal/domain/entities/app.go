package entities

import (
	"strings"
	"time"
)

// App represents an application that can be launched
type App struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	DisplayName string    `json:"displayName"`
	Description string    `json:"description"`
	ExecPath    string    `json:"execPath"`
	IconPath    string    `json:"iconPath"`
	Categories  []string  `json:"categories"`
	Keywords    []string  `json:"keywords"`
	InstalledAt time.Time `json:"installedAt"`
	LastUsed    time.Time `json:"lastUsed"`
	UseCount    int       `json:"useCount"`
	IsFavorite  bool      `json:"isFavorite"`
}

// NewApp creates a new App instance
func NewApp(name, displayName, execPath string) *App {
	return &App{
		ID:          generateID(),
		Name:        name,
		DisplayName: displayName,
		ExecPath:    execPath,
		Categories:  []string{},
		Keywords:    []string{},
		InstalledAt: time.Now(),
		UseCount:    0,
		IsFavorite:  false,
	}
}

// IncrementUseCount increases the use count and updates last used time
func (a *App) IncrementUseCount() {
	a.UseCount++
	a.LastUsed = time.Now()
}

// ToggleFavorite toggles the favorite status
func (a *App) ToggleFavorite() {
	a.IsFavorite = !a.IsFavorite
}

// AddCategory adds a category to the app
func (a *App) AddCategory(category string) {
	// Don't add empty categories
	if strings.TrimSpace(category) == "" {
		return
	}

	for _, cat := range a.Categories {
		if cat == category {
			return
		}
	}
	a.Categories = append(a.Categories, category)
}

// AddKeyword adds a keyword to the app
func (a *App) AddKeyword(keyword string) {
	// Don't add empty keywords
	if strings.TrimSpace(keyword) == "" {
		return
	}

	for _, kw := range a.Keywords {
		if kw == keyword {
			return
		}
	}
	a.Keywords = append(a.Keywords, keyword)
}

// generateID generates a simple ID for the app
func generateID() string {
	return time.Now().Format("20060102150405")
}
