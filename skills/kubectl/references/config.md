# kubeconfig and Contexts

How kubectl discovers cluster credentials through kubeconfig files, and how `kubectl config` commands manage clusters, users, contexts, and the current context.

> **Source of truth**: `kubectl config --help`, <https://kubernetes.io/docs/concepts/configuration/organize-cluster-access-kubeconfig/>. For **global connection flags**, see [cli.md](cli.md). For **authentication methods**, see [concepts.md](concepts.md).

## kubeconfig File

A kubeconfig file stores cluster connection information — clusters, users, and contexts. By default, kubectl looks for it at `~/.kube/config`.

### Resolution Order

kubectl resolves the kubeconfig file in this order:

1. **`--kubeconfig <path>` flag** — explicit path, highest priority
2. **`KUBECONFIG` env var** — one or more paths, merged in order
   - Linux/macOS: colon-separated (`KUBECONFIG=~/.kube/config:~/.kube/prod-config`)
   - Windows: semicolon-separated
3. **`~/.kube/config`** — default location

### Merging Multiple Files

When `KUBECONFIG` points to multiple files, kubectl **merges** them:

- Clusters, users, and contexts are combined (union)
- Later files override earlier ones for the same name
- The `current-context` from the **first** file wins

This enables layering: a base config with shared clusters, plus an overlay with personal contexts.

```shell
# Set multiple kubeconfigs — contexts and clusters are merged
export KUBECONFIG=~/.kube/config:~/.kube/work-config
kubectl config get-contexts
```

### kubeconfig Structure

```yaml
apiVersion: v1
kind: Config
clusters:
  - name: minikube
    cluster:
      server: https://192.168.49.2:8443
      certificate-authority: ~/.minikube/ca.crt
      insecure-skip-tls-verify: false
users:
  - name: minikube-user
    user:
      client-certificate: ~/.minikube/client.crt
      client-key: ~/.minikube/client.key
      token: <bearer-token>
contexts:
  - name: minikube
    context:
      cluster: minikube
      user: minikube-user
      namespace: default
current-context: minikube
preferences: {}
```

### Components

| Component | Purpose | Key Fields |
|-----------|---------|------------|
| **cluster** | API server endpoint and CA | `server`, `certificate-authority`, `insecure-skip-tls-verify`, `tls-server-name` |
| **user** (credentials) | Authentication credentials | `client-certificate`, `client-key`, `token`, `username`, `password`, `exec` |
| **context** | A (cluster, user, namespace) tuple | `cluster`, `user`, `namespace` |
| **current-context** | Which context kubectl uses by default | Name of a context |

### Embedded Certificates

Certificates can be inline (base64-encoded) or file paths:

```yaml
# Inline (base64)
cluster:
  certificate-authority-data: LS0tLS1CRUdJTi...

# File path
cluster:
  certificate-authority: /path/to/ca.crt
```

## `kubectl config` Subcommands

### Viewing

| Command | Description |
|---------|-------------|
| `kubectl config view` | Display merged kubeconfig |
| `kubectl config view --raw` | Display without redacting certificate data |
| `kubectl config view -o jsonpath='{.contexts[*].name}'` | Extract specific fields |
| `kubectl config get-contexts` | List all contexts (marks current with `*`) |
| `kubectl config get-clusters` | List all clusters |
| `kubectl config get-users` | List all users |
| `kubectl config current-context` | Print the current context name |

```shell
kubectl config get-contexts
# CURRENT   NAME       CLUSTER    AUTHINFO   NAMESPACE
# *         minikube   minikube   minikube   default
#           prod       prod-api   prod-user  production
```

### Switching Context

```shell
kubectl config use-context minikube
kubectl config use-context prod
```

This sets `current-context` in the kubeconfig file. All subsequent kubectl commands use the selected context's cluster, user, and namespace.

### Creating and Modifying Contexts

```shell
# Create a new context
kubectl config set-context dev \
  --cluster=dev-cluster \
  --user=dev-user \
  --namespace=dev

# Rename a context
kubectl config rename-context old-name new-name

# Set the namespace for the current context
kubectl config set-context --current --namespace=staging

# Delete a context
kubectl config delete-context old-context
```

### Creating and Modifying Clusters

```shell
kubectl config set-cluster prod-cluster \
  --server=https://prod-api.example.com:6443 \
  --certificate-authority=~/certs/prod-ca.crt \
  --insecure-skip-tls-verify=false

# Set embedded CA (base64)
kubectl config set-cluster prod-cluster \
  --server=https://prod-api.example.com:6443 \
  --embed-certs=true \
  --certificate-authority-data=<base64>

kubectl config delete-cluster old-cluster
kubectl config get-clusters
```

### Creating and Modifying Users (Credentials)

