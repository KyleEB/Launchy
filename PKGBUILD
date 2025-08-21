# Maintainer: Your Name <your.email@example.com>
pkgname=launchy
pkgver=0.1.0
pkgrel=1
pkgdesc="Application Launcher for CachyOS - A modern, fast application launcher built with Wails3 and Svelte"
arch=('x86_64' 'aarch64')
url="https://github.com/yourusername/launchy"
license=('MIT')
depends=('gtk3' 'webkit2gtk-4.1')
makedepends=('go' 'nodejs' 'npm' 'git')
provides=('launchy')
conflicts=('launchy')
source=("$pkgname-$pkgver.tar.gz")
sha256sums=('SKIP')

prepare() {
    cd "$srcdir/$pkgname-$pkgver"
    
    # Install frontend dependencies
    cd frontend
    npm install
    cd ..
    
    # Tidy Go modules
    go mod tidy
}

build() {
    cd "$srcdir/$pkgname-$pkgver"
    
    # Build the application
    wails3 build -config ./build/config.yml
    
    # Generate desktop file
    wails3 generate .desktop \
        -name "Launchy" \
        -exec "launchy" \
        -icon "launchy" \
        -outputfile build/linux/Launchy.desktop \
        -categories "Development;Utility;"
}

package() {
    cd "$srcdir/$pkgname-$pkgver"
    
    # Install binary
    install -Dm755 "bin/Launchy" "$pkgdir/usr/local/bin/launchy"
    
    # Install desktop file
    install -Dm644 "build/linux/Launchy.desktop" "$pkgdir/usr/share/applications/launchy.desktop"
    
    # Install icon (if available)
    if [ -f "build/appicon.png" ]; then
        install -Dm644 "build/appicon.png" "$pkgdir/usr/share/icons/hicolor/128x128/apps/launchy.png"
    fi
    
    # Install license (if available)
    if [ -f "LICENSE" ]; then
        install -Dm644 "LICENSE" "$pkgdir/usr/share/licenses/$pkgname/LICENSE"
    fi
}
