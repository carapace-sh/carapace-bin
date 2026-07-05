# Core Concepts

The mental model behind kubeadm — what it bootstraps, how the control plane and etcd fit together, how nodes establish trust, and the key abstractions a developer needs to reason about cluster lifecycle.

> **Source of truth**: <https://kubernetes.io/docs/reference/setup-tools/kubeadm/>. For the completer implementation, see `completers/common/kubeadm_completer/` in this repo.

## What kubeadm Is

kubeadm provides `kubeadm init` and `kubeadm join` as best-practice "fast paths" for creating Kubernetes clusters. It performs only the actions necessary to get a **minimum viable cluster** up and running. By design it cares about **bootstrapping**, not about provisioning machines or installing addons like dashboards, monitoring, or cloud-specific plugins.

### What kubeadm Does

- Generates the cluster PKI (CA + component certificates)
- Writes kubeconfig files for the control plane components and admin
- Generates static Pod manifests for the control plane (and local etcd)
- Configures bootstrap tokens and TLS bootstrap RBAC
- Installs CoreDNS and kube-proxy addons

### What kubeadm Does Not Do

- Provision machines or configure networks
- Install a CNI plugin (the user must do this after `init`)
- Install dashboards, monitoring, or cloud-provider addons
- Manage certificate renewal on a schedule (use `kubeadm certs renew` manually or via automation)

## The Control Plane

The control plane is established via **static Pod manifests** written to `/etc/kubernetes/manifests/`. The kubelet watches this directory and creates the Pods on startup. The three core control plane components:

| Component | Manifest | Purpose |
|-----------|----------|---------|
| **kube-apiserver** | `kube-apiserver.yaml` | Exposes the Kubernetes API |
| **kube-controller-manager** | `kube-controller-manager.yaml` | Runs core controllers |
| **kube-scheduler** | `kube-scheduler.yaml` | Schedules Pods onto nodes |

Custom flags for these components are passed via the kubeadm config API (see [config.md](config.md)), not directly through kubeadm CLI flags.

### `--control-plane-endpoint`

A stable IP address or DNS name for the control plane. This abstracts the API server address behind a single endpoint — essential for high-availability clusters using a load balancer. Without it, the advertise address of the first control-plane node is used.

## etcd

kubeadm supports two etcd topologies:

| Mode | How it runs | When to use |
|------|-------------|-------------|
| **Stacked (local)** | Static Pod on the control-plane node (default) | Simple, single-node or small HA clusters |
| **External** | User-provided external etcd endpoints | When etcd is managed separately |

With stacked etcd, `kubeadm init` generates a static Pod manifest for a single-node local etcd instance (`init phase etcd local`). When joining additional control-plane nodes, kubeadm adds a new etcd member to the existing cluster.

With external etcd, the user provides endpoints and CA/cert paths in `ClusterConfiguration.etcd.external`; kubeadm skips the etcd static Pod manifest entirely.

For certificate details, see [certs.md](certs.md).

## Bootstrap Tokens

Bootstrap tokens are simple bearer tokens used to establish bidirectional trust between a joining node and the control plane. They are stable since Kubernetes v1.18.

### Token Format

```
abcdef.0123456789abcdef
```

- Regex: `[a-z0-9]{6}\.[a-z0-9]{16}`
- **Token ID** (first 6 chars) — public, used to reference the token
- **Token Secret** (last 16 chars) — private, shared only with trusted parties

### Storage and Identity

Tokens are stored as Secrets of type `bootstrap.kubernetes.io/token` in the `kube-system` namespace. They authenticate as `system:bootstrap:<token-id>` in the `system:bootstrappers` group.

### TTL and Usages

| Property | Default | Notes |
|----------|---------|-------|
| TTL | `24h` | Auto-deleted after expiry. Set `0` for no expiry. |
| Usages | `signing,authentication` | `signing` for ConfigMap signing, `authentication` for bearer auth |
| Groups | `system:bootstrappers:kubeadm:default-node-token` | Extra groups for authentication |

For token management commands, see [token.md](token.md).

### The `cluster-info` ConfigMap

The `cluster-info` ConfigMap in the `kube-public` namespace is signed by bootstrap tokens. This is what `kubeadm join` fetches during discovery — the signature lets the joining node verify the information came from the real control plane before it has the CA certificate.

