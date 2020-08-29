# carapace-bin

Multi-shell multi-command argument completer based on [rsteube/carapace](https://github.com/rsteube/carapace).

Supported shells:
- [bash](https://www.gnu.org/software/bash/)
- [elvish](https://elv.sh/)
- [fish](https://fishshell.com/)
- [powershell](https://microsoft.com/powershell)
- [zsh](https://www.zsh.org/)


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
