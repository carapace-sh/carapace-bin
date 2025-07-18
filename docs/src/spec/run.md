# Run

Specs containing a `run` field can be executed using [Shims](shim.md).

## Alias

Alias bridges a command while retaining the argument completion.

```yaml
# yaml-language-server: $schema=https://carapace.sh/schemas/command.json
name: example
run: "[gh, issue, --repo, '${REPO}']"
```

![](./run/alias.cast)

## Script

Script macro is executed with `sh` on unix systems and `pwsh` on windows.
Flags are used for [environment substitution](https://github.com/drone/envsubst) and positional arguments are passed to the script.

```yaml
# yaml-language-server: $schema=https://carapace.sh/schemas/command.json
name: ls-remote
run: "$(git ls-remote --sort='${C_FLAG_SORT:-HEAD}' $@)"
flags:
  --sort=: field name to sort on
completion:
  flag:
    sort: [version:refname, authordate]
  positional:
    - ["$carapace.tools.git.RepositorySearch"]
  positionalany: ["$carapace.tools.git.LsRemoteRefs({url: '${C_ARG0}', branches: true, tags: true})"]
```

![](./run/script.cast)
