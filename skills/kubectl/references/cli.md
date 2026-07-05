# kubectl CLI Reference

Command-line reference for kubectl — the command groups, global flags, and command structure. For command-specific details, see the cross-referenced documents.

> **Source of truth**: `kubectl --help`, `kubectl options`, <https://kubernetes.io/docs/reference/kubectl/generated/kubectl/>. For **concepts**, see [concepts.md](concepts.md). For **kubeconfig flags**, see [config.md](config.md). For **output format flags**, see [output.md](output.md).

## Command Syntax

```
kubectl [command] [TYPE] [NAME] [flags]
```

| Component | Description |
|-----------|-------------|
| `command` | The operation to perform (`get`, `apply`, `delete`, etc.) |
| `TYPE` | Resource type (`pod`, `deployment`, `svc`) — case-insensitive, singular/plural/abbreviated |
| `NAME` | Resource name — case-sensitive |
| `flags` | Command-specific and global flags |

## Command Groups

kubectl organizes commands into groups (shown in `kubectl --help`):

### Basic Commands (Beginner)

| Command | Description |
|---------|-------------|
| `create` | Create a resource from file or stdin |
| `get` | Display one or many resources |
| `run` | Run a particular image on the cluster |
| `expose` | Expose a resource as a new Kubernetes Service |
| `set` | Set specific features on objects |

### Basic Commands (Intermediate)

| Command | Description |
|---------|-------------|
| `explain` | Documentation of a resource |
| `delete` | Delete resources by filenames, stdin, resources, labels, or selectors |
| `wait` | Wait for a specific condition on one or many resources |
| `events` | List events |

### Deploy Commands

| Command | Description |
|---------|-------------|
| `rollout` | Manage the rollout of a resource |
| `scale` | Set a new size for a Deployment, ReplicaSet, or ReplicationController |
| `autoscale` | Auto-scale a Deployment, ReplicaSet, or ReplicationController |

### Cluster Management Commands

| Command | Description |
|---------|-------------|
| `certificate` | Modify certificate resources |
| `cluster-info` | Display cluster info |
| `top` | Display Resource (CPU/Memory/Storage) usage |
| `cordon` | Mark node as unschedulable |
| `uncordon` | Mark node as schedulable |
| `drain` | Drain node in preparation for maintenance |
| `taint` | Update the taints on one or more nodes |

### Troubleshooting and Debugging Commands

| Command | Description |
|---------|-------------|
| `describe` | Show details of a specific resource or group of resources |
| `logs` | Print logs for a container in a pod |
| `attach` | Attach to a running container |
| `exec` | Execute a command in a container |
| `port-forward` | Forward one or more local ports to a pod |
| `proxy` | Run a proxy to the Kubernetes API server |
| `cp` | Copy files and directories to and from containers |
| `auth` | Inspect authorization |
| `debug` | Create debugging sessions for troubleshooting workloads and nodes |

### Advanced Commands

| Command | Description |
|---------|-------------|
| `diff` | Diff live version against would-be applied version |
| `apply` | Apply a configuration to a resource by filename or stdin |
| `patch` | Update field(s) of a resource |
| `replace` | Replace a resource by filename or stdin |
| `kustomize` | Build a kustomization target from a directory or URL |

### Settings Commands

| Command | Description |
|---------|-------------|
| `label` | Update the labels on a resource |
| `annotate` | Update the annotations on a resource |
| `config` | Modify kubeconfig files |

### Plugin Commands

| Command | Description |
|---------|-------------|
| `plugin` | Provides utilities for interacting with plugins |

### Other Commands

| Command | Description |
|---------|-------------|
| `alpha` | Commands for features in alpha |
| `api-resources` | Print the supported API resources on the server |
| `api-versions` | Print the supported API versions on the server |
| `completion` | Output shell completion code for the specified shell |
| `kuberc` | Modify kuberc preferences |
| `options` | Print the list of flags inherited by all commands |
| `version` | Print the client and server version information |

## Global Flags

All kubectl commands inherit these persistent flags:

### Cluster Connection

| Flag | Description |
|------|-------------|
| `--kubeconfig <path>` | Path to the kubeconfig file to use |
| `--server`, `-s` | Address and port of the Kubernetes API server |
| `--cluster` | The name of the kubeconfig cluster to use |
| `--context` | The name of the kubeconfig context to use |
| `--user` | The name of the kubeconfig user to use |
| `--cache-dir` | Default cache directory (default `$HOME/.kube/cache`) |

### Namespace and Scope

| Flag | Description |
|------|-------------|
| `--namespace`, `-n` | Namespace scope for this CLI request |
| `--all-namespaces`, `-A` | If present, list across all namespaces (context namespace ignored) |

### TLS

| Flag | Description |
|------|-------------|
| `--certificate-authority` | Path to cert file for the CA |
| `--client-certificate` | Path to client certificate for TLS |
| `--client-key` | Path to client key for TLS |
| `--insecure-skip-tls-verify` | Skip server certificate validation (insecure) |
| `--tls-server-name` | Server name for cert validation |
| `--disable-compression` | Opt-out of response compression |

### Authentication

| Flag | Description |
|------|-------------|
| `--token` | Bearer token for API server auth |
| `--username` | Username for basic auth (deprecated) |
| `--password` | Password for basic auth (deprecated) |

