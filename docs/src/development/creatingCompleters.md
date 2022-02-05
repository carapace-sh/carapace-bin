# Creating completers


- copy a basic completer for simplicity
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
- run the generator once to add the new completer
```sh
# cmd/carapace
go generate ./...
go install
```

## Example

[![asciicast](https://asciinema.org/a/466859.svg)](https://asciinema.org/a/466859)
