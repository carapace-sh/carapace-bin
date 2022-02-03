# Installation

## Manually

Download from [releases](https://github.com/rsteube/carapace-bin/releases) and add `carapace` to [PATH](https://en.wikipedia.org/wiki/PATH_(variable)).

## AUR

Install [carapace-bin](https://aur.archlinux.org/packages/carapace-bin/) from [AUR](https://aur.archlinux.org/).

```sh
# e.g. with pamac
pamac install carapace-bin
```

## Homebrew

Install from [rsteube/homebrew-tap](https://github.com/rsteube/homebrew-tap)

```sh
brew tap rsteube/homebrew-tap
brew install rsteube/tap/carapace
```

## Scoop

Install from [rsteube/scoop-bucket](https://github.com/rsteube/scoop-bucket)

```sh
scoop bucket add rsteube https://github.com/rsteube/scoop-bucket.git
scoop install carapace-bin
```
# Setup

Depending on your shell, use the following changes to the required config files:

```sh
# bash (~/.bashrc)
source <(carapace _carapace)

# elvish (~/.elvish/rc.elv)
eval (carapace _carapace|slurp)

# fish (~/.config/fish/config.fish)
mkdir -p ~/.config/fish/completions
carapace --list | awk '{print $1}' | xargs -I{} touch ~/.config/fish/completions/{}.fish # disable auto-loaded completions (#185)
carapace _carapace | source

# nushell [needs fork: https://github.com/rsteube/nushell]
carapace _carapace nushell | save carapace.nu ; nu -c 'source carapace.nu'

# oil (~/.config/oil/oshrc)
source <(carapace _carapace)

# powershell (~/.config/powershell/Microsoft.PowerShell_profile.ps1)
Set-PSReadlineKeyHandler -Key Tab -Function MenuComplete
carapace _carapace | Out-String | Invoke-Expression

# tcsh (~/.tcshrc)
set autolist
eval `carapace _carapace`

# xonsh (~/.config/xonsh/rc.xsh)
COMPLETIONS_CONFIRM=True
exec($(carapace _carapace))

# zsh (~/.zshrc)
source <(carapace _carapace)
```
