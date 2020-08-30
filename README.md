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
