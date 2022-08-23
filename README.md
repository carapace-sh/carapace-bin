# carapace-bin

[![PkgGoDev](https://pkg.go.dev/badge/github.com/rsteube/carapace-bin/pkg/actions)](https://pkg.go.dev/github.com/rsteube/carapace-bin/pkg/actions)
[![GoReportCard](https://goreportcard.com/badge/github.com/rsteube/carapace-bin)](https://goreportcard.com/report/github.com/rsteube/carapace-bin)
[![documentation](https://img.shields.io/badge/&zwnj;-documentation-blue?logo=gitbook)](https://rsteube.github.io/carapace-bin/)
[![Completers](https://rsteube.github.io/carapace-bin/badge.svg)](https://rsteube.github.io/carapace-bin/completers.html)
[![Macros](https://rsteube.github.io/carapace-bin/macros-badge.svg)](https://rsteube.github.io/carapace-bin/spec/macros.html)

Carapace-bin provides argument completions for many CLI commands: [see the full list here](https://rsteube.github.io/carapace-bin/completers.html), and works across many POSIX and non-POSIX shells.
This multi-shell multi-command argument completer is based on [rsteube/carapace](https://github.com/rsteube/carapace).
You can read more about this tool here: _[A pragmatic approach to shell completion](https://dev.to/rsteube/a-pragmatic-approach-to-shell-completion-4gp0)_.

[![asciicast](https://asciinema.org/a/357191.svg)](https://asciinema.org/a/357191)

Supported shells:
- [Bash](https://www.gnu.org/software/bash/)
- [Elvish](https://elv.sh/)
- [Fish](https://fishshell.com/)
- [Ion](https://doc.redox-os.org/ion-manual/) ([experimental](https://github.com/rsteube/carapace/issues/88))
- [Nushell](https://www.nushell.sh/)
- [Oil](http://www.oilshell.org/)
- [Powershell](https://microsoft.com/powershell)
- [Tcsh](https://www.tcsh.org/) ([experimental](https://github.com/rsteube/carapace/issues/331))
- [Xonsh](https://xon.sh/)
- [Zsh](https://www.zsh.org/)

## Status

A major part of the completers has been generated from help pages so there will be some quirks here and there. Also completion depends on what [rsteube/carapace](https://github.com/rsteube/carapace) is capable of so far.

## Getting Started

Ensure carapace is added to [PATH](https://en.wikipedia.org/wiki/PATH_(variable)) ([Installation](https://rsteube.github.io/carapace-bin/installation.html)).
Then register the [completers](https://rsteube.github.io/carapace-bin/completers.html) ([Setup](https://rsteube.github.io/carapace-bin/setup.html)):

```sh
# bash (~/.bashrc)
source <(carapace _carapace)

# elvish (~/.elvish/rc.elv)
eval (carapace _carapace|slurp)

# fish (~/.config/fish/config.fish)
mkdir -p ~/.config/fish/completions
carapace --list | awk '{print $1}' | xargs -I{} touch ~/.config/fish/completions/{}.fish # disable auto-loaded completions (#185)
carapace _carapace | source

# nushell (~/.config/nushell/config.nu)
carapace _carapace nushell # update config.nu manually according to output

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
