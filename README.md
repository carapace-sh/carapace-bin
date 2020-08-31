# carapace-bin

Multi-shell multi-command argument completer based on [rsteube/carapace](https://github.com/rsteube/carapace).

Supported shells:
- [Bash](https://www.gnu.org/software/bash/)
- [Elvish](https://elv.sh/)
- [Fish](https://fishshell.com/)
- [Powershell](https://microsoft.com/powershell)
- [Zsh](https://www.zsh.org/)


## Example

```
docker-compose run --rm build
docker-compose run --rm [bash|elvish|fish|powershell|zsh]
[ln|mkdir|chown...] <TAB>
```

for smallest file size (300kb instead of 6mb) use gccgo with flags "-s -w" and upx (upx slows down invocation but should still be fast enough).

## generate completion

``
carapace [ln|mkdir|...] [bash|elvish|fish|powershell|zsh]
``
(the completer binary needs to be added to PATH)

# Creating completers
[caraparse](https://github.com/rsteube/carapace/tree/support-shorthand-only-flags/caraparse) is a helper tool that uses regex to parse gnu help pages.
Due to strong inconsistencies between these the results may differ, but generally give a good head start.

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
		carapace.ActionDirectories(),
	)
```
- run the generator
```sh
cd cmd && ./generate.sh > completers.go
```
- build & test
```sh
docker-compose run --rm build
docker-compose run --rm [bash|elvish|fish|powershell|zsh]
```
