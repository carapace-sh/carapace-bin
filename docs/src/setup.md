# Setup

Depending on your shell, add the following changes to the corresponding config files:

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
carapace _carapace | save carapace.nu ; nu -c 'source carapace.nu'

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

> This registers all the available [completers](./completers.md). It is also possible to load a single one by replacing `_carapace` with the completer name (e.g. `carapace chmod`).
