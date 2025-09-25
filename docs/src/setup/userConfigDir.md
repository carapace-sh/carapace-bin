# UserConfigDir

Depending on your operating system the [UserConfigDir] can vary.

> - On **Unix** systems, it returns `$XDG_CONFIG_HOME` as specified by [XDG] if non-empty, else `$HOME/.config`.
> - On **Darwin**, it returns `$HOME/Library/Application Support`.
> - On **Windows**, it returns `%AppData%`.
> - On **Plan 9**, it returns `$home/lib`.

> Contrary to the above [Carapace] acknowledges `$XDG_CONFIG_HOME` on **every** operating system for its **config directory**.
> ```sh
> carapace --help
> # Some completions are cached at [/home/rsteube/.cache/carapace].
> # Config is written to [/home/rsteube/.config/carapace].
> # Specs are loaded from [/home/rsteube/.config/carapace/specs].
> ```

[Carapace]:https://carapace.sh
[UserConfigDir]:https://pkg.go.dev/os#UserConfigDir
[XDG]:https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html
