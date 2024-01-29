# Embed

The [bridged](./bridge.md) completion can also be embedded as subcommand.

```yaml
# yaml-language-server: $schema=https://carapace.sh/schemas/command.json
name: embed
commands:
  - name: git
    completion:
      positionalany: ["$carapace.bridge.CarapaceBin([git]) ||| $chdir(~/.password-store)"]
```

![](./embed.cast)

## Plugin

Embedding is also internally used to enable plugin completion for tools like `cargo-[plugin]`, `gh-[plugin]`, `git-[plugin]`. Simply add a [Spec](../spec.md) with the corresponding name.

E.g. the [gh-repo-collab](https://github.com/mislav/gh-repo-collab) extension for [GitHub CLI](https://cli.github.com/):

```yaml
# yaml-language-server: $schema=https://carapace.sh/schemas/command.json
name: gh-repo-collab
description: manage repository collaborators
commands:
  -
    name: list
    completion:
      positional:
        - ["$carapace.tools.gh.OwnerRepositories"]
  -
    name: add
    flags:
      --permission=: set permission
    completion:
      flag:
        permission: ["pull", "triage", "push", "maintain", "admin\t\tred"]
      positional:
        - ["$carapace.tools.gh.OwnerRepositories"]
        - ["$carapace.tools.gh.Users"]
  -
    name: remove
    completion:
      positional:
        - ["$carapace.tools.gh.OwnerRepositories"]
        - ["$carapace.tools.gh.Users"]
```

![](./embed-plugin.cast)
