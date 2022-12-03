# Installation

## Manually

Download from [releases](https://github.com/rsteube/carapace-bin/releases) and add `carapace` to [PATH](https://en.wikipedia.org/wiki/PATH_(variable)).

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

Install from [rsteube/homebrew-tap](https://github.com/rsteube/homebrew-tap)

```sh
brew tap rsteube/homebrew-tap
brew install rsteube/tap/carapace
```

## Nix

Install from [nixpkgs](https://search.nixos.org/packages?show=carapace)

```sh
nix-shell -p carapace
```

## RPM

Install from [fury.io](https://rsteube.fury.site/)

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

## Scoop

Install from [rsteube/scoop-bucket](https://github.com/rsteube/scoop-bucket)

```sh
scoop bucket add rsteube https://github.com/rsteube/scoop-bucket.git
scoop install carapace-bin
```
