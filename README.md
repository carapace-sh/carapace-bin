# carapace-bin

[![CircleCI](https://circleci.com/gh/rsteube/carapace-bin.svg?style=svg)](https://circleci.com/gh/rsteube/carapace-bin)

Multi-shell multi-command argument completer based on [rsteube/carapace](https://github.com/rsteube/carapace).

[![asciicast](https://asciinema.org/a/357191.svg)](https://asciinema.org/a/357191)

Supported shells:
- [Bash](https://www.gnu.org/software/bash/)
- [Elvish](https://elv.sh/)
- [Fish](https://fishshell.com/)
- [Powershell](https://microsoft.com/powershell)
- [Zsh](https://www.zsh.org/)

## Status

A major part of the completers has been generated from help pages so there will be some quirks here and there. Also completion depends on what [rsteube/carapace](https://github.com/rsteube/carapace) is capable of so far which by itself is limited by [spf13/pflag](https://github.com/spf13/pflag/) (e.g. shorthand-only flags are currently only possible due to [PR 256](https://github.com/spf13/pflag/pull/256)).

Notable shortcomings:
- no support for long shorthands (e.g. `-ldflags`)
- no support for combined shorthands (e.g. `tar -xzvf`)

## Example

```
docker-compose run --rm build
docker-compose run --rm [bash|elvish|fish|powershell|zsh]
[ln|mkdir|chown...] <TAB>
```

## Build

```sh
cd carapace && go build -ldflags="-s -w"
```
For smallest file size (300kb instead of 6mb) use gccgo with flags "-s -w" and upx (upx slows down invocation but should still be fast enough).

Completers can also be built separately:
```sh
cd completers/ln_completer && go build -ldflags="-s -w"
./ln_completer _carapace [bash|elvish|fish|powershell|zsh]
```

## Generate completion

Ensure carapace is added to PATH.

- completion for carapace itself
```sh
carapace _carapace [bash|elvish|fish|powershell|zsh]
```
- completion for commands
```sh
carapace [ln|mkdir|...] [bash|elvish|fish|powershell|zsh]
```
- list completions
```sh
carapace --list
```
If the shell is left out carapace will try to determine it by the parent process name.

## Creating completers
[caraparse](/caraparse) is a helper tool that uses regex to parse gnu help pages.
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
docker-compose run --rm [bash|elvish|fish|powershell|zsh]
```
