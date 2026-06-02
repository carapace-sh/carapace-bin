---
name: carapace-setup
description: >
  Use when user needs to set up, configure, or troubleshoot carapace shell completion —
  including per-shell integration, environment variables, overlays, and extensions.
  For choices/variants/bridges, use the carapace-choice skill instead.
  Triggers on: "setup carapace", "configure carapace", "carapace setup", "shell completion setup",
  "add completion to shell", "carapace env vars", "carapace overlay", "carapace extension".
user-invocable: true
---

# Carapace Setup & Configuration

## Install

Download from [releases](https://github.com/carapace-sh/carapace-bin/releases) and add `carapace` to PATH.

| Method | Command |
|--------|---------|
| Homebrew | `brew install carapace` |
| AUR | `pamac install carapace-bin` |
| Mise | `mise use -g carapace@latest` |
| Nix | `nix-shell -p carapace` |
| Scoop | `scoop install extras/carapace-bin` |
| Winget | `winget install -e --id rsteube.Carapace` |
| Termux | `pkg install carapace` |
| DEB | `apt-get update && apt-get install carapace-bin` (from [fury.io](https://rsteube.fury.site/)) |
| RPM | `yum install carapace-bin` (from [fury.io](https://rsteube.fury.site/)) |

## Shell Integration

The config directory is `${UserConfigDir}/carapace/`:

| OS | Path |
|----|------|
| Linux | `~/.config/carapace/` |
| macOS | `~/Library/Application Support/carapace/` |
| Windows | `%APPDATA%\carapace\` |

> On **all** OSes, carapace respects `$XDG_CONFIG_HOME` for its config directory.

The `CARAPACE_BRIDGES` line is optional — see the **carapace-choice** skill for bridge details.

### Bash

```sh
# ~/.bashrc
export CARAPACE_BRIDGES='zsh,fish,bash,inshellisense' # optional
source <(carapace _carapace)
```

### Zsh

```sh
# ${UserConfigDir}/zsh/.zshrc
autoload -U compinit && compinit
export CARAPACE_BRIDGES='zsh,fish,bash,inshellisense' # optional
zstyle ':completion:*' format $'\e[2;37mCompleting %d\e[m'
source <(carapace _carapace)
```

Group order: `zstyle ':completion:*:git:*' group-order 'main commands' 'alias commands' 'external commands'`

### Fish

```sh
# ${UserConfigDir}/fish/config.fish
set -Ux CARAPACE_BRIDGES 'zsh,fish,bash,inshellisense' # optional
carapace _carapace | source
```

### Elvish

```sh
# ${UserConfigDir}/elvish/rc.elv
set-env CARAPACE_BRIDGES 'zsh,fish,bash,inshellisense' # optional
eval (carapace _carapace|slurp)
```

### Nushell

```sh
# ${UserConfigDir}/nushell/env.nu
$env.CARAPACE_BRIDGES = 'zsh,fish,bash,inshellisense' # optional
mkdir $"($nu.cache-dir)"
carapace _carapace nushell | save --force $"($nu.cache-dir)/carapace.nu"

# ${UserConfigDir}/nushell/config.nu
source $"($nu.cache-dir)/carapace.nu"
```

### Powershell

```powershell
# ${UserConfigDir}/powershell/Microsoft.PowerShell_profile.ps1
$env:CARAPACE_BRIDGES = 'zsh,fish,bash,inshellisense' # optional
Set-PSReadLineOption -Colors @{ "Selection" = "`e[7m" }
Set-PSReadlineKeyHandler -Key Tab -Function MenuComplete
carapace _carapace | Out-String | Invoke-Expression
```

> `Set-PSReadlineKeyHandler -Key Tab -Function MenuComplete` is **required**. Without it, raw ANSI escape codes appear. If using `Set-PSReadLineOption -EditMode Emacs`, place it **before** the key handler.

### Oil

```sh
# ${UserConfigDir}/oil/oshrc
export CARAPACE_BRIDGES='zsh,fish,bash,inshellisense' # optional
source <(carapace _carapace)
```

### Xonsh

```sh
# ${UserConfigDir}/xonsh/rc.xsh
$CARAPACE_BRIDGES='zsh,fish,bash,inshellisense' # optional
$COMPLETIONS_CONFIRM=True
exec($(carapace _carapace))
```

### Tcsh

```sh
# ~/.tcshrc
setenv CARAPACE_BRIDGES 'zsh,fish,bash,inshellisense' # optional
set autolist
eval `carapace _carapace`
```

### Cmd (via clink)

```lua
-- ~/AppData/Local/clink/carapace.lua
load(io.popen('carapace _carapace cmd-clink'):read("*a"))()
```

Requires [clink](https://chrisant996.github.io/clink/).

## Environment Variables

Set all variables **before** sourcing the carapace init snippet.

### CARAPACE_BRIDGES

Comma-separated implicit bridge sources. Enables completion for commands carapace lacks a native completer for, by delegating to the shell's own completion. Default order: `zsh,fish,bash,inshellisense`. See the **carapace-choice** skill for full bridge details.

### CARAPACE_EXCLUDES

Comma-separated completer names to exclude: `export CARAPACE_EXCLUDES='git,docker'`

### CARAPACE_COLOR

Whether to output color (default: enabled).

### CARAPACE_DESCRIPTION_LENGTH

Maximum description length in completions.

### CARAPACE_ENV

Register `get-env`/`set-env`/`unset-env` shell functions for env var completion. `0` = disabled, `1` = enabled (default).

### CARAPACE_HIDDEN

Show hidden commands/flags. `0` = disabled (default), `1` = enabled, `2` = enabled including `_carapace`.

### CARAPACE_LENIENT

Allow unknown flags. `0` = disabled (default), `1` = enabled. Overlays implicitly set this.

### CARAPACE_LOG

Enable logging. `0` = disabled (default), `1` = enabled.

### CARAPACE_MATCH

Case-insensitive matching. `0` = case sensitive (default), `1` = case insensitive.

### CARAPACE_MERGEFLAGS

Merge flags to single tag group. `0` = disabled, `1` = enabled (enabled by default in Zsh).

### CARAPACE_NOSPACE

Suffixes that prevent space after completion. `*` = match all (no space after any value).

### CARAPACE_TOOLTIP

Enable tooltip style. `0` = disabled (default), `1` = enabled. Only affects Powershell.

### CARAPACE_UNFILTERED

Skip final filtering step. Enables fuzzy completion in Fish, but only works for (mostly) static values. `0` = disabled (default), `1` = enabled.

## Overlays

YAML files in `${UserConfigDir}/carapace/overlays/` are merged on top of an existing completer at runtime, adding flags/subcommands without replacing it:

```yaml
# ${UserConfigDir}/carapace/overlays/doctl.yaml
name: doctl
persistentflags:
  --output=: Desired output format [text|json]
completion:
  flag:
    output: [text, json]
```

Overlays implicitly set `CARAPACE_LENIENT`.

## Extensions

Extensions are standalone completer binaries that carapace-bin discovers automatically when on `PATH`. They use the carapace bridge protocol to provide richer completion than the built-in completer.

### carapace-aws

The only extension currently: [carapace-sh/carapace-aws](https://github.com/carapace-sh/carapace-aws).

Provides enriched AWS CLI completion by parsing [botocore](https://github.com/boto/botocore) service definitions (descriptions, colored highlighting, custom completions). Falls back to the official `aws_completer` for undefined completions.

**Setup**: Install `carapace-aws` and ensure it's on `PATH`. The built-in `aws` completer in carapace-bin automatically delegates to `carapace-aws` when detected:

```go
// completers/common/aws_completer/cmd/root.go
if _, err := exec.LookPath("carapace-aws"); err == nil {
    return bridge.ActionCarapace("carapace-aws")
}
return bridge.ActionAws("aws")
```

No configuration needed — just install and it takes effect.