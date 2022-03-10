# Bridge

With `--bridge` completions from other engines can be bridged for shells not natively supported by it.

```sh
carapace --bridge [command]/[argcomplete|cobra|posener]
```

Currently supported are:
- [fish-shell/fish-shell](https://github.com/fish-shell/fish-shell)
- [kislyuk/argcomplete](https://github.com/kislyuk/argcomplete)
- [posener/complete](https://github.com/posener/complete)
- [spf13/cobra](https://github.com/spf13/cobra)

E.g. for [vault](https://www.vaultproject.io/):
```sh
# bash (~/.bashrc)
source <(carapace --bridge vault/posener)

# elvish (~/.elvish/rc.elv)
eval (carapace --bridge vault/posener|slurp)

# fish (~/.config/fish/config.fish)
carapace --bridge vault/posener | source

# oil (~/.config/oil/oshrc)
source <(carapace --bridge vault/posener)

# powershell (~/.config/powershell/Microsoft.PowerShell_profile.ps1)
carapace --bridge vault/posener | Out-String | Invoke-Expression

# tcsh (~/.tcshrc)
eval `carapace --bridge vault/posener`

# xonsh (~/.config/xonsh/rc.xsh)
exec($(carapace --bridge vault/posener))

# zsh (~/.zshrc)
source <(carapace --bridge vault/posener)
```
