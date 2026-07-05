# kubeadm reset — Teardown

The `kubeadm reset` command performs a best-effort revert of changes made to a host by `kubeadm init` or `kubeadm join`. It cleans up certificates, kubeconfig files, static Pod manifests, and removes the node from the cluster (including etcd membership for control-plane nodes).

> **Source of truth**: <https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm-reset/>. For the phase tree, see [phases.md](phases.md).

## Synopsis

```
kubeadm reset [flags]
```

## What reset Does

1. **Preflight checks** — validates state before cleanup
2. **Remove etcd member** (control-plane nodes only) — removes the local etcd member from the cluster
3. **Cleanup node** — removes:
   - `/etc/kubernetes/manifests/` (static Pod manifests)
   - `/etc/kubernetes/pki/` (certificates)
   - `/etc/kubernetes/*.conf` (kubeconfig files)
   - kubelet configuration
   - CRI containers and volumes created by kubeadm

## What reset Does NOT Do

- Clean CNI configuration (`/etc/cni/net.d`)
- Remove iptables / nftables / IPVS rules
- Remove `$HOME/.kube/config`
- Delete external etcd data (only affects stacked etcd)

You must clean these up manually if needed.

## Phase Tree

```
preflight             Run reset pre-flight checks
remove-etcd-member    Remove a local etcd member (control-plane nodes only)
cleanup-node          Run cleanup node (remove certs, configs, manifests)
```

See [phases.md](phases.md) for invoking individual phases.

## Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--cert-dir` | `/etc/kubernetes/pki` | Path to certificates directory to clean. |
| `--cleanup-tmp-dir` | `false` | Cleanup `/etc/kubernetes/tmp` directory. |
| `--config` | — | Path to a kubeadm configuration file (v1beta4 `ResetConfiguration`). |
| `--cri-socket` | auto-detect | Path to the CRI socket. |
| `--dry-run` | `false` | Output what would be done without applying. |
| `-f, --force` | `false` | Reset without prompting for confirmation. |
| `--ignore-preflight-errors` | — | Checks to treat as warnings. See [flags.md](flags.md). |
| `--kubeconfig` | `/etc/kubernetes/admin.conf` | Kubeconfig file to use for talking to the cluster. |
| `--skip-phases` | — | Phases to skip. See [phases.md](phases.md). |

### Inherited (global) flags

| Flag | Description |
|------|-------------|
| `--rootfs` | Path to the 'real' host root filesystem. |
| `--v` | Log level verbosity. |
| `--vmodule` | Comma-separated `pattern=N` for file-filtered logging. |

## `--force`

Skips the interactive confirmation prompt. Useful for scripting:

```bash
kubeadm reset --force
```

## `--cleanup-tmp-dir`

Also removes `/etc/kubernetes/tmp`, which contains temporary working files from the last `init` or `join`. Without this flag, the tmp directory is left intact.

## Removing etcd Membership

For control-plane nodes with stacked etcd, the `remove-etcd-member` phase removes the local etcd member from the cluster before cleaning up. This prevents the cluster from having a dead etcd member.

If the etcd cluster has already lost this member (e.g. the node is being reset after a failure), you may need to skip this phase:

```bash
kubeadm reset --skip-phases=remove-etcd-member
```

## Config File (v1beta4)

In v1beta4, `kubeadm reset` accepts a `ResetConfiguration` via `--config`:

```yaml
apiVersion: kubeadm.k8s.io/v1beta4
kind: ResetConfiguration
cleanupTmpDir: true
certificatesDir: /etc/kubernetes/pki
force: true
ignorePreflightErrors:
  - all
```

See [config.md](config.md) for the full `ResetConfiguration` API.

## Edge Cases

- **Stale etcd members**: If a control-plane node is already down and you're resetting it, the etcd member may already be removed. Use `--skip-phases=remove-etcd-member`.
- **CNI cleanup**: `reset` does not remove CNI config or network rules. Manually clean `/etc/cni/net.d` and flush iptables/nftables if needed.
- **`$HOME/.kube/config`**: Not touched by reset. Remove manually if you want to clear the user's kubeconfig.
- **External etcd**: `reset` does not delete external etcd data. Only stacked etcd data (in the static Pod volume) is cleaned.
- **Dry run**: `--dry-run` shows what would be cleaned without removing anything.

## References

- [kubeadm reset docs](https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm-reset/)

## Related

- [phases.md](phases.md) — reset phase tree
- [config.md](config.md) — `ResetConfiguration` in v1beta4
- [flags.md](flags.md) — common flags like `--ignore-preflight-errors`
