# Embed

The [bridged](./bride.md) completion can also be embedded as subcommand.

```yaml
name: embed
commands:
  - name: git
    completion:
      positionalany: ["$chdir(~/.password-store)", "$_bridge.CarapaceBin(git)"]
```

![](./embed.cast)

## Plugin

Embedding is also internally used to enable plugin completion for tools like `cargo-[plugin]`, `gh-[plugin]`, `git-[plugin]`. Simply add a [Spec](../spec.md) with the corresponding name.

E.g. the [gh-repo-collab](https://github.com/mislav/gh-repo-collab) extension for [GitHub CLI](https://cli.github.com/):

```yaml
name: gh-repo-collab
description: manage repository collaborators
commands:
  -
    name: list
    completion:
      positional:
        - ["$_tools.gh.OwnerRepositories"]
  -
    name: add
    flags:
      --permission=: set permission
    completion:
      flag:
        permission: ["pull", "triage", "push", "maintain", "admin\t\tred"]
      positional:
        - ["$_tools.gh.OwnerRepositories"]
        - ["$_tools.gh.Users"]
  -
    name: remove
    completion:
      positional:
        - ["$_tools.gh.OwnerRepositories"]
        - ["$_tools.gh.Users"]
```

![](./embed-plugin.cast)