### Impersonation

| Flag | Description |
|------|-------------|
| `--as` | Username to impersonate |
| `--as-group` | Group to impersonate (repeatable) |
| `--as-uid` | UID to impersonate |
| `--as-user-extra` | User extras to impersonate (repeatable) |

### Request Control

| Flag | Description |
|------|-------------|
| `--request-timeout` | Length to wait before giving up on a single request (e.g. `1s`, `2m`) |
| `--match-server-version` | Require server version to match client version |

### Logging and Profiling

| Flag | Description |
|------|-------------|
| `--v`, `-v` | Log level verbosity (e.g. `--v=5` for detailed, `--v=8` for very verbose) |
| `--vmodule` | Comma-separated `pattern=N` for file-filtered logging |
| `--log-flush-frequency` | Maximum seconds between log flushes |
| `--profile` | Profile to capture: `none`, `cpu`, `heap`, `goroutine`, `threadcreate`, `block`, `mutex` |
| `--profile-output` | File to write the profile to |
| `--warnings-as-errors` | Treat server warnings as errors |

### Other

| Flag | Description |
|------|-------------|
| `--kuberc` | Path to the kuberc preferences file |

## Common Command-Specific Flags

These flags appear on many resource-management commands:

### File Input

| Flag | Description |
|------|-------------|
| `-f`, `--filename` | Filename, directory, or URL to resource config |
| `-k`, `--kustomize` | Process a kustomization directory |
| `-R`, `--recursive` | Process directory in `-f` recursively |

### Output

| Flag | Description |
|------|-------------|
| `-o`, `--output` | Output format (see [output.md](output.md)) |
| `--show-kind` | Show resource type in output |
| `--show-labels` | Show all labels in output |
| `--show-managed-fields` | Keep managedFields in JSON/YAML output |
| `--no-headers` | Don't print table headers |
| `--label-columns`, `-L` | Additional label columns |
| `--sort-by` | Sort by JSONPath expression |

### Selection

| Flag | Description |
|------|-------------|
| `-l`, `--selector` | Label selector |
| `--field-selector` | Field selector |
| `--all` | Select all resources in namespace |
| `-A`, `--all-namespaces` | Across all namespaces |

### Validation and Dry Run

| Flag | Description |
|------|-------------|
| `--dry-run` | `none`, `server`, or `client` — preview without applying |
| `--validate` | `strict`/`true`, `warn`, `ignore`/`false` — schema validation |
| `--field-manager` | Name of manager for field ownership tracking |

### Template

| Flag | Description |
|------|-------------|
| `--template` | Template string or file path for go-template output |
| `--allow-missing-template-keys` | Ignore missing keys in templates |

## Verbosity Levels

The `--v` flag controls logging verbosity:

| Level | What you see |
|-------|-------------|
| `--v=0` | Default — minimal output |
| `--v=2` | Useful steady-state info |
| `--v=4` | Debug-level verbosity |
| `--v=6` | Request URLs and method |
| `--v=7` | Request and response headers |
| `--v=8` | Request and response body |

```shell
kubectl get pods --v=6       # show API request URLs
kubectl get pods --v=8       # show full request/response bodies
```

## Completion

kubectl can generate shell completion scripts:

```shell
kubectl completion bash > /etc/bash_completion.d/kubectl
kubectl completion zsh > "${fpath[1]}/_kubectl"
kubectl completion fish > ~/.config/fish/completions/kubectl.fish
kubectl completion powershell | Out-String | Invoke-Expression
```

When carapace provides the kubectl completer, completion is handled by carapace instead (see the **carapace** skill).

## `kubectl options`

Lists all global command-line options inherited by every command:

```shell
kubectl options
```

## `kubectl version`

Prints client and server version info:

```shell
kubectl version                       # both client and server
kubectl version --client              # client only (no cluster connection)
kubectl version --output=yaml         # structured output
```

## Edge Cases and Known Issues

- **`--namespace` is ignored for cluster-scoped resources**: `kubectl get nodes -n dev` lists all nodes regardless of namespace.
- **`--all-namespaces` overrides `--namespace`**: If both are set, `--all-namespaces` takes precedence.
- **Global flags must come after the subcommand**: `kubectl get pods --namespace=dev` (correct), not `kubectl --namespace=dev get pods` (may work but is not the documented form).
- **`--record` is deprecated**: The `--record` flag that annotated resources with the kubectl command is deprecated and hidden. Use server-side apply with field management instead.
- **Basic auth flags are deprecated**: `--username` and `--password` are deprecated; use client certificates or tokens.
- **`completion` output is large**: The generated completion script can be hundreds of lines. Redirect to a file rather than printing to terminal.
- **Some commands have hidden flags**: Flags like `--record` may be hidden from `--help` but still accepted.

## References

- [kubectl Commands](https://kubernetes.io/docs/reference/kubectl/generated/kubectl/)
- [kubectl Overview](https://kubernetes.io/docs/reference/kubectl/)

## Related Skills

- For kubeconfig and context management, see [config.md](config.md).
- For output formatting, see [output.md](output.md).
- For workload commands (rollout, scale, autoscale, expose), see [workloads.md](workloads.md).
- For debugging commands, see [debugging.md](debugging.md).
