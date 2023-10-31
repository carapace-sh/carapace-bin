# Style

[Transformations](#transformations) can be applied to files and specific values.
 
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

Values can be styled with a comma separated list of [transformations](#transformations):

```sh
# set
carapace --style 'carapace.Value=bold,magenta'

# clear
carapace --style 'carapace.Description='
```

![](./style-value.cast)

> Generic configuration like default value and description style can be found under `carapace.{key}`

## Scheme

How the default colors look depends on your [terminal color scheme](https://github.com/rsteube/vincent).

> It is recommended to choose one with a high contrast so that every color is [well readable](https://gogh-co.github.io/Gogh/).

![](./style-scheme.cast)


## Transformations

Transformations are adopted from [elvish](https://elv.sh/ref/builtin.html#styled):

> Each `$style-transformer` can be one of the following:
> - A boolean attribute name:
>   - One of `bold`, `dim`, `italic`, `underlined`, `blink` and `inverse` for
>     setting the corresponding attribute.
>   - An attribute name prefixed by `no-` for unsetting the attribute.
>   - An attribute name prefixed by `toggle-` for toggling the attribute
>     between set and unset.
> - A color name for setting the text color, which may be one of the
>   following:
>   - One of the 8 basic ANSI colors: `black`, `red`, `green`, `yellow`,
>     `blue`, `magenta`, `cyan` and `white`.
>   - The bright variant of the 8 basic ANSI colors, with a `bright-` prefix.
>   - Any color from the xterm 256-color palette, as `colorX` (such as
>     `color12`).
>   - A 24-bit RGB color written as `#RRGGBB` (such as `'#778899'`).
>     **Note**: You need to quote such values, since an unquoted `#` introduces
>     a comment (e.g. use `'bg-#778899'` instead of `bg-#778899`).
> - A color name prefixed by `fg-` to set the foreground color. This has
>   the same effect as specifying the color name without the `fg-` prefix.
> - A color name prefixed by `bg-` to set the background color.


[Elvish]:https://elv.sh/
[Powershell]:https://microsoft.com/powershell
[vivid]:https://github.com/sharkdp/vivid
[Zsh]:https://www.zsh.org/
