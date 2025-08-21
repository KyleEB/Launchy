# Maintainer: Your Name <your.email@example.com>
pkgname=launchy
pkgver=0.1.0
pkgrel=1
pkgdesc="Application Launcher for CachyOS - Wails3 + Svelte"
arch=('x86_64' 'aarch64')
url="https://github.com/KyleEB/launchy"
license=('MIT')
depends=('gtk3' 'webkit2gtk-4.1')
makedepends=('go' 'nodejs' 'npm' 'git')
provides=('launchy')
conflicts=('launchy')
source=("$url/releases/download/v$pkgver/launchy-$pkgver.tar.gz")
sha256sums=('SKIP')

# local tool bin for bootstrapping wails3
_toolbin="$srcdir/.bin"

prepare() {
  mkdir -p "$_toolbin"
  export GOBIN="$_toolbin"
  export PATH="$_toolbin:$PATH"

  # Bootstrap Wails v3 CLI locally (pin version if you want reproducibility)
  go install github.com/wailsapp/wails/v3/cmd/wails3@latest

  (cd frontend && npm ci)
  go mod tidy
}

build() {
  export PATH="$_toolbin:$PATH"

  wails3 build -config ./build/config.yml

  # If your CLI doesn't support this, keep a static .desktop in repo instead.
  if command -v wails3 >/dev/null 2>&1; then
    wails3 generate .desktop \
      -name "Launchy" \
      -exec "launchy" \
      -icon "launchy" \
      -outputfile build/linux/Launchy.desktop \
      -categories "Utility;"
  fi
}

package() {
  install -Dm755 "bin/Launchy" "$pkgdir/usr/bin/launchy"
  [ -f "build/linux/Launchy.desktop" ] && install -Dm644 "build/linux/Launchy.desktop" \
                                                         "$pkgdir/usr/share/applications/launchy.desktop"
  [ -f "build/appicon.png" ] && install -Dm644 "build/appicon.png" \
                                               "$pkgdir/usr/share/icons/hicolor/128x128/apps/launchy.png"
  [ -f "LICENSE" ] && install -Dm644 "LICENSE" "$pkgdir/usr/share/licenses/$pkgname/LICENSE"
}
