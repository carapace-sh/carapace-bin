# kubeadm join

The `kubeadm join` command bootstraps a new Kubernetes node — either a worker node or an additional control-plane node — and joins it to an existing cluster.

> **Source of truth**: <https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm-join/>. For the phase tree, see [phases.md](phases.md).

## Synopsis

```
kubeadm join [api-server-endpoint] [flags]
```

The `api-server-endpoint` is the address (and optional port) of the control plane API server. This is typically the `--control-plane-endpoint` or the advertise address from `kubeadm init`.

## Two Join Modes

| Mode | Flag | What happens |
|------|------|--------------|
| **Worker node** | (default) | Discovery + TLS bootstrap → kubelet gets permanent identity |
| **Control-plane node** | `--control-plane` | Above + download shared certs + generate control plane manifests + join etcd |

## Discovery

Before a node can join, it must discover cluster information and verify it's talking to the real control plane. kubeadm supports three discovery methods:

### 1. Token-based with CA pinning (default, strongest)

```bash
kubeadm join 10.0.0.1:6443 \
  --token abcdef.0123456789abcdef \
  --discovery-token-ca-cert-hash sha256:1234...
```

The joining node fetches the `cluster-info` ConfigMap from `kube-public`, validates its signature using the bootstrap token, then validates that the root CA public key matches the provided hash. This provides a strong out-of-band root of trust.

### 2. Token-based without CA pinning (weaker)

```bash
kubeadm join 10.0.0.1:6443 \
  --token abcdef.0123456789abcdef \
  --discovery-token-unsafe-skip-ca-verification
```

Skips CA hash verification. Relies only on the HMAC-SHA256 signature of the `cluster-info` ConfigMap. Use only in controlled environments — this is vulnerable to MITM attacks.

### 3. File/HTTPS-based

```bash
kubeadm join --discovery-file /path/to/cluster-info.yaml
kubeadm join --discovery-file https://internal.example.com/cluster-info.yaml
```

The discovery file is a kubeconfig containing the cluster CA and API server endpoint. This provides a strong out-of-band root of trust without token-based discovery.

### Discovery Flags

| Flag | Description |
|------|-------------|
| `--discovery-file` | File or URL for file-based discovery |
| `--discovery-token` | Token for token-based discovery |
| `--discovery-token-ca-cert-hash` | Validate root CA matches this hash (format: `sha256:<hex>`) |
| `--discovery-token-unsafe-skip-ca-verification` | Skip CA pinning (weaker security) |
| `--tls-bootstrap-token` | Token for TLS bootstrap authentication (if different from discovery token) |
| `--token` | Convenience: use same token for both discovery and TLS bootstrap |

## TLS Bootstrap

After discovery, the node performs TLS Bootstrap:

1. The kubelet connects to the API server using the bootstrap token (temporary auth)
2. The kubelet submits a CSR for its client certificate
3. The control plane auto-approves the CSR (kubeadm configured this RBAC during `init`)
4. The kubelet receives its permanent client certificate and switches to it

See [concepts.md](concepts.md) for the full TLS Bootstrap explanation.

## Worker Node Workflow

```
preflight
  → discovery (fetch + verify cluster-info)
    → kubelet-start (write bootstrap kubeconfig, start kubelet)
      → TLS bootstrap (CSR submitted, auto-approved, permanent cert received)
        → kubelet-wait-bootstrap
```

The node is now a functioning worker. No control plane components run on it.

## Control-Plane Node Workflow

When `--control-plane` is set, the join does everything a worker does, plus:

```
control-plane-prepare
  → download-certs (decrypt shared certs using --certificate-key)
  → certs (generate any remaining certs)
  → kubeconfig (generate control-plane kubeconfigs)
  → control-plane (generate static Pod manifests)
etcd-join (add new etcd member to the cluster)
control-plane-join
  → mark-control-plane
wait-control-plane
```

### `--certificate-key`

Required for control-plane join when certs were uploaded during `init` with `--upload-certs`. This key decrypts the `kubeadm-certs` Secret. The key expires 2 hours after `init` uploaded the certs.

```bash
kubeadm join 10.0.0.1:6443 \
  --token abcdef.0123456789abcdef \
  --discovery-token-ca-cert-hash sha256:1234... \
  --control-plane \
  --certificate-key 1234...
```

## Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--apiserver-advertise-address` | string | auto-detect | IP for API Server (control-plane join only). |
| `--apiserver-bind-port` | int32 | `6443` | Port for API Server (control-plane join only). |
| `--certificate-key` | string | — | Key to decrypt certificate secrets uploaded by `init`. |
| `--config` | string | — | Path to a kubeadm configuration file. |
| `--control-plane` | bool | `false` | Create a new control plane instance on this node. |
| `--cri-socket` | string | auto-detect | Path to the CRI socket. |
| `--discovery-file` | string | — | File or URL for file-based discovery. |
| `--discovery-token` | string | — | Token for token-based discovery. |
| `--discovery-token-ca-cert-hash` | stringSlice | — | Validate root CA matches this hash (`<type>:<value>`). |
| `--discovery-token-unsafe-skip-ca-verification` | bool | `false` | Skip CA pinning. |
| `--dry-run` | bool | `false` | Output what would be done without applying. |
| `--ignore-preflight-errors` | stringSlice | — | Checks to treat as warnings. See [flags.md](flags.md). |
| `--node-name` | string | hostname | Node name. |
| `--patches` | string | — | Directory with patch files. See [flags.md](flags.md). |
| `--skip-phases` | stringSlice | — | Phases to skip. See [phases.md](phases.md). |
| `--tls-bootstrap-token` | string | — | Token for TLS bootstrap (if different from discovery token). |
| `--token` | string | — | Convenience: same token for discovery and TLS bootstrap. |

### Inherited (global) flags

| Flag | Description |
|------|-------------|
| `--rootfs` | Path to the 'real' host root filesystem. |
| `--v` | Log level verbosity. |
| `--vmodule` | Comma-separated `pattern=N` for file-filtered logging. |

## Config File

Instead of flags, you can pass a `JoinConfiguration` (and optionally `ClusterConfiguration`) via `--config`:

```yaml
apiVersion: kubeadm.k8s.io/v1beta4
kind: JoinConfiguration
discovery:
  bootstrapToken:
    token: abcdef.0123456789abcdef
    apiServerEndpoint: 10.0.0.1:6443
    caCertHashes:
      - sha256:1234...
nodeRegistration:
  name: worker-1
  criSocket: unix:///var/run/containerd/containerd.sock
controlPlane:
  localAPIEndpoint:
    advertiseAddress: 10.0.0.2
    bindPort: 6443
```

See [config.md](config.md) for the full `JoinConfiguration` API.

## References

- [kubeadm join docs](https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm-join/)
- [TLS bootstrap](https://kubernetes.io/docs/reference/access-authn-authz/kubelet-tls-bootstrapping/)

## Related

- [init.md](init.md) — the control-plane bootstrap that produces the join command
- [phases.md](phases.md) — the join phase tree
- [token.md](token.md) — bootstrap token management
- [concepts.md](concepts.md) — TLS bootstrap and discovery concepts
