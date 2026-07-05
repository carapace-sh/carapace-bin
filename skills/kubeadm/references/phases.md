# The Phase System

kubeadm executes each command (`init`, `join`, `reset`, `upgrade apply`, `upgrade node`) as a sequence of **phases**. Each phase is a discrete step that can be invoked individually via the `phase` subcommand, skipped via `--skip-phases`, or composed into custom workflows.

> **Source of truth**: <https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm-init-phase/> and the per-command phase docs. For the completer's phase action, see `completers/common/kubeadm_completer/cmd/action/phase.go`.

## How Phases Work

Every phaseable command (`init`, `join`, `reset`, `upgrade apply`, `upgrade node`) has a `phase` subcommand that exposes the individual phases. Phases are hierarchical — a phase like `certs` contains sub-phases like `certs/ca`, `certs/apiserver`, etc.

```
kubeadm init phase certs ca
kubeadm init phase certs apiserver
kubeadm init phase kubeconfig admin
```

### `--skip-phases`

The `--skip-phases` flag accepts a comma-separated list of phase paths to skip when running the full command:

```bash
kubeadm init --skip-phases=addon/kube-proxy,bootstrap-token
```

Phase paths use `/` to denote the hierarchy. Skipping a parent phase skips all its sub-phases.

### Phase Ordering

Phases run in a fixed order. Running a parent phase runs its sub-phases in order. You can run individual sub-phases out of order via the `phase` subcommand, but doing so may produce an inconsistent state — phases have dependencies on earlier phases (e.g. `control-plane` depends on `certs` and `kubeconfig`).

## init Phase Tree

```
preflight                         Run pre-flight checks
certs                             Certificate generation
  /ca                               Generate self-signed Kubernetes CA
  /apiserver                        Generate cert for serving the Kubernetes API
  /apiserver-kubelet-client         Generate cert for API server → kubelet
  /front-proxy-ca                   Generate self-signed CA for front proxy
  /front-proxy-client               Generate cert for front proxy client
  /etcd-ca                          Generate self-signed CA for etcd
  /etcd-server                      Generate cert for serving etcd
  /etcd-peer                        Generate cert for etcd peer communication
  /etcd-healthcheck-client          Generate cert for etcd liveness probes
  /apiserver-etcd-client            Generate cert for apiserver → etcd
  /sa                               Generate service account signing key pair
kubeconfig                        Generate all kubeconfig files
  /admin                            Generate admin kubeconfig
  /super-admin                      Generate super-admin kubeconfig (bypasses RBAC)
  /kubelet                          Generate kubelet bootstrap kubeconfig
  /controller-manager               Generate controller-manager kubeconfig
  /scheduler                        Generate scheduler kubeconfig
etcd                              Generate static Pod manifest for local etcd
  /local                            Generate static Pod manifest for single-node etcd
control-plane                     Generate static Pod manifests for control plane
  /apiserver                        kube-apiserver static Pod manifest
  /controller-manager               kube-controller-manager static Pod manifest
  /scheduler                        kube-scheduler static Pod manifest
kubelet-start                     Write kubelet settings and (re)start kubelet
wait-control-plane                Wait for the control plane to start up
upload-config                     Upload kubeadm and kubelet config to ConfigMaps
  /kubeadm                          Upload ClusterConfiguration to a ConfigMap
  /kubelet                          Upload kubelet component config to a ConfigMap
upload-certs                      Upload certificates to kubeadm-certs Secret
mark-control-plane                Mark node as control-plane (labels + taints)
bootstrap-token                   Generate bootstrap tokens for joining nodes
kubelet-finalize                  Update kubelet settings after TLS bootstrap
  /enable-client-cert-rotation      Enable kubelet client certificate rotation
addon                             Install required addons
  /coredns                          Install CoreDNS addon
  /kube-proxy                       Install kube-proxy addon
show-join-command                 Show the join command for control-plane and worker nodes
```

## join Phase Tree

