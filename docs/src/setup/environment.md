# Environment

## CARAPACE_BRIDGES

Which implicit [bridges](https://github.com/carapace-sh/carapace-bridge) to enable.

- [bash](https://www.gnu.org/software/bash/)
- [fish](https://fishshell.com/)
- [inshellisense](https://github.com/microsoft/inshellisense)
- [zsh](https://www.zsh.org/)

![](./bridges.cast)

> The list of completers is cached to not impact shell startup time.
> Clear the cache with `carapace --clear-cache` if your system changes.

## CARAPACE_COLOR

Whether to output color (default).

![](./color.cast)

## CARAPACE_COVERDIR

Coverage directory for sandbox tests (internal).
      
## CARAPACE_ENV

Whether to register `get-env`, `set-env` and `unset-env` functions.

- `0` - disabled
- `1` - enabled

![](./env.cast)

## CARAPACE_EXCLUDES

Which completers to exclude.

![](./excludes.cast)
          
## CARAPACE_HIDDEN

Whether to show hidden commands/flags.
        
- `0` - disabled
- `1` - enabled
- `2` - enabled including `_carapace`

![](./hidden.cast)

## CARAPACE_LENIENT

Whether to allow unknown flags.

- `0` - disabled
- `1` - enabled

![](./lenient.cast)
      
## CARAPACE_LOG

Whether to enable logging.

- `0` - disabled
- `1` - enabled
          
![](./log.cast)

## CARAPACE_MATCH

Whether to match case insensitive.

- `0` - case sensitive
- `1` - case insensitive

![](./match.cast)

## CARAPACE_MERGEFLAGS

Whether to merge flags to single tag group.

- `0` - disabled
- `1` - enabled

![](./mergeflags.cast)

> Enabled by default in Zsh for side-by-side view.

## CARAPACE_NOSPACE

Extend suffixes that prevent space suffix.

- `*` - matches all

![](./nospace.cast)
        
## CARAPACE_SANDBOX

Mock context for sandbox tests (internal).
      
## CARAPACE_TOOLTIP

Whether to enable tooltip style.

- `0` - disabled
- `1` - enabled

![](./tooltip.cast)

> Only affects Powershell.

## CARAPACE_UNFILTERED

Whether to skip the final filtering step.

![](./unfiltered.cast)

> This enables fuzzy completion in [Fish], but only works for (mostly) static values.

## CARAPACE_ZSH_HASH_DIRS

Zsh hash directories (internal).


[Fish]:https://fishshell.com/
[Zsh]:https://www.zsh.org/
