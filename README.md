# 🚀 Launchy - Application Launcher for CachyOS

A modern, fast application launcher built with Wails3 and Svelte for CachyOS and other Linux distributions.

## Features

- **🔍 Smart Search**: Search applications by name, description, or categories
- **🎨 Modern UI**: Beautiful, responsive interface with smooth animations
- **⚡ Fast Performance**: Built with Wails3 for native performance
- **🔧 Clean Architecture**: Follows Clean Code and Clean Architecture patterns

## Architecture

This project follows Clean Architecture principles with clear separation of concerns:

### Backend (Go)
- **AppService**: Main service layer handling business logic
- **AppRepository**: Interface for data operations (implemented as FileAppRepository)
- **AppLauncher**: Interface for launching applications (implemented as SystemAppLauncher)
- **DesktopFileParser**: Interface for parsing .desktop files

### Frontend (Svelte)
- **App.svelte**: Main application component
- **SearchBar.svelte**: Search functionality with debouncing
- **AppCard.svelte**: Individual application card component

## Getting Started

### Prerequisites

- Go 1.21 or later
- Node.js 18 or later
- Wails3 CLI

### Local Setup

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd Launchy
   ```

2. Install Wails3 CLI:
   ```bash
   go install github.com/wailsapp/wails/v3/cmd/wails3@latest
   ```

3. Install frontend dependencies:
   ```bash
   cd frontend
   npm install
   cd ..
   ```

### Development

To run the application in development mode:

```bash
wails3 dev
```

This will start the application with hot-reloading for both frontend and backend changes.

### Building

To build the application for production:

```bash
wails3 build
```

This creates a production-ready executable in the `build` directory.

## Installation

### Arch Linux

Launchy can be installed on Arch Linux using the provided PKGBUILD:

#### From Release (Recommended)
1. **Download the PKGBUILD:**
   ```bash
   curl -O https://raw.githubusercontent.com/KyleEB/launchy/main/PKGBUILD
   ```

2. **Build and install:**
   ```bash
   makepkg -si
   ```

### Other Linux Distributions


## Usage

1. **Search Applications**: Type in the search bar to find applications by name, description, or category. Use the up and down arrow keys to navigate.
2. **Launch Applications**: Click on any application card to launch it or hit the Enter key

## Development Guidelines

### Go Best Practices
- Follow Clean Architecture principles
- Use interfaces for dependency inversion
- Implement proper error handling with wrapped errors
- Use meaningful variable and function names
- Add comments for exported functions and types

### Svelte Best Practices
- Use TypeScript for type safety
- Follow component composition patterns

### Code Style
- Use consistent formatting (gofmt for Go, Prettier for Svelte)
- Follow naming conventions
- Write self-documenting code
- Keep functions small and focused
- Use meaningful commit messages

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes following the development guidelines
4. Test your changes thoroughly
5. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- Built with [Wails3](https://wails.io/) for cross-platform desktop applications
- Frontend powered by [Svelte](https://svelte.dev/) for reactive user interfaces
- Designed for [CachyOS](https://cachyos.org/) and other Linux distributions
