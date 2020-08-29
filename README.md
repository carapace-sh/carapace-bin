# carapace-bin

Multi-shell multi-command argument completer based on [rsteube/carapace](https://github.com/rsteube/carapace).

Supported shells:
- bash
- elvish
- fish
- powershell
- zsh


## build

```
docker-compose run --rm build
```

for glibc (ubuntu) version or invoke `go build` yourself - for smallest file size (300kb instead of 6mb) use gccgo with flags "-s -w" and upx (upx slows down invocation but should still be fast enough).

## generate completion

``
carapace [ln|mkdir|...] [bash|elvish|fish|powershell|zsh]
``
(the completer binary needs to be added to PATH)
