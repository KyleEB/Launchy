.PHONY: help dev build clean install-deps test lint format

# Default target
help:
	@echo "Launchy - CachyOS App Launcher"
	@echo ""
	@echo "Available commands:"
	@echo "  help         - Show this help message"
	@echo "  install-deps - Install all dependencies (Go and Node.js)"
	@echo "  dev          - Start development server with hot reload"
	@echo "  build        - Build for production"
	@echo "  clean        - Clean build artifacts"
	@echo "  test         - Run tests"
	@echo "  lint         - Run linters"
	@echo "  format       - Format code"
	@echo "  install      - Install the application"
	@echo "  package      - Package the application for distribution"

# Install dependencies
install-deps:
	@echo "Installing Go dependencies..."
	go mod tidy
	@echo "Installing frontend dependencies..."
	cd frontend && npm install
	@echo "Dependencies installed successfully!"

# Development mode
dev:
	@echo "Starting development server..."
	wails dev

# Build for production
build:
	@echo "Building for production..."
	wails build

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -rf build/
	rm -rf dist/
	rm -rf frontend/dist/
	rm -rf frontend/build/
	@echo "Clean complete!"

# Run tests
test:
	@echo "Running tests..."
	go test ./...
	cd frontend && npm test

# Run linters
lint:
	@echo "Running linters..."
	golangci-lint run
	cd frontend && npm run lint

# Format code
format:
	@echo "Formatting code..."
	go fmt ./...
	cd frontend && npm run format

# Install the application
install: build
	@echo "Installing Launchy..."
	sudo cp build/bin/Launchy /usr/local/bin/launchy
	@echo "Launchy installed successfully!"

# Package for distribution
package:
	@echo "Packaging application..."
	wails build -package
	@echo "Package created in build/"

# Build for specific platforms
build-linux:
	@echo "Building for Linux..."
	wails build -platform linux/amd64

build-windows:
	@echo "Building for Windows..."
	wails build -platform windows/amd64

build-macos:
	@echo "Building for macOS..."
	wails build -platform darwin/amd64

# Development utilities
frontend-dev:
	@echo "Starting frontend development server..."
	cd frontend && npm run dev

frontend-build:
	@echo "Building frontend..."
	cd frontend && npm run build

# Database operations
refresh-apps:
	@echo "Refreshing application database..."
	go run . refresh

# Quick development setup
setup: install-deps
	@echo "Development environment setup complete!"
	@echo "Run 'make dev' to start development server"
