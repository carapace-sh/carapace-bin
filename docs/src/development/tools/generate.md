# generate

[generate](https://github.com/rsteube/carapace-bin/tree/master/cmd/generate) is invoked by `go:generate` to implicitly register the completers.
It also copies them to `completers_release` for an optimized build where only the `init()` function of the relevant completer is invoked.
