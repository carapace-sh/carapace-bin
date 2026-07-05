# kubeadm token — Bootstrap Token Management

The `kubeadm token` command manages bootstrap tokens used by `kubeadm join` to establish trust between joining nodes and the control plane.

> **Source of truth**: <https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm-token/>. For token concepts, see [concepts.md](concepts.md).

## Subcommands

| Subcommand | Purpose |
|------------|---------|
| `kubeadm token create [token]` | Create a bootstrap token on the server |
| `kubeadm token delete [token-value] ...` | Delete bootstrap tokens |
| `kubeadm token generate` | Generate and print a random token (does not create on server) |
| `kubeadm token list` | List bootstrap tokens on the server |

## `token create`

Creates a new bootstrap token on the cluster. The token is stored as a Secret of type `bootstrap.kubernetes.io/token` in `kube-system`.

```bash
kubeadm token create
kubeadm token create abcdef.0123456789abcdef
kubeadm token create --print-join-command
kubeadm token create --ttl 48h --description "for worker pool"
```

### Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--certificate-key` | — | When used with `--print-join-command`, prints full control-plane join flags |
| `--config` | — | Path to a kubeadm config file |
| `--description` | — | Human-friendly description of the token's purpose |
| `--groups` | `system:bootstrappers:kubeadm:default-node-token` | Extra groups for authentication |
| `--print-join-command` | `false` | Print the full `kubeadm join` command instead of just the token |
| `--ttl` | `24h0m0s` | Duration before auto-deletion. `0` = never expire. |
| `--usages` | `signing,authentication` | Ways the token can be used |

### Token Format

```
abcdef.0123456789abcdef
```

- Regex: `[a-z0-9]{6}\.[a-z0-9]{16}`
- First part (6 chars) = **Token ID** (public)
- Second part (16 chars) = **Token Secret** (private)

If you provide a token argument, it must match this format. If you don't, kubeadm generates a random one.

### Usages

| Usage | Description |
|-------|-------------|
| `signing` | Sign the `cluster-info` ConfigMap during discovery |
| `authentication` | Use as a bearer token for API server authentication |

### `--print-join-command`

Prints the full `kubeadm join` command with the token filled in. Combined with `--certificate-key`, it prints the full control-plane join command:

```bash
$ kubeadm token create --print-join-command --certificate-key 1234...
kubeadm join 10.0.0.1:6443 --token abcdef.0123456789abcdef --discovery-token-ca-cert-hash sha256:... --control-plane --certificate-key 1234...
```

### Inherited Flags

| Flag | Description |
|------|-------------|
| `--dry-run` | Dry run mode (don't create the token) |
| `--kubeconfig` | Kubeconfig file to use (default: `/etc/kubernetes/admin.conf`) |

## `token delete`

Deletes one or more bootstrap tokens from the server. You can specify tokens by their full value or just the token ID (first 6 chars).

```bash
kubeadm token delete abcdef.0123456789abcdef
kubeadm token delete abcdef
kubeadm token delete abcdef xyz123
```

## `token generate`

Generates and prints a random bootstrap token. **Does not create it on the server** — use `token create` for that. This is useful for scripting:

```bash
TOKEN=$(kubeadm token generate)
kubeadm token create "$TOKEN"
```

Or to pass to `kubeadm init`:

```bash
TOKEN=$(kubeadm token generate)
kubeadm init --token "$TOKEN"
```

## `token list`

Lists all bootstrap tokens on the server.

```bash
kubeadm token list
```

### Flags

| Flag | Description |
|------|-------------|
| `--allow-missing-template-keys` | Ignore missing template fields (golang/jsonpath) |
| `-o, --output` | Output format: `text`, `json`, `yaml`, `kyaml`, `go-template`, `go-template-file`, `template`, `templatefile`, `jsonpath`, `jsonpath-as-json`, `jsonpath-file` |
| `--show-managed-fields` | Keep managedFields in JSON/YAML output |

### Output Format

The default `text` output shows a table with columns: TOKEN, TTL, EXPIRES, USAGES, DESCRIPTION, EXTRA GROUPS.

## Relationship to init and join

- `kubeadm init` automatically creates a bootstrap token (the `bootstrap-token` phase)
- The join command printed at the end of `init` includes this token
- Tokens expire (default 24h) — use `token create` to make new ones for joining nodes later
- `kubeadm join` uses the token for both discovery and TLS bootstrap

For the join workflow, see [join.md](join.md).

## Edge Cases

- **Expired tokens**: A token with a TTL of `24h` is automatically deleted by a controller after 24h. If you need a permanent token, set `--ttl 0`.
- **Token ID vs full token**: `token delete` accepts either the full token or just the ID. `token list` shows only the ID (not the secret).
- **Dry run**: `--dry-run` on `token create` shows what would happen without creating the Secret.

## References

- [kubeadm token docs](https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm-token/)
- [Bootstrap tokens](https://kubernetes.io/docs/reference/access-authn-authz/bootstrap-tokens/)

## Related

- [concepts.md](concepts.md) — bootstrap token concepts and TLS bootstrap
- [join.md](join.md) — how tokens are used during join
- [init.md](init.md) — token creation during init
