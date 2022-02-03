# Build

Execute [`go generate`](https://go.dev/blog/generate) and build with the `release` [tag](https://www.digitalocean.com/community/tutorials/customizing-go-binaries-with-build-tags) for a performance optimized build.

```sh
cd cmd/carapace && go generate ./... && go build -ldflags="-s -w" -tags release
```

Completers can also be built separately:

```sh
cd completers/ln_completer && go build -ldflags="-s -w"
./ln_completer _carapace [bash|elvish|fish|nushell|oil|powershell|tcsh|xonsh|zsh]
```
