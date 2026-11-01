# carapace-lint

A [golangci-lint](https://golangci-lint.run) plugin enforcing carapace command definition conventions.
The plugin is built into a custom golangci-lint binary using `.custom-gcl.yml` and configured in `.golangci.yml`.

```sh
golangci-lint custom      # build custom-gcl from .custom-gcl.yml
./custom-gcl run ./...    # lint
./custom-gcl run --fix ./... # lint with auto-fixes
```

## Analyzers

| Setting | Description |
|---------|-------------|
| `flags-order` | flags defined with `cmd.Flags().<Type>(...)` have to be alphabetically ordered |
| `flag-actions-order` | entries of `FlagCompletion(carapace.ActionMap{...})` have to be alphabetically ordered |
| `macro-syntax` | exported actions under `pkg/actions/` have to match the macro syntax |
| `completion-values` | completion values have to be paired correctly and must be unique |

All checks are enabled by default.
Ordering and duplicate-entry findings support auto-fix with `--fix`.
