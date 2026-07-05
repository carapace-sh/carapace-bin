# kubeadm upgrade — Cluster Upgrades

The `kubeadm upgrade` command upgrades a Kubernetes cluster to a newer version. It wraps complex upgrade logic behind four subcommands.

> **Source of truth**: <https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm-upgrade/>. For the phase tree, see [phases.md](phases.md).

## Subcommands

| Subcommand | Purpose |
|------------|---------|
| `kubeadm upgrade plan [version]` | Check available upgrade versions and validate upgradeability |
| `kubeadm upgrade apply [version]` | Upgrade the cluster to the specified version |
| `kubeadm upgrade diff [version]` | Show differences that would be applied to static Pod manifests |
| `kubeadm upgrade node` | Upgrade a specific node in the cluster |

## Upgrade Process

The standard upgrade workflow is:

1. **On the first control-plane node**: `kubeadm upgrade plan` → `kubeadm upgrade apply <version>`
2. **On remaining control-plane nodes**: `kubeadm upgrade node`
3. **On worker nodes**: `kubeadm upgrade node`
4. **On every node**: upgrade kubelet and kubectl packages, then restart kubelet

kubeadm only upgrades the control plane components and kubelet config — **not** the kubelet binary itself (that's done via your package manager).

## `upgrade plan`

Checks which versions are available to upgrade to and validates whether your current cluster is upgradeable.

```bash
kubeadm upgrade plan
kubeadm upgrade plan v1.30.0
```

### Flags

| Flag | Description |
|------|-------------|
| `--allow-experimental-upgrades` | Show alpha/beta/RC versions |
| `--allow-missing-template-keys` | Ignore missing template fields |
| `--allow-release-candidate-upgrades` | Show RC versions |
| `--config` | Path to a kubeadm config file |
| `--etcd-upgrade` | Whether to upgrade etcd (default: true) |
| `--ignore-preflight-errors` | Checks to treat as warnings |
| `--kubeconfig` | Kubeconfig file (default: `/etc/kubernetes/admin.conf`) |
| `-o, --output` | Output format |
| `--print-config` | Print the config that will be used |
| `--show-managed-fields` | Keep managedFields in JSON/YAML output |

### Output

Shows the current cluster version, available upgrade targets, and any issues that would block an upgrade.

## `upgrade apply`

Upgrades the cluster to the specified version. Run this on the **first control-plane node** only.

```bash
kubeadm upgrade apply v1.30.0
kubeadm upgrade apply v1.30.0 --yes
```

### Flags

| Flag | Description |
|------|-------------|
| `--allow-experimental-upgrades` | Allow alpha/beta/RC versions |
| `--allow-release-candidate-upgrades` | Allow RC versions |
| `--certificate-renewal` | Renew certificates during upgrade (default: true) |
| `--config` | Path to a kubeadm config file |
| `--dry-run` | Output actions without changing state |
| `--etcd-upgrade` | Perform etcd upgrade (default: true) |
| `--force` | Force upgrade non-interactively (implies non-interactive) |
| `--ignore-preflight-errors` | Checks to treat as warnings |
| `--kubeconfig` | Kubeconfig file |
| `--patches` | Directory with patch files for static Pod manifests |
| `--print-config` | Print the config that will be used |
| `--skip-phases` | Phases to skip |
| `-y, --yes` | Non-interactive mode (skip confirmation prompt) |

### apply Phase Workflow

```
preflight          Run preflight checks
control-plane      Upgrade the control plane components
upload-config      Upload kubeadm and kubelet configs to ConfigMaps
  /kubeadm           Upload ClusterConfiguration
  /kubelet           Upload kubelet configuration
kubelet-config     Upgrade kubelet configuration for this node
bootstrap-token    Configure bootstrap token and cluster-info RBAC
addon              Upgrade default addons
  /coredns           Upgrade CoreDNS
  /kube-proxy        Upgrade kube-proxy
post-upgrade       Run post-upgrade tasks
```

## `upgrade node`

Upgrades a node in the cluster. Run this on **additional control-plane nodes** and **worker nodes** after `upgrade apply` has been run on the first control-plane node.

```bash
kubeadm upgrade node
```

### Flags

| Flag | Description |
|------|-------------|
| `--certificate-renewal` | Renew certificates during upgrade (default: true) |
| `--config` | Path to a kubeadm config file |
| `--dry-run` | Output actions without changing state |
| `--etcd-upgrade` | Perform etcd upgrade (default: true) |
| `--ignore-preflight-errors` | Checks to treat as warnings |
| `--kubeconfig` | Kubeconfig file |
| `--patches` | Directory with patch files |
| `--skip-phases` | Phases to skip |

### node Phase Workflow

```
preflight       Run upgrade node pre-flight checks
control-plane   Upgrade control plane instance on this node (if any)
kubelet-config  Upgrade kubelet configuration for this node
addon           Upgrade default addons
  /coredns        Upgrade CoreDNS
  /kube-proxy     Upgrade kube-proxy
post-upgrade    Run post-upgrade tasks
```

### node Phase Sub-phases

The `upgrade node phase` subcommand exposes individual phases:

| Phase | Description |
|-------|-------------|
| `preflight` | Run preflight checks |
| `control-plane` | Upgrade control plane instance |
| `addon` | Upgrade addons |
| `addon all` | Upgrade all addons |
| `addon coredns` | Upgrade CoreDNS |
| `addon kube-proxy` | Upgrade kube-proxy |
| `post-upgrade` | Run post-upgrade tasks |

## `upgrade diff`

Shows what differences would be applied to existing static Pod manifests. See also `kubeadm upgrade apply --dry-run`.

```bash
kubeadm upgrade diff v1.30.0
kubeadm upgrade diff v1.30.0 --context-lines 5
```

### Flags

| Flag | Description |
|------|-------------|
| `--config` | Path to a kubeadm config file |
| `-c, --context-lines` | Number of context lines in the diff |
| `--kubeconfig` | Kubeconfig file |

## Certificate Renewal During Upgrade

`--certificate-renewal` (default: `true`) renews all control-plane certificates during the upgrade. This is convenient but means the kubeconfig files (`admin.conf`, etc.) get new embedded certs — you must update your local `~/.kube/config` afterward.

To skip cert renewal during upgrade:

```bash
kubeadm upgrade apply v1.30.0 --certificate-renewal=false
```

For standalone cert renewal, see [certs.md](certs.md).

## etcd Upgrade

`--etcd-upgrade` (default: `true`) upgrades etcd as part of the cluster upgrade. This updates the etcd container image in the static Pod manifest. For stacked etcd, this happens on each control-plane node. For external etcd, this flag has no effect (etcd is managed externally).

## Skip Phases

`--skip-phases` works the same as other commands — see [phases.md](phases.md).

```bash
kubeadm upgrade apply v1.30.0 --skip-phases=addon/kube-proxy
```

## Edge Cases

- **Version skew**: kubeadm must be the same version as the target Kubernetes version (±1 minor). Upgrade kubeadm itself first via your package manager.
- **HA clusters**: Run `upgrade apply` on the first control-plane node, then `upgrade node` on each remaining control-plane node, then `upgrade node` on each worker node.
- **Rollback**: kubeadm does not support automatic rollback. If an upgrade fails, you must manually revert the static Pod manifests and kubelet config.
- **`--force`**: Bypasses some preflight checks and implies non-interactive mode. Use with caution.
- **Dry run**: `--dry-run` shows what would happen without applying changes. `upgrade diff` is a more focused alternative for just the manifest changes.

## References

- [kubeadm upgrade docs](https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm-upgrade/)
- [Upgrading kubeadm clusters](https://kubernetes.io/docs/tasks/administer-cluster/kubeadm/kubeadm-upgrade/)

## Related

- [phases.md](phases.md) — upgrade phase trees
- [certs.md](certs.md) — certificate renewal (used by `--certificate-renewal`)
- [config.md](config.md) — `UpgradeConfiguration` in v1beta4
