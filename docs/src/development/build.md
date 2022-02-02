# Build

```sh
cd cmd/carapace && go generate ./... && go build -ldflags="-s -w" -tags release
```

Completers can also be built separately:
```sh
cd completers/ln_completer && go build -ldflags="-s -w"
./ln_completer _carapace [bash|elvish|fish|nushell|oil|powershell|tcsh|xonsh|zsh]
```
