# Spec

Custom completions can be defined using yaml files.

> see [carapace-spec] for more documentation

```yaml
# yaml-language-server: $schema=https://carapace.sh/schemas/command.json
name: mycmd
description: my command
flags:
  --optarg?: optarg flag
  -r, --repeatable*: repeatable flag
  -v=: flag with value
persistentflags:
  --help: bool flag
completion:
  flag:
    optarg: ["one", "two\twith description", "three\twith style\tblue"]
    v: ["$files"]
commands:
- name: sub
  description: subcommand
  completion:
    positional:
      - ["$list(,)", "1", "2", "3"]
      - ["$directories"]
```

![](./spec.cast)

## Custom Macros

Carapace provides a range of [custom macros](./spec/macros.md):

```sh
carapace --macro                       # list macros
carapace --macro color.HexColors       # show macro details
carapace --macro color.HexColors <TAB> # test macro
```

![](./spec-macros.cast)


[carapace-spec]:https://github.com/rsteube/carapace-spec
