# kubeadm init

The `kubeadm init` command bootstraps a Kubernetes control-plane node. It runs a sequence of phases — from pre-flight checks through certificate generation, kubeconfig creation, static Pod manifest generation, and addon installation — to produce a minimum viable cluster.

> **Source of truth**: <https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm-init/>. For the phase tree, see [phases.md](phases.md).

## Synopsis

```
kubeadm init [flags]
```

`kubeadm init` takes no positional arguments. All configuration is via flags or a `--config` file (see [config.md](config.md)).

## What Happens During init

1. **Pre-flight checks** — validates system state (kernel, swap, ports, CRI, etc.)
2. **Certificate generation** — self-signed CA + all component certificates in `/etc/kubernetes/pki/`
3. **Kubeconfig generation** — writes `admin.conf`, `super-admin.conf`, `kubelet.conf`, `controller-manager.conf`, `scheduler.conf` to `/etc/kubernetes/`
4. **Static Pod manifests** — generates manifests for kube-apiserver, kube-controller-manager, kube-scheduler, and etcd (if stacked) to `/etc/kubernetes/manifests/`
5. **Labels and taints** — marks the node as control-plane with appropriate taints
6. **Bootstrap token** — generates a token (24h default TTL) for joining nodes
7. **TLS Bootstrap RBAC** — configures auto-approval for CSRs from bootstrap tokens
8. **Addons** — installs CoreDNS and kube-proxy via the API server

After completion, kubeadm prints the join command for both worker and control-plane nodes.

## Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--apiserver-advertise-address` | string | auto-detect | IP the API Server advertises. Uses default network interface if unset. |
| `--apiserver-bind-port` | int32 | `6443` | Port for the API Server to bind to. |
| `--apiserver-cert-extra-sans` | stringSlice | — | Extra SANs (IPs and DNS names) for the API Server serving cert. |
| `--cert-dir` | string | `/etc/kubernetes/pki` | Path to save/store certificates. |
| `--certificate-key` | string | — | Hex-encoded AES key (32 bytes) to encrypt control-plane certs in the kubeadm-certs Secret. |
| `--config` | string | — | Path to a kubeadm configuration file. |
| `--control-plane-endpoint` | string | — | Stable IP or DNS name for the control plane (use with load balancer for HA). |
| `--cri-socket` | string | auto-detect | Path to the CRI socket. Use only if multiple CRIs or non-standard socket. |
| `--dry-run` | bool | `false` | Output what would be done without applying changes. |
| `--feature-gates` | string | — | `key=value` pairs for feature gates. See [concepts.md](concepts.md). |
| `--ignore-preflight-errors` | stringSlice | — | Checks to treat as warnings. `all` ignores all. See [flags.md](flags.md). |
| `--image-repository` | string | `registry.k8s.io` | Container registry for control plane images. |
| `--kubernetes-version` | string | `stable-1` | Specific Kubernetes version for the control plane. |
| `--node-name` | string | hostname | Node name (passes `--hostname-override` to kubelet). |
| `--patches` | string | — | Directory with patch files for static Pod manifests. See [flags.md](flags.md). |
| `--pod-network-cidr` | string | — | CIDR range for the pod network. If set, control plane auto-allocates per-node CIDRs. |
| `--service-cidr` | string | `10.96.0.0/12` | CIDR range for service VIPs. |
| `--service-dns-domain` | string | `cluster.local` | Domain for services (e.g. `myorg.internal`). |
| `--skip-certificate-key-print` | bool | `false` | Don't print the certificate encryption key. |
| `--skip-phases` | stringSlice | — | List of phases to skip. See [phases.md](phases.md). |
| `--skip-token-print` | bool | `false` | Skip printing the default bootstrap token. |
| `--token` | string | auto-generated | Bootstrap token. Format: `[a-z0-9]{6}\.[a-z0-9]{16}`. |
| `--token-ttl` | duration | `24h0m0s` | Time before token auto-deletion. `0` = never expire. |
| `--upload-certs` | bool | `false` | Upload control-plane certificates to the kubeadm-certs Secret. |

