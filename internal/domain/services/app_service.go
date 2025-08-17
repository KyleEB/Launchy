package services

import (
	"context"
	"fmt"
	"os/exec"
	"sort"
	"strings"
	"launchy/internal/domain/entities"
	"launchy/internal/domain/repositories"
)

// AppService handles business logic for applications
type AppService struct {
	repo repositories.AppRepository
}

// NewAppService creates a new AppService instance
func NewAppService(repo repositories.AppRepository) *AppService {
	return &AppService{
		repo: repo,
	}
}

// LaunchApp launches an application and updates usage statistics
func (s *AppService) LaunchApp(ctx context.Context, appID string) error {
	app, err := s.repo.GetByID(ctx, appID)
	if err != nil {
		return fmt.Errorf("failed to get app: %w", err)
	}

	// Launch the application
	if err := s.executeApp(app); err != nil {
		return fmt.Errorf("failed to launch app: %w", err)
	}

	// Update usage statistics
	app.IncrementUseCount()
	if err := s.repo.Save(ctx, app); err != nil {
		return fmt.Errorf("failed to update app usage: %w", err)
	}

	return nil
}

// SearchApps searches for applications with smart ranking
func (s *AppService) SearchApps(ctx context.Context, query string) ([]*entities.App, error) {
	if strings.TrimSpace(query) == "" {
		return s.repo.GetAll(ctx)
	}

	apps, err := s.repo.Search(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to search apps: %w", err)
	}

	// Sort by relevance and usage
	sort.Slice(apps, func(i, j int) bool {
		return s.calculateRelevance(apps[i], query) > s.calculateRelevance(apps[j], query)
	})

	return apps, nil
}

// GetFavorites returns favorite applications
func (s *AppService) GetFavorites(ctx context.Context) ([]*entities.App, error) {
	return s.repo.GetFavorites(ctx)
}

// GetRecentlyUsed returns recently used applications
func (s *AppService) GetRecentlyUsed(ctx context.Context, limit int) ([]*entities.App, error) {
	return s.repo.GetRecentlyUsed(ctx, limit)
}

// ToggleFavorite toggles the favorite status of an app
func (s *AppService) ToggleFavorite(ctx context.Context, appID string) error {
	app, err := s.repo.GetByID(ctx, appID)
	if err != nil {
		return fmt.Errorf("failed to get app: %w", err)
	}

	app.ToggleFavorite()
	return s.repo.Save(ctx, app)
}

// RefreshApps refreshes the application database
func (s *AppService) RefreshApps(ctx context.Context) error {
	return s.repo.Refresh(ctx)
}

// GetCategories returns all available categories
func (s *AppService) GetCategories(ctx context.Context) ([]string, error) {
	return s.repo.GetCategories(ctx)
}

// GetAppsByCategory returns apps in a specific category
func (s *AppService) GetAppsByCategory(ctx context.Context, category string) ([]*entities.App, error) {
	return s.repo.GetByCategory(ctx, category)
}

// executeApp launches an application using exec.Command
func (s *AppService) executeApp(app *entities.App) error {
	cmd := exec.Command(app.ExecPath)
	return cmd.Start()
}

// calculateRelevance calculates the relevance score for search ranking
func (s *AppService) calculateRelevance(app *entities.App, query string) float64 {
	query = strings.ToLower(query)
	score := 0.0

	// Exact name match gets highest score
	if strings.Contains(strings.ToLower(app.Name), query) {
		score += 100
	}

	// Display name match
	if strings.Contains(strings.ToLower(app.DisplayName), query) {
		score += 80
	}

	// Description match
	if strings.Contains(strings.ToLower(app.Description), query) {
		score += 40
	}

	// Keyword match
	for _, keyword := range app.Keywords {
		if strings.Contains(strings.ToLower(keyword), query) {
			score += 60
		}
	}

	// Category match
	for _, category := range app.Categories {
		if strings.Contains(strings.ToLower(category), query) {
			score += 30
		}
	}

	// Boost favorites and frequently used apps
	if app.IsFavorite {
		score += 20
	}

	// Boost based on usage count (but not too much)
	usageBoost := float64(app.UseCount) * 0.1
	if usageBoost > 10 {
		usageBoost = 10
	}
	score += usageBoost

	return score
}
