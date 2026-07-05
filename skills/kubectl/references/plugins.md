# kubectl Plugin System

How kubectl discovers and executes plugins, the `kubectl plugin` command, the `kubectl-` naming convention, the Krew plugin manager, and the carapace completer's plugin integration.

> **Source of truth**: <https://kubernetes.io/docs/tasks/extend-kubectl/kubectl-plugins/>. For **carapace completer plugin integration** (how `completers/common/kubectl_completer/` discovers plugins), see the **carapace** skill.

## Plugin Discovery

kubectl supports **plugins** — external executables that extend kubectl with new subcommands. Plugins are discovered by scanning the `PATH` for executables with the `kubectl-` prefix.

### How It Works

1. kubectl receives a command (e.g., `kubectl foo bar`)
2. If `foo` is not a built-in command, kubectl checks for a plugin
3. It looks for an executable named `kubectl-foo` in the `PATH`
4. If found, kubectl executes it with the remaining arguments (`kubectl-foo bar`)

### Naming Convention

```
kubectl-<command>          → kubectl <command>
kubectl-<command>-<sub>    → kubectl <command> <sub>
```

Examples:

| Executable | Invoked as |
|-----------|------------|
| `kubectl-foo` | `kubectl foo` |
| `kubectl-bar-baz` | `kubectl bar baz` |
| `kubectl-do-thing` | `kubectl do thing` |

The plugin name is derived by stripping the `kubectl-` prefix and splitting on `-` into subcommands.

### Plugin Path

kubectl searches directories in the `PATH` environment variable for executables matching the `kubectl-` prefix. There is no separate plugin directory or plugin path environment variable in modern kubectl (the legacy `~/.kube/plugins` directory and `KUBECTL_PLUGIN_PATH` env var were removed in kubectl 1.12).

## `kubectl plugin`

The `plugin` command manages plugins:

```shell
kubectl plugin list
kubectl plugin list --name-only
```

### `plugin list`

Lists all discovered plugins:

```shell
kubectl plugin list
# The following compatible plugins are available:
# /usr/local/bin/kubectl-foo
#   - foo
# /usr/local/bin/kubectl-krew
#   - krew
```

The `--name-only` flag outputs just the plugin names (one per line, with a header), which is useful for scripting.

## Writing a Plugin

A plugin is any executable (shell script, Python, Go binary, etc.) named `kubectl-<name>`:

### Shell Script Plugin

```bash
#!/bin/bash
# kubectl-hello — a simple kubectl plugin

echo "Hello from kubectl plugin!"
echo "Arguments: $@"
kubectl get pods "$@"
```

Install it:

```shell
chmod +x kubectl-hello
sudo mv kubectl-hello /usr/local/bin/
kubectl hello                    # runs the plugin
kubectl hello -n kube-system     # passes args through
```

### Plugin Environment

Plugins inherit the environment from kubectl, including:
- `KUBECONFIG` — kubeconfig path
- All global kubectl settings

Plugins can call `kubectl` internally to interact with the cluster:

```bash
#!/bin/bash
# kubectl-pod-logs — get logs for the first pod matching a label

POD=$(kubectl get pods -l app="$1" -o jsonpath='{.items[0].metadata.name}')
kubectl logs "$POD"
```

### Plugin Limitations

- Plugins **cannot override** built-in kubectl commands. If `foo` is a built-in, `kubectl-foo` is never invoked.
- Plugins are **not sandboxed** — they run with the user's permissions and can access kubeconfig credentials.
- kubectl does not verify plugin integrity or authenticity.
- Flag parsing is the plugin's responsibility — kubectl passes all arguments through.

## Krew — Plugin Manager

