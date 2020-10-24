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
- [Oil](http://www.oilshell.org/) *broken* ([carapace#86](https://github.com/rsteube/carapace/issues/86))
- [Powershell](https://microsoft.com/powershell)
- [Xonsh](https://xon.sh/)
- [Zsh](https://www.zsh.org/)

## Status

A major part of the completers has been generated from help pages so there will be some quirks here and there. Also completion depends on what [rsteube/carapace](https://github.com/rsteube/carapace) is capable of so far.

## Example

```
docker-compose run --rm build
docker-compose run --rm [bash|elvish|fish|oil|powershell|xonsh|zsh]
[ln|mkdir|chown...] <TAB>
```

## Getting Started

Ensure carapace is added to PATH.

```sh
# bash
source <(carapace _carapace)

# elvish
eval (carapace _carapace|slurp)

# fish
carapace _carapace fish | source

# oil
source <(carapace _carapace)

# powershell
Set-PSReadlineKeyHandler -Key Tab -Function MenuComplete
carapace _carapace | Out-String | Invoke-Expression

# xonsh
COMPLETIONS_CONFIRM=True
exec($(carapace _carapace xonsh))

# zsh
source <(carapace _carapace zsh)
```

## Build

```sh
cd cmd/carapace && go generate ./... && go build -ldflags="-s -w"
```

Completers can also be built separately:
```sh
cd completers/ln_completer && go build -ldflags="-s -w"
./ln_completer _carapace [bash|elvish|fish|oil|powershell|xonsh|zsh]
```

## Generate completion

Ensure carapace is added to PATH.

- completion for commands
```sh
carapace [ln|mkdir|...] [bash|elvish|fish|oil|powershell|xonsh|zsh]
```
- list completions
```sh
carapace --list
```
If the shell is left out carapace will try to determine it by the parent process name.

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
cd carapace/cmd && ./generate.sh > completers.go
```
- build & test
```sh
docker-compose run --rm build
docker-compose run --rm [bash|elvish|fish|oil|powershell|xonsh|zsh]
```

[![asciicast](https://asciinema.org/a/357895.svg)](https://asciinema.org/a/357895)
