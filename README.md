# 🚀 Launchy - Application Launcher for CachyOS

A modern, fast application launcher built with Wails3 and Svelte for CachyOS and other Linux distributions.

## Features

- **🔍 Smart Search**: Search applications by name, description, or categories
- **⭐ Favorites**: Mark your most-used applications as favorites for quick access
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
- **FavoritesSection.svelte**: Favorites display section

## Getting Started

### Prerequisites

- Go 1.21 or later
- Node.js 18 or later
- Wails3 CLI

### Installation

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

#### From Source
1. **Clone the repository:**
   ```bash
   git clone https://github.com/KyleEB/launchy.git
   cd launchy
   ```

2. **Install dependencies:**
   ```bash
   sudo pacman -S go nodejs npm git
   go install github.com/wailsapp/wails/v3/cmd/wails3@latest
   ```

3. **Build and install:**
   ```bash
   makepkg -si
   ```

4. **Run Launchy:**
   ```bash
   launchy
   ```

   The application will also be available in your desktop environment's application menu.

### Other Linux Distributions

For other distributions, build from source:

1. **Install dependencies:**
   ```bash
   # Ubuntu/Debian
   sudo apt install golang-go nodejs npm
   
   # Fedora
   sudo dnf install golang nodejs npm
   
   # Arch Linux (if not using PKGBUILD)
   sudo pacman -S go nodejs npm
   ```

2. **Install Wails3 CLI:**
   ```bash
   go install github.com/wailsapp/wails/v3/cmd/wails3@latest
   ```

3. **Build the application:**
   ```bash
   wails3 build
   ```

4. **Run the binary:**
   ```bash
   ./bin/Launchy
   ```

## Usage

1. **Search Applications**: Type in the search bar to find applications by name, description, or category
2. **Launch Applications**: Click on any application card to launch it
3. **Manage Favorites**: Click the star icon on any application to add/remove it from favorites
4. **View Favorites**: Your favorite applications appear at the top of the application list

## Configuration

Favorites are automatically saved to `~/.config/launchy/favorites.json` and persist between sessions.

## Development Guidelines

### Go Best Practices
- Follow Clean Architecture principles
- Use interfaces for dependency inversion
- Implement proper error handling with wrapped errors
- Use meaningful variable and function names
- Add comments for exported functions and types

### Svelte Best Practices
- Use TypeScript for type safety
- Implement proper event handling with `createEventDispatcher`
- Follow component composition patterns
- Use CSS custom properties for theming
- Implement accessibility features (ARIA labels, keyboard navigation)

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
