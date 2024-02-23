# Manually

- Copy a basic completer for simplicity.

```sh
cp -r completers/ln_completer completers/manually_completer
```

- Update the package name in `main.go`.

```diff
-import "github.com/rsteube/carapace-bin/completers/ln_completer/cmd"
+import "github.com/rsteube/carapace-bin/completers/manually_completer/cmd"
```

- Create the root command.

```sh
echo | carapace-parse -n manually > root.go
```

- Add subcommands.

```sh
echo | carapace-parse -n subcommand -p root > subcommand.go
```

- Define flags and completions.

- [Build and install](../build.md#development).

![](./manually.cast)