```
preflight                         Run join pre-flight checks
control-plane-prepare             Prepare machine for serving a control plane
  /download-certs                   Download certs shared among control-plane nodes
  /certs                            Generate certs for new control-plane components
  /kubeconfig                       Generate kubeconfig for new control-plane components
  /control-plane                    Generate manifests for new control-plane components
kubelet-start                     Write kubelet settings, certs, and (re)start kubelet
etcd-join                         Join etcd for control-plane nodes
kubelet-wait-bootstrap            Wait for kubelet to bootstrap itself
control-plane-join                Join machine as a control plane instance
  /mark-control-plane               Mark node as control-plane
wait-control-plane                Wait for control plane to start (control-plane join only)
```

The `control-plane-prepare`, `control-plane-join`, `etcd-join`, and `wait-control-plane` phases only run when `--control-plane` is set. Worker node joins run `preflight`, `kubelet-start`, and `kubelet-wait-bootstrap`.

## reset Phase Tree

```
preflight                         Run reset pre-flight checks
remove-etcd-member                Remove a local etcd member (control-plane nodes)
cleanup-node                      Run cleanup node (remove certs, configs, manifests)
```

See [reset.md](reset.md).

## upgrade apply Phase Tree

```
preflight                         Run preflight checks before upgrade
control-plane                     Upgrade the control plane components
upload-config                     Upload kubeadm and kubelet configurations
  /kubeadm                          Upload ClusterConfiguration to a ConfigMap
  /kubelet                          Upload kubelet configuration to a ConfigMap
kubelet-config                    Upgrade the kubelet configuration for this node
bootstrap-token                   Configure bootstrap token and cluster-info RBAC
addon                             Upgrade default kubeadm addons
  /coredns                          Upgrade the CoreDNS addon
  /kube-proxy                       Upgrade the kube-proxy addon
post-upgrade                      Run post-upgrade tasks
```

## upgrade node Phase Tree

```
preflight                         Run upgrade node pre-flight checks
control-plane                     Upgrade control plane instance on this node (if any)
kubelet-config                    Upgrade the kubelet configuration for this node
addon                             Upgrade default kubeadm addons
  /coredns                          Upgrade the CoreDNS addon
  /kube-proxy                       Upgrade the kube-proxy addon
post-upgrade                      Run post-upgrade tasks
```

See [upgrade.md](upgrade.md).

## Invoking Individual Phases

```bash
# Run just the certificate generation
kubeadm init phase certs all

# Generate only the CA
kubeadm init phase certs ca

# Generate kubeconfig files
kubeadm init phase kubeconfig all

# Upload certs after init (for adding control-plane nodes later)
kubeadm init phase upload-certs --upload-certs --certificate-key <key>

# Run a single join phase
kubeadm join phase preflight 10.0.0.1:6443 --token abcdef.0123456789abcdef
```

### `all` sub-phase

Most parent phases accept `all` as a sub-phase to run all sub-phases in order:

```bash
kubeadm init phase certs all
kubeadm init phase kubeconfig all
kubeadm init phase control-plane all
```

## Building Custom Workflows

The phase system lets you build custom workflows by running phases individually. Common patterns:

### Generate certs only (for external CA mode)

```bash
kubeadm init phase certs all --cert-dir /path/to/certs
```

### Generate manifests without starting the cluster

```bash
kubeadm init phase control-plane all --dry-run
```

### Re-upload certs after expiry

```bash
kubeadm init phase upload-certs --upload-certs --certificate-key $(kubeadm certs certificate-key)
```

### Re-run bootstrap token setup

```bash
kubeadm init phase bootstrap-token
```

## Edge Cases

- **Out-of-order phases**: Running `control-plane` before `certs` will fail because the manifests reference certificates that don't exist yet.
- **`--skip-phases` with parent paths**: Skipping `certs` skips the entire `certs` tree. To skip only one cert, use the full path like `certs/etcd-server`.
- **Phase flags**: Individual phase invocations accept the same flags as the parent command (e.g. `--config`, `--cert-dir`, `--dry-run`).
- **`show-join-command`**: The last phase of `init` that prints the join command for both control-plane and worker nodes. Not listed in `--skip-phases` completion.

## References

- [kubeadm init phase](https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm-init-phase/)
- [kubeadm join phase](https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm-join-phase/)

## Related

- [init.md](init.md) — the init command and its flags
- [join.md](join.md) — the join command and discovery
- [reset.md](reset.md) — the reset command
- [upgrade.md](upgrade.md) — the upgrade command
