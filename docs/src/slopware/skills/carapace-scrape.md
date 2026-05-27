# carapace-scrape

The [`carapace-scrape`](https://github.com/carapace-sh/carapace-bin/blob/master/skills/carapace-scrape/SKILL.md) skill creates a [Spec](../../spec.md) by scraping the source code of a CLI tool.

![](./carapace-scrape/carapace-scrape.cast)

## Instructions

1. start [crush] in an empty directory
2. `ctrl+p`, `TAB`, then select the **skill**
3. run with `scrape https://github.com/some/repo`

## Result

```yaml
{{#include ./carapace-scrape/pgxcli.yaml}}
```

## Intermediate build

```dockerfile
{{#include ./carapace-scrape/Dockerfile}}
```

```patch
{{#include ./carapace-scrape/pgxcli.patch}}
```


[crush]:https://github.com/charmbracelet/crush
