# Creating completers


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
docker-compose run --rm [bash|elvish|fish|ion|nushell|oil|powershell|tcsh|xonsh|zsh]
```

[![asciicast](https://asciinema.org/a/357895.svg)](https://asciinema.org/a/357895)