[Krew](https://krew.sigs.k8s.io/) is the plugin manager for kubectl. It installs, updates, and removes plugins from a central index:

```shell
# Install Krew
(
  set -x; cd "$(mktemp -d)" &&
  OS="$(uname | tr '[:upper:]' '[:lower:]')" &&
  ARCH="$(uname -m | sed -e 's/x86_64/amd64/' -e 's/\(arm64\|aarch64\)/arm64/')" &&
  KREW="krew-${OS}_${ARCH}" &&
  curl -fsSLO "https://github.com/kubernetes-sigs/krew/releases/latest/download/${KREW}.tar.gz" &&
  tar zxvf "${KREW}.tar.gz" &&
  ./"${KREW}" install krew
)

# Use Krew
kubectl krew search
kubectl krew install neat          # install the "neat" plugin
kubectl krew install konfig        # install the "konfig" plugin
kubectl krew list                  # list installed plugins
kubectl krew upgrade               # update all installed plugins
kubectl krew uninstall neat        # remove a plugin
kubectl krew update                # update the plugin index
kubectl krew info neat             # show plugin details
```

### Common Krew Plugins

| Plugin | Description |
|--------|-------------|
| `krew` | The plugin manager itself |
| `neat` | Clean up `kubectl get -o yaml` output (remove defaults) |
| `konfig` | Merge, extract, and manage kubeconfig files |
| `ctx` | Quick context switching |
| `ns` | Quick namespace switching |
| `tree` | Show ownership tree of resources |
| `view-secret` | Decode Kubernetes secrets |
| `exec-as` | Exec into pods as a specific user |
| `debug-pod` | Debug a pod by creating a copy with debug tools |
| `tail` | Tail logs from multiple pods simultaneously |
| `status` | Show pod status with more detail |
| `who-can` | Show who can perform an action (RBAC) |
| `access-matrix` | Show RBAC access matrix |
| `trace` | bpftrace for Kubernetes |
| `node-shell` | Get a shell on a node |

## `kubectl alpha`

Some kubectl features start in alpha and are accessible via the `alpha` command:

```shell
kubectl alpha --help
```

Alpha commands are experimental and may be removed or changed in future versions.

## `kubectl convert`

The `convert` command converts config files between API versions. It's provided as a **plugin** (not built-in) in modern kubectl versions:

```shell
kubectl convert -f deployment.yaml --output-version=apps/v1
```

If `kubectl convert` is not found, install the `kubectl-convert` plugin.

## Carapace Completer Plugin Integration

The carapace kubectl completer (`completers/common/kubectl_completer/`) has special handling for plugins. In `cmd/root.go`, a `PreRun` hook dynamically adds plugin commands to the cobra command tree:

```go
carapace.Gen(rootCmd).PreRun(func(cmd *cobra.Command, args []string) {
    if _, _, err := cmd.Find(args); len(args) > 1 && err == nil {
        return // core command - skip plugin commands
    }
    addPluginCommands()
})
```

The `addPluginCommands()` function:
1. Runs `kubectl plugin list --name-only`
2. For each discovered plugin, creates a synthetic cobra command
3. Sets `DisableFlagParsing: true` (flags are passed through to the plugin)
4. Assigns it to the `plugin` command group
5. Registers `PositionalAnyCompletion` that bridges to the plugin via `bridge.ActionCarapaceBin(name)`

This means carapace completion works for kubectl plugins that themselves have carapace completion installed.

For details on the carapace bridge mechanism, see the **carapace** skill → `references/integrate.md`.

## Plugin Best Practices

### Naming

- Use a short, descriptive name: `kubectl-who-can`, not `kubectl-check-rbac-permissions-for-user`
- For multi-word plugins, use hyphens: `kubectl-view-secret`
- Avoid names that conflict with built-in commands

### Distribution

- **Krew**: Submit to the [krew-index](https://github.com/kubernetes-sigs/krew-index) for discoverability
- **GitHub releases**: Distribute the binary with installation instructions
- **Homebrew**: For macOS users, consider a Homebrew formula

### Implementation

- Handle `--help` and `-h` flags
- Validate arguments and provide clear error messages
- Use `kubectl` internally for cluster operations (reuse auth, kubeconfig)
- Support the `KUBECONFIG` environment variable
- Make the plugin work with any namespace (`-n` flag passthrough)

## Edge Cases and Known Issues

- **Built-in commands take priority**: If a plugin has the same name as a built-in command, the built-in always wins. The plugin is never invoked.
- **Plugin name splitting**: `kubectl-foo-bar` becomes `kubectl foo bar` (two subcommands), not `kubectl foo-bar`. This can be surprising.
- **Plugins are not validated**: kubectl runs any executable matching the naming convention. Be cautious about installing plugins from untrusted sources.
- **Krew plugins live in `~/.krew/bin`**: This directory must be in your `PATH` for kubectl to discover Krew-installed plugins.
- **`plugin list` shows warnings for invalid plugins**: If a plugin file is not executable or has the wrong naming, it's listed with a warning.
- **Flag parsing differs**: Built-in kubectl commands parse flags with pflag. Plugins receive raw arguments and must parse their own flags. Global flags like `--namespace` are not automatically parsed by kubectl for plugins.
- **`--as` and other global flags**: Global flags are passed through to plugins as arguments. The plugin must handle them (or call `kubectl` internally, which will re-parse them).

## References

- [kubectl Plugins](https://kubernetes.io/docs/tasks/extend-kubectl/kubectl-plugins/)
- [Krew Plugin Manager](https://krew.sigs.k8s.io/)
- [krew-index Repository](https://github.com/kubernetes-sigs/krew-index)

## Related Skills

- For carapace bridge and completion integration, see the **carapace** skill → `references/integrate.md`.
- For the kubectl command structure and groups, see [cli.md](cli.md).
