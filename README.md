# carapace-bin

[![CircleCI](https://circleci.com/gh/rsteube/carapace-bin.svg?style=svg)](https://circleci.com/gh/rsteube/carapace-bin)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/rsteube/carapace-bin/pkg/actions)](https://pkg.go.dev/github.com/rsteube/carapace-bin/pkg/actions)
[![GoReportCard](https://goreportcard.com/badge/github.com/rsteube/carapace-bin)](https://goreportcard.com/report/github.com/rsteube/carapace-bin)

Multi-shell multi-command argument completer based on [rsteube/carapace](https://github.com/rsteube/carapace).

[![asciicast](https://asciinema.org/a/357191.svg)](https://asciinema.org/a/357191)

Supported shells:
- [Bash](https://www.gnu.org/software/bash/)
- [Elvish](https://elv.sh/)
- [Fish](https://fishshell.com/)
- [Ion](https://doc.redox-os.org/ion-manual/html/) ([experimental](https://github.com/rsteube/carapace/issues/88))
- [Nushell](https://www.nushell.sh/) ([experimental](https://github.com/rsteube/carapace/issues/89))
- [Oil](http://www.oilshell.org/)
- [Powershell](https://microsoft.com/powershell)
- [Tcsh](https://www.tcsh.org/) ([experimental](https://github.com/rsteube/carapace/issues/331))
- [Xonsh](https://xon.sh/)
- [Zsh](https://www.zsh.org/)

## Status

A major part of the completers has been generated from help pages so there will be some quirks here and there. Also completion depends on what [rsteube/carapace](https://github.com/rsteube/carapace) is capable of so far.

## Example

```
docker-compose run --rm build
docker-compose run --rm [bash|elvish|fish|ion|nushell|oil|powershell|tcsh|xonsh|zsh]
[ln|mkdir|chown...] <TAB>
```

## Getting Started

Ensure carapace is added to PATH.

- completion for commands (replace `_carapace` with completer name to load a single one)
```sh
# bash (~/.bashrc)
source <(carapace _carapace)

# elvish (~/.elvish/rc.elv)
eval (carapace _carapace|slurp)

# fish (~/.config/fish/config.fish)
mkdir -p ~/.config/fish/completions
carapace --list | awk '{print $1}' | xargs -I{} touch ~/.config/fish/completions/{}.fish # disable auto-loaded completions (#185)
carapace _carapace | source

# oil (~/.config/oil/oshrc)
source <(carapace _carapace)

# powershell (~/.config/powershell/Microsoft.PowerShell_profile.ps1)
Set-PSReadlineKeyHandler -Key Tab -Function MenuComplete
carapace _carapace | Out-String | Invoke-Expression

# tcsh
set autolist
eval `carapace _carapace`

# xonsh (~/.config/xonsh/rc.xsh)
COMPLETIONS_CONFIRM=True
exec($(carapace _carapace))

# zsh (~/.zshrc)
source <(carapace _carapace)
```

- list completions
```sh
carapace --list
```

## Build

```sh
cd cmd/carapace && go generate ./... && go build -ldflags="-s -w"
```

Completers can also be built separately:
```sh
cd completers/ln_completer && go build -ldflags="-s -w"
./ln_completer _carapace [bash|elvish|fish|oil|powershell|tcsh|xonsh|zsh]
```

## Creating completers
[caraparse](/cmd/caraparse) is a helper tool that uses regex to parse gnu help pages.
Due to strong inconsistencies between these the results may differ but generally give a good head start.

- copy a completer for simplicity
```sh
cp -r completers/cp_completer completers/ln_completer
```
- update the package name in `main.go`
- replace `root.go`
```sh
ln --help | caraparse -n ln > completers/ln_completer/cmd/root.go
```
- fix issues and add completions as required
```go
	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"backup":           carapace.ActionValues("existing", "nil", "none", "off", "numbered", "t", "simple", "never"),
		"target-directory": carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(""),
	)
```
- run the generator
```sh
go generate ./...
```
- build & test
```sh
docker-compose run --rm build
docker-compose run --rm [bash|elvish|fish|oil|powershell|tcsh|xonsh|zsh]
```

[![asciicast](https://asciinema.org/a/357895.svg)](https://asciinema.org/a/357895)
