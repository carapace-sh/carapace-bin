> unreleased

# Choices

The **default priority** of completers can be **overridden** with `carapace --choice {name}[/{variant}][@{group}]`.

![](./choices/choices.cast)

> Choices are stored as simple text files in your [config directory](http://localhost:3000/setup/userConfigDir.html).
> ```sh
> carapace
> └── choices
>     ├── sed  # sed@bsd
>     └── tldr # tldr/tldr-python-client
> ```

## Bridges

[Bridges](https://github.com/carapace-sh/carapace-bridge) are a special form of choices as they can be added as **additional completers** otherwise unknown.

![](./choices/bridges.cast)

> With [CARAPACE_BRIDGES](http://localhost:3000/setup/environment.html#carapace_bridges) completers from `zsh`, `fish`, `bash`, and `inshellisense` can largely determined and act as implicit fallback.
> But frameworks like [cobra](http://localhost:3000/setup/environment.html#carapace_bridges) need to be set explicitly.
