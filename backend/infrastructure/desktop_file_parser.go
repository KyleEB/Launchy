package infrastructure

import (
	"fmt"
	"os"
	"strings"

	"changeme/backend/domain"
)

// DesktopFileParser defines the interface for parsing desktop files
type DesktopFileParser interface {
	ParseDesktopFile(filePath string) (*domain.App, error)
}

// DesktopFileParserImpl implements DesktopFileParser
type DesktopFileParserImpl struct{}

// NewDesktopFileParser creates a new DesktopFileParser
func NewDesktopFileParser() DesktopFileParser {
	return &DesktopFileParserImpl{}
}

// ParseDesktopFile parses a desktop file and returns App
func (p *DesktopFileParserImpl) ParseDesktopFile(filePath string) (*domain.App, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic parsing desktop file %s: %v\n", filePath, r)
		}
	}()

	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read desktop file: %w", err)
	}

	lines := strings.Split(string(content), "\n")
	var name, exec, icon, comment, categories string
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
			name = strings.TrimPrefix(line, "Name=")
		} else if strings.HasPrefix(line, "Exec=") {
			exec = strings.TrimPrefix(line, "Exec=")
		} else if strings.HasPrefix(line, "Icon=") {
			icon = strings.TrimPrefix(line, "Icon=")
		} else if strings.HasPrefix(line, "Comment=") {
			comment = strings.TrimPrefix(line, "Comment=")
		} else if strings.HasPrefix(line, "Categories=") {
			categories = strings.TrimPrefix(line, "Categories=")
		}
	}

	// Only return apps that have a name and exec
	if name != "" && exec != "" {
		return domain.NewApp(name, exec, icon, comment, categories)
	}

	return nil, fmt.Errorf("invalid desktop file: missing name or exec")
}