## TLS Bootstrap

TLS Bootstrap is the mechanism by which a joining node obtains its definitive kubelet identity:

1. The joining node connects to the API server using the bootstrap token (temporary auth)
2. The kubelet submits a Certificate Signing Request (CSR)
3. The control plane auto-approves the CSR (kubeadm configures the RBAC for this)
4. The kubelet receives its client certificate and switches to its permanent identity

This means bootstrap tokens are **short-lived credentials** that bridge the gap between "node knows nothing" and "node has its own cert."

For the join workflow, see [join.md](join.md).

## Kubeconfig Files

kubeadm generates these kubeconfig files in `/etc/kubernetes/` during `init`:

| File | Purpose |
|------|---------|
| `admin.conf` | For the admin user and kubeadm itself |
| `super-admin.conf` | For a super-admin that can bypass RBAC |
| `kubelet.conf` | For the kubelet (cluster bootstrapping only) |
| `controller-manager.conf` | For the controller manager |
| `scheduler.conf` | For the scheduler |

The `kubelet.conf` generated by `init` is a **bootstrap** kubeconfig — it uses the bootstrap token. After TLS bootstrap completes, the kubelet switches to its own client certificate. When joining a node, the join workflow generates a bootstrap kubeconfig then completes TLS bootstrap to get the permanent `kubelet.conf`.

## Feature Gates

kubeadm supports feature gates via `--feature-gates key=value` on the CLI or `featureGates` in the config API.

### Current Feature Gates

| Gate | Default | Status | Description |
|------|---------|--------|-------------|
| `NodeLocalCRISocket` | `true` | GA 1.36 | kubeadm reads/writes CRI socket from `/var/lib/kubelet/instance-config.yaml` instead of Node annotations |

### Deprecated Feature Gates

| Gate | Default | Deprecated | Replacement |
|------|---------|------------|-------------|
| `PublicKeysECDSA` | `false` | 1.31 | `encryptionAlgorithm` field in v1beta4 |
| `RootlessControlPlane` | `false` | 1.31 | — |

Feature gates follow the standard Kubernetes lifecycle: Alpha → Beta → GA → Removal. Once GA, the gate is locked to `true` and the flag becomes a no-op. Removed gates (e.g. `EtcdLearnerMode`, `IPv6DualStack`, `WaitForAllControlPlaneComponents`) no longer have any effect.

## Filesystem Layout

After `kubeadm init`, the control-plane node has this structure:

```
/etc/kubernetes/
├── admin.conf              # admin kubeconfig
├── super-admin.conf        # super-admin kubeconfig (bypasses RBAC)
├── kubelet.conf            # kubelet kubeconfig
├── controller-manager.conf
├── scheduler.conf
├── manifests/              # static Pod manifests (kubelet watches this)
│   ├── kube-apiserver.yaml
│   ├── kube-controller-manager.yaml
│   ├── kube-scheduler.yaml
│   └── etcd.yaml           # only with stacked etcd
├── pki/                    # cluster PKI
│   ├── ca.crt / ca.key
│   ├── apiserver.crt / .key
│   ├── apiserver-kubelet-client.crt / .key
│   ├── front-proxy-ca.crt / .key
│   ├── front-proxy-client.crt / .key
│   ├── etcd/ca.crt / .key
│   ├── etcd/server.crt / .key
│   ├── etcd/peer.crt / .key
│   ├── etcd/healthcheck-client.crt / .key
│   ├── apiserver-etcd-client.crt / .key
│   └── sa.key / sa.pub     # service account signing key pair
└── tmp/                    # temporary working files
```

## References

- [kubeadm CLI overview](https://kubernetes.io/docs/reference/setup-tools/kubeadm/)
- [Bootstrap tokens](https://kubernetes.io/docs/reference/access-authn-authz/bootstrap-tokens/)
- [TLS bootstrap](https://kubernetes.io/docs/reference/access-authn-authz/kubelet-tls-bootstrapping/)

## Related

- [config.md](config.md) — the v1beta4 configuration API
- [certs.md](certs.md) — certificate inventory and renewal
- [token.md](token.md) — bootstrap token management commands
