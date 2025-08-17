# Launchy - CachyOS App Launcher

A modern, fast, and beautiful application launcher for CachyOS built with Go and Wails.

## Features

- 🚀 **Lightning Fast**: Built with Go for optimal performance
- 🎨 **Modern UI**: Beautiful, responsive interface with dark mode support
- 🔍 **Smart Search**: Intelligent search with relevance ranking
- ⭐ **Favorites**: Mark and quickly access your favorite applications
- 📱 **Responsive Design**: Works perfectly on different screen sizes
- 🏷️ **Categories**: Browse applications by category
- 📊 **Usage Tracking**: Track and display recently used applications
- 🔄 **Auto-refresh**: Automatically detects new applications
- 🌙 **Dark Mode**: Toggle between light and dark themes

## Architecture

Launchy follows clean architecture principles with clear separation of concerns:

```
launchy/
├── internal/
│   ├── domain/           # Business logic and entities
│   │   ├── entities/     # Core domain objects
│   │   ├── repositories/ # Data access interfaces
│   │   └── services/     # Business logic services
│   └── infrastructure/   # External concerns
│       └── repositories/ # Concrete implementations
├── frontend/             # Vue.js frontend
│   ├── src/
│   │   ├── components/   # Vue components
│   │   └── ...
│   └── ...
└── ...
```

## Prerequisites

- Go 1.21 or later
- Node.js 16 or later
- Wails CLI (`go install github.com/wailsapp/wails/v2/cmd/wails@latest`)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/your-username/launchy.git
cd launchy
```

2. Install frontend dependencies:
```bash
cd frontend
npm install
cd ..
```

3. Build the application:
```bash
wails build
```

## Development

### Running in Development Mode

```bash
wails dev
```

This will start the development server with hot reloading for both frontend and backend.

### Building for Production

```bash
wails build
```

This creates a production build for your current platform.

### Cross-platform Building

```bash
# Build for Linux
wails build -platform linux/amd64

# Build for Windows
wails build -platform windows/amd64

# Build for macOS
wails build -platform darwin/amd64
```

## Usage

1. **Search**: Type in the search bar to find applications
2. **Launch**: Click on any application to launch it
3. **Favorites**: Click the star icon to add/remove from favorites
4. **Categories**: Browse applications by category

## Keyboard Shortcuts

- `Ctrl/Cmd + K`: Focus search bar
- `Enter`: Launch selected application
- `Escape`: Clear search
- `Ctrl/Cmd + /`: Toggle dark mode

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Development Guidelines

- Follow clean architecture principles
- Write tests for new features
- Use meaningful commit messages
- Follow Go and Vue.js best practices
- Ensure responsive design works on all screen sizes

## License

This project is licensed under the GNU General Public License v3.0 - see the [LICENSE](LICENSE) file for details.

### License Summary

Launchy is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.

## Acknowledgments

- Built with [Wails](https://wails.io/) - Go framework for building desktop apps
- UI powered by [Vue.js](https://vuejs.org/) and [Tailwind CSS](https://tailwindcss.com/)
- Icons from [Lucide](https://lucide.dev/)
- Inspired by modern application launchers like Spotlight and Alfred

## Support

For support, please open an issue on GitHub or join our community discussions.
