# carapace-completers

Currently a playground to test standalone mode of [rsteube/carapace](https://github.com/rsteube/carapace).

## build

```
docker-compose run --rm build
```

for glibc (ubuntu) version or invoke `go build` yourself - for smallest file size (300kb instead of 6mb) use gccgo with flags "-s -w" and upx (upx slows down invocation but should still be fast enough).

## generate completion

``
[xxx_completer] _carapace [bash|elvish|fish|powershell|zsh]
``

(the completer binary needs to be added to PATH)

## completers

- exa (works)
- git (rough experiment)
- gradle (works and has task caching)
