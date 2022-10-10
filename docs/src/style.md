# Style

[Text transformations](https://elv.sh/ref/builtin.html#styled) can be applied to files and specific values.

> This is only supported in [elvish], [powershell] and [zsh].

## File

Files can be styled with the `LS_COLORS` environment variable (e.g. using [vivid]):

```sh
# elvish
set-env LS_COLORS (vivid generate dracula)

# powershell
$env:LS_COLORS = (vivid generate dracula)

# zsh
export LS_COLORS=$(vivid generate dracula)
```

![](./style-filename.cast)

## Value

Values can be styled with `carapace --style 'key=value'`

```sh
# set
carapace --style 'carapace.Value=bold,magenta'

# clear
carapace --style 'carapace.Description='
```

![](./style-value.cast)


[Elvish]:https://elv.sh/
[Powershell]:https://microsoft.com/powershell
[vivid]:https://github.com/sharkdp/vivid
[Zsh]:https://www.zsh.org/
