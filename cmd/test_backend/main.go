package main

import (
	"fmt"
	"log"

	"changeme/backend/service"
)

func main() {
	fmt.Println("Testing Launchy Backend...")

	// Create the app service
	appService := service.NewAppService()

	// Test GetApps
	fmt.Println("\nTesting GetApps...")
	apps, err := appService.GetApps()
	if err != nil {
		log.Fatalf("GetApps failed: %v", err)
	}

	fmt.Printf("Successfully loaded %d applications\n", len(apps))

	// Show first 5 apps
	fmt.Println("\nFirst 5 applications:")
	for i, app := range apps {
		if i >= 5 {
			break
		}
		fmt.Printf("  %d. %s (%s)\n", i+1, app.Name, app.Exec)
	}

	// Test search
	fmt.Println("\nTesting search for 'terminal'...")
	searchResults, err := appService.SearchApps("terminal")
	if err != nil {
		log.Fatalf("SearchApps failed: %v", err)
	}

	fmt.Printf("Found %d apps matching 'terminal'\n", len(searchResults))
	for i, app := range searchResults {
		if i >= 3 {
			break
		}
		fmt.Printf("  %d. %s (%s)\n", i+1, app.Name, app.Exec)
	}

	fmt.Println("\nBackend test completed successfully!")
}
