# Examples

## github.yaml

```yaml
name: github
completion:
  positional:
    - ["$_tools.gh.OwnerRepositories"]  # ${C_ARG0}
    - ["$_tools.git.LsRemoteRefs({url: 'https://github.com/${C_ARG0}', branches: true, tags: true})"]
```

## zipfile.yaml

```yaml
name: zipfile
completion:
  positional:
    - ["$files([.zip])"] # ${C_ARG0}
  positionalany: ["$_fs.ZipFileContents(${C_ARG0})"] # ${C_ARG1},${C_ARG2},...
```

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
    - ["$_tools.git.Refs"]
    - ["$(env)"]
```
