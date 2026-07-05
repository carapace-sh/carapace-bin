# kubeadm kubeconfig, version, alpha

The remaining kubeadm commands: `kubeconfig` for generating kubeconfig files for additional users, `version` for printing the kubeadm version, and `alpha` for experimental features.

> **Source of truth**: <https://kubernetes.io/docs/reference/setup-tools/kubeadm/>.

## kubeadm kubeconfig

Kubeconfig file utilities.

### Subcommands

| Subcommand | Purpose |
|------------|---------|
| `kubeadm kubeconfig user` | Output a kubeconfig file for an additional user |

### `kubeconfig user`

Generates a kubeconfig file for an additional user. The kubeconfig is printed to stdout — redirect to a file to save it.

```bash
kubeadm kubeconfig user --client-name dev-user > dev-user.conf
```

#### Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--client-name` | (required) | Name of the user (used as the CN if client certs are created). |
| `--config` | — | Path to a kubeadm configuration file. |
| `--org` | — | Organizations of the client certificate (used as the O). Can be repeated. |
| `--token` | — | Use token authentication instead of client certificates. |
| `--validity-period` | `8760h0m0s` (1 year) | Validity period of the client certificate (offset from now). |

#### Examples

Generate a kubeconfig with a client certificate:

```bash
kubeadm kubeconfig user --client-name dev-user --org dev-team
```

Generate a kubeconfig with token authentication instead of a cert:

```bash
kubeadm kubeconfig user --client-name ci-bot --token abcdef.0123456789abcdef
```

Set a custom validity period:

```bash
kubeadm kubeconfig user --client-name temp-user --validity-period 720h
```

#### How It Works

The command connects to the cluster using the admin kubeconfig (`/etc/kubernetes/admin.conf`), generates a client certificate signed by the cluster CA, and embeds it in a new kubeconfig. The `--token` option skips cert generation and uses a bootstrap token instead.

## kubeadm version

Prints the version of kubeadm.

```bash
$ kubeadm version
kubeadm version: &version.Info{Major:"1", Minor:"30", GitVersion:"v1.30.0", ...}
```

### Flags

| Flag | Description |
|------|-------------|
| `-o, --output` | Output format: `yaml`, `json`, or `short` |

```bash
$ kubeadm version -o short
v1.30.0

$ kubeadm version -o yaml
gitVersion: v1.30.0
gitCommit: abc123...
gitTreeState: clean
buildDate: "2024-04-18T..."
goVersion: go1.22.2
compiler: gc
platform: linux/amd64
```

## kubeadm alpha

Preview features for gathering community feedback. This is a placeholder for experimental sub-commands.

Currently, there are **no experimental commands** under `kubeadm alpha`. Features that were previously here have either graduated to stable or been removed. The `alpha` command exists to provide a staging ground for future experimental features.

## References

- [kubeadm kubeconfig](https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm-kubeconfig/)
- [kubeadm version](https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm-version/)
- [kubeadm alpha](https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm-alpha/)

## Related

- [concepts.md](concepts.md) — kubeconfig files generated during init
- [certs.md](certs.md) — the CA that signs client certs for `kubeconfig user`
