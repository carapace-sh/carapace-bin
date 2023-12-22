# Examples

## github.yaml

```yaml
name: github
completion:
  positional:
    - ["$carapace.tools.gh.OwnerRepositories"]  # ${C_ARG0}
    - ["$carapace.tools.git.LsRemoteRefs({url: 'https://github.com/${C_ARG0}', branches: true, tags: true})"]
```

![](./examples-github.cast)

## zipfile.yaml

```yaml
name: zipfile
completion:
  positional:
    - ["$files([.zip])"] # ${C_ARG0}
  positionalany: ["$carapace.fs.ZipFileContents(${C_ARG0})"] # ${C_ARG1},${C_ARG2},...
```

![](./examples-zipfile.cast)

## refs.yaml

```yaml
name: refs
flags:
  -t, --tags: include tags # ${C_FLAG_TAGS}
  --localbranches: include local branches # ${C_FLAG_LOCALBRANCHES}
  --c=: amount of commits # ${C_FLAG_C}
completion:
  positional:
    - ["$carapace.tools.git.Refs({tags: ${C_FLAG_TAGS:-false}, localbranches: ${C_FLAG_LOCALBRANCHES:-false}, commits: ${C_FLAG_C:-0}})"]
    - ["$carapace.tools.git.Refs"] # default
```

![](./examples-refs.cast)


## g.yaml

```yaml
name: g
commands:
  - name: checkout
    description: Switch branches or restore working tree files
    group: git
    run: "[git, checkout]"

  - name: commit
    description: Record changes to the repository
    group: git
    run: "[git, commit]"

  - name: diff
    description: Show changes between commits
    group: git
    run: "[git, diff]"

  - name: log
    description: Show commit logs
    group: git
    run: "[git, log]"

  - name: push
    description: Update remote refs along with associated objects
    group: git
    run: "[git, push]"

  - name: switch
    description: Switch branches
    group: git
    run: "[git, checkout]"

  - name: issue
    description: Manage issues
    group: gh
    run: "[gh, issue]"
    commands:
      - name: bugs
        description: List bugs
        run: "[gh, issue, list, --label, bug]"

  - name: pr
    description: Manage pull requests
    group: gh
    run: "[gh, pr]"

  - name: edit
    description: Edit changed files
    run: "$(hx $@)"
    flags:
      -s, --staged: include staged files
    completion:
      positionalany: ["$carapace.tools.git.Changes({staged: ${C_FLAG_STAGED:-false}, unstaged: true})"]
```

![](./examples-g.cast)