```shell
kubectl config set-credentials prod-user \
  --client-certificate=~/certs/prod-client.crt \
  --client-key=~/certs/prod-client.key

# Token-based
kubectl config set-credentials prod-user --token=<bearer-token>

# Exec plugin (e.g., cloud auth)
kubectl config set-credentials gcp-user \
  --exec-command=gcloud \
  --exec-args=auth,print-access-token \
  --exec-api-version=client.authentication.k8s.io/v1

kubectl config delete-user old-user
kubectl config get-users
```

### Setting and Unsetting Arbitrary Values

```shell
kubectl config set clusters.prod-cluster.proxy-url http://proxy.example.com
kubectl config unset users.prod-user.client-key
```

These use a dotted path syntax into the kubeconfig YAML structure.

## Common Workflows

### Switching Between Clusters

```shell
kubectl config use-context dev
kubectl get pods                    # lists pods in dev cluster

kubectl config use-context prod
kubectl get pods                    # lists pods in prod cluster
```

### Temporary Namespace Override

```shell
# Override namespace for one command (doesn't change context)
kubectl get pods -n kube-system

# Permanently set namespace in current context
kubectl config set-context --current --namespace=monitoring
kubectl get pods                    # now lists pods in monitoring namespace
```

### Temporary Context Override

```shell
# Override context for one command (doesn't change current-context)
kubectl get pods --context=prod
kubectl get nodes --context=prod --namespace=default
```

### Adding a New Cluster

```shell
# 1. Add the cluster
kubectl config set-cluster staging \
  --server=https://staging-api.example.com:6443 \
  --certificate-authority=~/certs/staging-ca.crt

# 2. Add credentials
kubectl config set-credentials staging-admin \
  --token=<staging-token>

# 3. Create a context binding them together
kubectl config set-context staging \
  --cluster=staging \
  --user=staging-admin \
  --namespace=default

# 4. Switch to it
kubectl config use-context staging
```

### Merging a kubeconfig File

To merge a provided kubeconfig (e.g., from a cloud provider) into your default config:

```shell
# Append to KUBECONFIG for this session
export KUBECONFIG=~/.kube/config:~/Downloads/new-cluster.yaml

# Or merge permanently into default config
KUBECONFIG=~/.kube/config:~/Downloads/new-cluster.yaml kubectl config view --flatten > ~/.kube/config
```

The `--flatten` flag produces a single merged file with embedded certificates.

## Programmatic Access

`kubectl config view` with jsonpath can extract specific values for scripts:

```shell
# Get current context
kubectl config current-context

# Get current namespace
kubectl config view --minify --output 'jsonpath={..namespace}'

# Get API server URL for current context
kubectl config view --minify --output 'jsonpath={.clusters[0].cluster.server}'

# Get all context names
kubectl config view -o jsonpath='{.contexts[*].name}'
```

The `--minify` flag outputs only the current context's relevant entries, which is useful for scripting.

## kubeconfig and Global Flags

Global flags override kubeconfig settings for a single command:

| Flag | Overrides |
|------|-----------|
| `--kubeconfig` | Which file(s) to load |
| `--server` | Context's cluster server URL |
| `--cluster` | Which cluster entry to use |
| `--user` | Which user/credentials to use |
| `--context` | Which context to use (cluster + user + namespace) |
| `--namespace` | Context's namespace |
| `--token` | User's bearer token |
| `--certificate-authority` | Cluster's CA |
| `--client-certificate` | User's client cert |
| `--client-key` | User's client key |
| `--insecure-skip-tls-verify` | Cluster's TLS verification |

```shell
# Use a specific context but override the namespace
kubectl get pods --context=prod --namespace=monitoring

# Connect to a server directly, bypassing kubeconfig
kubectl get pods --server=https://api.example.com:6443 --token=abc123 --insecure-skip-tls-verify
```

## Edge Cases and Known Issues

- **`KUBECONFIG` path separator differs by OS**: Colon on Linux/macOS, semicolon on Windows.
- **Merging order matters**: Later files in `KUBECONFIG` override earlier entries with the same name. The first file's `current-context` wins.
- **`config view` redacts sensitive data**: Certificate data and tokens are hidden by default. Use `--raw` to see them (for debugging, not for sharing).
- **`--flatten` embeds certificates**: `kubectl config view --flatten` converts file-path certificates to inline base64, producing a self-contained file.
- **Exec credential plugins**: Users configured with `exec` run an external command to fetch credentials. The command must be on `PATH` and is invoked on each kubectl request (cached briefly).
- **Context namespace is a default, not a constraint**: `--namespace` always overrides the context's namespace. Some cluster-scoped operations ignore it entirely.
- **Renaming a context does not update `current-context`**: If you rename the current context, you must `use-context` with the new name.
- **`config set-context --current`**: This is a shortcut to modify the current context in place — commonly used for changing the default namespace.

## References

- [Organize Cluster Access with kubeconfig](https://kubernetes.io/docs/concepts/configuration/organize-cluster-access-kubeconfig/)
- [kubectl config](https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands#config)

## Related Skills

- For global connection flags, see [cli.md](cli.md).
- For authentication methods (client cert, token, in-cluster), see [concepts.md](concepts.md).
