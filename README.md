# carapace-bin

[![PkgGoDev](https://pkg.go.dev/badge/github.com/carapace-sh/carapace-bin/pkg/actions)](https://pkg.go.dev/github.com/carapace-sh/carapace-bin/pkg/actions)
[![GoReportCard](https://goreportcard.com/badge/github.com/carapace-sh/carapace-bin)](https://goreportcard.com/report/github.com/carapace-sh/carapace-bin)
[![documentation](https://img.shields.io/badge/&zwnj;-documentation-blue?logo=gitbook)](https://carapace-sh.github.io/carapace-bin/)
[![Completers](https://carapace-sh.github.io/carapace-bin/badge.svg)](https://carapace-sh.github.io/carapace-bin/completers.html)
[![Macros](https://carapace-sh.github.io/carapace-bin/macros-badge.svg)](https://carapace-sh.github.io/carapace-bin/spec/macros.html)
[![Packaging status](https://repology.org/badge/tiny-repos/carapace.svg)](https://repology.org/project/carapace/versions)
[![faq](https://img.shields.io/badge/discussions-faq-white)](https://github.com/carapace-sh/carapace-bin/discussions/categories/q-a?discussions_q=category%3AQ%26A+label%3Afaq)

Carapace-bin provides argument completion for multiple CLI commands ([full list](https://carapace-sh.github.io/carapace-bin/completers.html)), and works across multiple POSIX and non-POSIX shells.
You can read more about it here: _[A pragmatic approach to shell completion](https://dev.to/rsteube/a-pragmatic-approach-to-shell-completion-4gp0)_.

[![asciicast](https://asciinema.org/a/533283.svg)](https://asciinema.org/a/533283)

Supported shells:
- [Bash](https://www.gnu.org/software/bash/)
- [Elvish](https://elv.sh/)
- [Fish](https://fishshell.com/)
- [Ion](https://doc.redox-os.org/ion-manual/) ([experimental](https://github.com/carapace-sh/carapace/issues/88))
- [Nushell](https://www.nushell.sh/)
- [Oil](http://www.oilshell.org/)
- [Powershell](https://microsoft.com/powershell)
- [Tcsh](https://www.tcsh.org/) ([experimental](https://github.com/carapace-sh/carapace/issues/331))
- [Xonsh](https://xon.sh/)
- [Zsh](https://www.zsh.org/)

## Getting Started

Ensure carapace is added to [PATH](https://en.wikipedia.org/wiki/PATH_(variable)) ([Installation](https://carapace-sh.github.io/carapace-bin/installation.html)).
Then register the [completers](https://carapace-sh.github.io/carapace-bin/completers.html) ([Setup](https://carapace-sh.github.io/carapace-bin/setup.html)):

```sh
# bash (~/.bashrc)
source <(carapace _carapace)

# elvish (~/.elvish/rc.elv)
eval (carapace _carapace|slurp)

# fish (~/.config/fish/config.fish)
mkdir -p ~/.config/fish/completions
carapace --list | awk '{print $1}' | xargs -I{} touch ~/.config/fish/completions/{}.fish # disable auto-loaded completions (#185)
carapace _carapace | source

# nushell

## ~/.config/nushell/env.nu
mkdir ~/.cache/carapace
carapace _carapace nushell | save --force ~/.cache/carapace/init.nu

## ~/.config/nushell/config.nu
source ~/.cache/carapace/init.nu

# oil (~/.config/oil/oshrc)
source <(carapace _carapace)

# powershell (~/.config/powershell/Microsoft.PowerShell_profile.ps1)
Set-PSReadLineOption -Colors @{ "Selection" = "`e[7m" }
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
