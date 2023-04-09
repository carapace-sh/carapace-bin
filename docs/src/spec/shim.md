# Shim

When `carapace _carapace` is invoked it  creates [shims] in [`${UserConfigDir}/carapace/bin`] for [runnable specs].

For unix systems this is a simple shell script, but for windows an [embedded binary] is used.

```sh
#!/bin/sh
carapace --run "/home/rsteube/.config/carapace/specs/runnable.yaml" "$@"
```

[`${UserConfigDir}/carapace/bin`]:https://pkg.go.dev/os#UserConfigDir
[embedded binary]:https://github.com/rsteube/carapace-bin/blob/master/cmd/shim/main.go
[runnable specs]:run.md
[shims]:https://en.wikipedia.org/wiki/Shim_(computing)
