# Examples

## github.yaml

```yaml
name: github
completion:
  positional:
    - ["$_tools.gh.OwnerRepositories"]  # ${C_ARG0}
    - ["$_tools.git.LsRemoteRefs({url: 'https://github.com/${C_ARG0}', branches: true, tags: true})"]
```

![](./examples-github.cast)

## zipfile.yaml

```yaml
name: zipfile
completion:
  positional:
    - ["$files([.zip])"] # ${C_ARG0}
  positionalany: ["$_fs.ZipFileContents(${C_ARG0})"] # ${C_ARG1},${C_ARG2},...
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
    - ["$_tools.git.Refs({tags: ${C_FLAG_TAGS:-false}, localbranches: ${C_FLAG_LOCALBRANCHES:-false}, commits: ${C_FLAG_C:-0}})"]
    - ["$_tools.git.Refs"] # default
```

![](./examples-refs.cast)

## embed.yaml

```yaml
name: embed
commands:
  - name: git
    completion:
      positionalany: ["$chdir(~/.password-store)", "$_bridge.CarapaceBin(git)"]
```

![](./examples-embed.cast)
