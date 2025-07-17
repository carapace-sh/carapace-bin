# Bridges

With `~/.config/carapace/bridges.yaml` generic bridges can be configured for commands.

```yaml
gh: cobra
opentofu: complete
tsh: kingpin
```

![](./bridges.cast)

> `bridges.yaml` is a simpler alternative to [Spec/Bridge](./spec/bridge.md).
> It has a higher precedence than [CARAPACES_BRIDGES](./setup/environment.md#carapace_bridges) but lower than internal completers and specs.
> So disable an internal completer with [CARAPACES_EXCLUDES](./setup/environment.md#carapace_excludes) if you want to replace it with a bridge.

Currently available bridges:
- [argcomplete](https://github.com/kislyuk/argcomplete)
- [argcomplete@v1](https://github.com/kislyuk/argcomplete) (legacy)
- [bash](https://www.gnu.org/software/bash/)
- [carapace](https://github.com/carapace-sh/carapace)
- [clap](https://github.com/clap-rs/clap) (needs [clap-rs/clap#3166](https://github.com/clap-rs/clap/issues/3166))
- [click](https://github.com/pallets/click)
- [cobra](https://github.com/spf13/cobra)
- [complete](https://github.com/posener/complete)
- [fish](https://fishshell.com/)
- [inshellisense](https://github.com/microsoft/inshellisense)
- [kingpin](https://github.com/alecthomas/kingpin)
- [powershell](https://microsoft.com/powershell)
- [urfavecli](https://github.com/urfave/cli)
- [urfavecli@v1](https://github.com/urfave/cli) (legacy)
- [yargs](https://github.com/yargs/yargs)
- [zsh](https://www.zsh.org/)

> Bridging frameworks should be preferred to shells (e.g. zsh) as these generally work better and have less overhead.

## Cache

The list of completers is cached to not impact shell startup time.
Clear the cache with `carapace --clear-cache` if your system changes.
