# Style

Style is only supported in [elvish], [powershell] and [zsh].

## Filename
Filenames can be configured with the `LS_COLORS` environment variable - e.g with [vivid]:
```sh
# elvish
set E:LS_COLORS = (vivid generate dracula)

# zsh
export LS_COLORS="$(vivid generate dracula)"
```

![](./ls_colors.png)

# Value

Values can be configured with `carapace --style 'key=value'`

```sh
# set
carapace --style 'carapace.Value=bold,magenta'

# clear
carapace --style 'carapace.Description='
```

![](./style-config.png)


[Elvish]:https://elv.sh/
[Powershell]:https://microsoft.com/powershell
[vivid]:https://github.com/sharkdp/vivid
[Zsh]:https://www.zsh.org/
