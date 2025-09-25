# Overlay

Overlays are essentially [Spec] files placed in [`${UserConfigDir}/carapace/overlays`] that provide additional completions.
These are merged with the existing completion and provide a workaround for issues that have yet to be fixed in upstream.

> Overlays implicitly set `CARAPACE_LENIENT` to allow unknown flags.

## Flag

```yaml
# ${UserConfigDir}/carapace/overlays/doctl.yaml
name: doctl
persistentflags:
  --output=: Desired output format [text|json]
completion:
  flag:
    output: [text, json]
commands:
  - name: compute
    description: Display commands that manage infrastructure
    commands:
      - name: region
        description: Display commands to list datacenter regions
        commands:
          - name: list
            description: List datacenter regions
            flags:
              --format=: Columns for output in a comma-separated list
            completion:
              flag:
                format: ["$uniquelist(,)", Slug, Name, Available]
```

![](./overlay-flag.cast)

## Command

```yaml
# ${UserConfigDir}/carapace/overlays/doctl.yaml
name: doctl
commands:
  - name: auth
    description: Display commands for authenticating doctl with an account
    group: management

  - name: compute
    description: Display commands that manage infrastructure
    group: core

  - name: custom
    description: custom command
    group: custom
    flags:
      -h, --help: show help
    completion:
      positional:
        - [one, two, three]
```

![](./overlay-command.cast)

[Spec]:./spec.md
[`${UserConfigDir}/carapace/overlays`]:https://pkg.go.dev/os#UserConfigDir
