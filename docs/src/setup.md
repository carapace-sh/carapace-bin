# Setup

> This registers all the available [completers](./completers.md).
> It is also possible to load a single one by replacing `_carapace` with the completer name (e.g. `carapace chmod`).

> **See [UserConfigDir](./setup/userConfigDir.md) for details.**

## Bash

```sh
# ~/.bashrc
export CARAPACE_BRIDGES='zsh,fish,bash,inshellisense' # optional
source <(carapace _carapace)
```

![](./setup-bash.png)

## Cmd

```lua
# ~/AppData/Local/clink/carapace.lua
load(io.popen('carapace _carapace cmd-clink'):read("*a"))()
```

![](./setup-cmd.png)

> Needs [clink](https://chrisant996.github.io/clink/).

## Elvish

```sh
# ${UserConfigDir}/elvish/rc.elv
set-env CARAPACE_BRIDGES 'zsh,fish,bash,inshellisense' # optional
eval (carapace _carapace|slurp)
```

![](./setup-elvish.png)

## Fish

```sh
# ${UserConfigDir}/fish/config.fish
set -Ux CARAPACE_BRIDGES 'zsh,fish,bash,inshellisense' # optional
carapace _carapace | source
```

![](./setup-fish.png)

## Nushell

```sh
## ${UserConfigDir}/nushell/env.nu
$env.CARAPACE_BRIDGES = 'zsh,fish,bash,inshellisense' # optional
mkdir $"($nu.cache-dir)"
carapace _carapace nushell | save --force $"($nu.cache-dir)/carapace.nu"

# ${UserConfigDir}/nushell/config.nu
source $"($nu.cache-dir)/carapace.nu"
```

![](./setup-nushell.png)

## Oil

```sh
# ${UserConfigDir}/oil/oshrc
export CARAPACE_BRIDGES='zsh,fish,bash,inshellisense' # optional
source <(carapace _carapace)
```

![](./setup-oil.png)

## Powershell

```sh
# ${UserConfigDir}/powershell/Microsoft.PowerShell_profile.ps1
$env:CARAPACE_BRIDGES = 'zsh,fish,bash,inshellisense' # optional
Set-PSReadLineOption -Colors @{ "Selection" = "`e[7m" }
Set-PSReadlineKeyHandler -Key Tab -Function MenuComplete
carapace _carapace | Out-String | Invoke-Expression
```

![](./setup-powershell.png)

## Tcsh
```sh
# ~/.tcshrc
setenv CARAPACE_BRIDGES 'zsh,fish,bash,inshellisense' # optional
set autolist
eval `carapace _carapace`
```

![](./setup-tcsh.png)

## Xonsh
```sh
# ${UserConfigDir}/xonsh/rc.xsh
$CARAPACE_BRIDGES='zsh,fish,bash,inshellisense' # optional
$COMPLETIONS_CONFIRM=True
exec($(carapace _carapace))
```

![](./setup-xonsh.png)

## Zsh

```sh
# ${UserConfigDir}/zsh/.zshrc
autoload -U compinit && compinit
export CARAPACE_BRIDGES='zsh,fish,bash,inshellisense' # optional
zstyle ':completion:*' format $'\e[2;37mCompleting %d\e[m'
source <(carapace _carapace)
```

Order of groups can be configured with the zstyle [group-order](https://zsh.sourceforge.io/Guide/zshguide06.html).
```sh
zstyle ':completion:*:git:*' group-order 'main commands' 'alias commands' 'external commands'
```

![](./setup-zsh.png)
