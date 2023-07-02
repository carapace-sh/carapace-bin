# Setup

> This registers all the available [completers](./completers.md).
> It is also possible to load a single one by replacing `_carapace` with the completer name (e.g. `carapace chmod`).
>
> Skipping a completer is not possible [yet](https://github.com/rsteube/carapace/pull/539), but registering a different one after carapace overrides it.
> The script can also be edited manually to remove unwanted entries.

## Bash

```sh
# ~/.bashrc
source <(carapace _carapace)
```

![](./setup-bash.png)

## Elvish

```sh
# ~/.elvish/rc.elv
eval (carapace _carapace|slurp)
```

![](./setup-elvish.png)

## Fish

```sh
# ~/.config/fish/config.fish
mkdir -p ~/.config/fish/completions
carapace --list | awk '{print $1}' | xargs -I{} touch ~/.config/fish/completions/{}.fish # disable auto-loaded completions (#185)
carapace _carapace | source
```

![](./setup-fish.png)

## Nushell

```sh
## ~/.config/nushell/env.nu
mkdir ~/.cache/carapace
carapace _carapace nushell | save --force ~/.cache/carapace/init.nu

#~/.config/nushell/config.nu
source ~/.cache/carapace/init.nu
```

![](./setup-nushell.png)

## Oil

```sh
# ~/.config/oil/oshrc
source <(carapace _carapace)
```

![](./setup-oil.png)

## Powershell

```sh
# ~/.config/powershell/Microsoft.PowerShell_profile.ps1
Set-PSReadLineOption -Colors @{ "Selection" = "`e[7m" }
Set-PSReadlineKeyHandler -Key Tab -Function MenuComplete
carapace _carapace | Out-String | Invoke-Expression
```

![](./setup-powershell.png)

## Tcsh
```sh
# ~/.tcshrc
set autolist
eval `carapace _carapace`
```

![](./setup-tcsh.png)

## Xonsh
```sh
# ~/.config/xonsh/rc.xsh
COMPLETIONS_CONFIRM=True
exec($(carapace _carapace))
```

![](./setup-xonsh.png)

## Zsh

```sh
# ~/.zshrc
zstyle ':completion:*' format $'\e[2;37mCompleting %d\e[m'
source <(carapace _carapace)
```

Order of groups can be configured with the zstyle [group-order](https://zsh.sourceforge.io/Guide/zshguide06.html).
```sh
zstyle ':completion:*:git:*' group-order 'main commands' 'alias commands' 'external commands'
```

![](./setup-zsh.png)
