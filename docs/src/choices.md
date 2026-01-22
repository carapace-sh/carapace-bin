# Choices

The **default priority** of completers can be **overridden** with `carapace --choice {name}[/{variant}][@{group}]`.

![](./choices/choices.cast)

> Choices are stored as simple text files in your [config directory](setup/userConfigDir.html).
> ```sh
> carapace
> └── choices
>     ├── sed  # sed@bsd
>     └── tldr # tldr/tldr-python-client
> ```

## Bridges

[Bridges](https://github.com/carapace-sh/carapace-bridge) are a special form of choices as they can be added as **additional completers** otherwise unknown.

![](./choices/bridges.cast)

> With [CARAPACE_BRIDGES](setup/environment.html#carapace_bridges) completers from `zsh`, `fish`, `bash`, and `inshellisense` can largely determined and act as implicit fallback.
> But frameworks like [cobra](setup/environment.html#carapace_bridges) need to be set explicitly.


Currently available bridges:
- [argcomplete](https://github.com/kislyuk/argcomplete)
- [argcomplete_v1](https://github.com/kislyuk/argcomplete) (legacy)
- [aws](https://github.com/aws/aws-cli)
- [bash](https://www.gnu.org/software/bash/)
- [carapace](https://github.com/carapace-sh/carapace)
- [clap](https://github.com/clap-rs/clap) (experimental - needs [clap-rs/clap#3166](https://github.com/clap-rs/clap/issues/3166))
- [click](https://github.com/pallets/click)
- [cobra](https://github.com/spf13/cobra)
- [complete](https://github.com/posener/complete)
- [fish](https://fishshell.com/)
- [inshellisense](https://github.com/microsoft/inshellisense)
- [kingpin](https://github.com/alecthomas/kingpin)
- [kitten](https://github.com/kovidgoyal/kitty)
- [powershell](https://microsoft.com/powershell)
- [urfavecli](https://github.com/urfave/cli)
- [urfavecli_v1](https://github.com/urfave/cli) (legacy)
- [yargs](https://github.com/yargs/yargs)
- [zsh](https://www.zsh.org/)

> Bridging frameworks should be preferred to shells (e.g. `zsh`) as these generally work better and have less overhead.
