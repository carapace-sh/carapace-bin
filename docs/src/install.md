# Install

## Manually

Download from [releases](https://github.com/carapace-sh/carapace-bin/releases) and add `carapace` to [PATH](https://en.wikipedia.org/wiki/PATH_(variable)).

## AUR

Install [carapace-bin](https://aur.archlinux.org/packages/carapace-bin/) from [AUR](https://aur.archlinux.org/).

```sh
# e.g. with pamac
pamac install carapace-bin
```

## DEB

Install from [fury.io](https://rsteube.fury.site/)

```toml
# /etc/apt/sources.list.d/fury.list
deb [trusted=yes] https://apt.fury.io/rsteube/ /
```

```sh
apt-get update && apt-get install carapace-bin
```

## Homebrew

Install from [homebrew-core](https://formulae.brew.sh/formula/carapace)

```sh
brew install carapace
```

---

Install from [rsteube/homebrew-tap](https://github.com/rsteube/homebrew-tap)

```sh
brew tap rsteube/homebrew-tap
brew install rsteube/tap/carapace
```

## Mise [mise](https://github.com/jdx/mise)

```sh
mise use -g carapace@latest
```

## Nix

Install from [nixpkgs](https://search.nixos.org/packages?show=carapace)

```sh
nix-shell -p carapace
```

## PKGX

Install from [pkgx.dev](https://pkgx.dev/pkgs/carapace.sh/)

```sh
pkgx install carapace
```

## RPM

Install from [fury.io](https://rsteube.fury.site/)

### Yum

```toml
# /etc/yum.repos.d/fury.repo
[fury]
name=Gemfury Private Repo
baseurl=https://yum.fury.io/rsteube/
enabled=1
gpgcheck=0
```

```sh
yum install carapace-bin
```

### Zypper

```sh
zypper ar --gpgcheck-allow-unsigned -f https://yum.fury.io/rsteube/ carapace
zypper install carapace-bin
```

## Scoop

Install from [ScoopInstaller/Extras](https://github.com/ScoopInstaller/Extras) (**recommended**)

```sh
scoop bucket add extras
scoop install extras/carapace-bin
```

---

Install from [rsteube/scoop-bucket](https://github.com/rsteube/scoop-bucket)

```sh
scoop bucket add rsteube https://github.com/rsteube/scoop-bucket.git
scoop install carapace-bin
```

## Termux

Install from [termux/termux-packages](https://github.com/termux/termux-packages)

```sh
pkg install carapace
```

---

Install from [carapace-sh/termux](https://github.com/carapace-sh/termux) (gh_pages)
> **WIP**: repo currently manually created

### Manually

```sh
# $PREFIX/etc/apt/sources.list.d
deb [trusted=yes] https://termux.carapace.sh termux extras  
```

```sh
apt update && apt install carapace-bin
```

### Script
```sh
curl termux.carapace.sh | sh
```

## Winget

Install from [winget-pkgs](https://github.com/microsoft/winget-pkgs)

```sh
winget install -e --id rsteube.Carapace
```

## X-CMD

Install from [x-cmd.com](https://www.x-cmd.com/pkg/carapace-bin)

```sh
x env use carapace-bin
```