### Inherited (global) flags

| Flag | Description |
|------|-------------|
| `--rootfs` | Path to the 'real' host root filesystem (kubeadm chroots into it). |
| `--v` | Log level verbosity. |
| `--vmodule` | Comma-separated `pattern=N` for file-filtered logging. |

## Networking

| Aspect | Flag | Default |
|--------|------|---------|
| API Server address | `--apiserver-advertise-address` | auto-detected from network interface |
| API Server port | `--apiserver-bind-port` | `6443` |
| Pod network CIDR | `--pod-network-cidr` | — (must be set to match your CNI) |
| Service CIDR | `--service-cidr` | `10.96.0.0/12` |
| Service DNS domain | `--service-dns-domain` | `cluster.local` |
| Control plane endpoint | `--control-plane-endpoint` | — (use for HA / load balancer) |

**Important**: `--pod-network-cidr` must match the CIDR your CNI plugin uses. kubeadm does not install a CNI — you must do this after `init`. CoreDNS will not be scheduled until a CNI is installed.

## Certificate Upload (`--upload-certs`)

When `--upload-certs` is passed, kubeadm encrypts the control-plane certificates and uploads them to a Secret named `kubeadm-certs` in `kube-system`. This Secret **expires after 2 hours**. The `--certificate-key` encrypts the certs; generate one with:

```bash
kubeadm certs certificate-key
```

This enables the **control-plane join** workflow — additional control-plane nodes use `--certificate-key` during `kubeadm join --control-plane` to download and decrypt the shared certificates. See [join.md](join.md).

## Certificate Key

The `--certificate-key` is a hex-encoded string representing an AES key of 32 bytes (256 bits). It is used to encrypt/decrypt the control-plane certificates stored in the `kubeadm-certs` Secret. Generate a new one with `kubeadm certs certificate-key`.

## Custom Images

To use a private registry or air-gapped environment:

```bash
# List images kubeadm will use
kubeadm config images list --image-repository my.registry.com

# Pre-pull images
kubeadm config images pull --image-repository my.registry.com
```

See [config.md](config.md) for the `config images` subcommands.

## Providing Your Own CA

To use a custom CA instead of the auto-generated one, place the CA cert and key in the cert directory before running `init`:

```
/etc/kubernetes/pki/ca.crt
/etc/kubernetes/pki/ca.key
```

kubeadm detects existing CA files and uses them, generating only the component certificates signed by your CA. The same applies to the front-proxy CA and etcd CA (in `/etc/kubernetes/pki/etcd/`).

For the full certificate inventory, see [certs.md](certs.md).

## Dry Run

`--dry-run` outputs what would be done without applying any changes. Useful for validating a configuration before committing to it.

## Skip Phases

`--skip-phases` accepts a comma-separated list of phase paths to skip:

```bash
kubeadm init --skip-phases=addon/kube-proxy
```

Phase paths use `/` to denote sub-phases (e.g. `certs/ca`, `addon/coredns`). See [phases.md](phases.md) for the complete phase tree.

## Automation

For scripted deployments:

```bash
# Generate a token in advance
TOKEN=$(kubeadm token generate)

# Generate a certificate key
CERT_KEY=$(kubeadm certs certificate-key)

# Init with the pre-generated values
kubeadm init --token="$TOKEN" --certificate-key="$CERT_KEY" --upload-certs
```

See [token.md](token.md) for `kubeadm token generate`.

## References

- [kubeadm init docs](https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm-init/)
- [Customizing control plane flags](https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/control-plane-flags/)

## Related

- [phases.md](phases.md) — the init phase tree and `phase` subcommand
- [join.md](join.md) — joining nodes to the cluster
- [certs.md](certs.md) — certificate inventory and renewal
- [config.md](config.md) — the `--config` file API
