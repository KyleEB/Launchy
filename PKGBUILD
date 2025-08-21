# Maintainer: KyleEB
pkgname=launchy
pkgver=0.1.1
pkgrel=1
pkgdesc="Application Launcher for CachyOS - Wails3 + Svelte"
arch=('x86_64' 'aarch64')
url="https://github.com/KyleEB/Launchy"
license=('MIT')
depends=('gtk3' 'webkit2gtk-4.1')
provides=('launchy')
conflicts=('launchy')
source=("Launchy::${url}/releases/download/v${pkgver}/Launchy")
sha256sums=('SKIP')

package() {
  install -Dm755 "Launchy" "$pkgdir/usr/bin/launchy"
}
