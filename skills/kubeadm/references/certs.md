# kubeadm certs — Certificate Management

The `kubeadm certs` command (aliased as `kubeadm certificates`) manages the Kubernetes cluster PKI: renewing certificates, generating certificate keys, checking expiration, and generating CSRs for external CA mode.

> **Source of truth**: <https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm-certs/>.

## Subcommands

| Subcommand | Purpose |
|------------|---------|
| `kubeadm certs renew` | Renew certificates (parent) |
| `kubeadm certs renew <cert>` | Renew a specific certificate |
| `kubeadm certs renew all` | Renew all known control-plane certificates |
| `kubeadm certs certificate-key` | Generate a new certificate key |
| `kubeadm certs check-expiration` | Check certificate expiration |
| `kubeadm certs generate-csr` | Generate keys and CSRs (external CA mode) |

## Certificate Inventory

kubeadm generates these certificates during `init` (stored in `/etc/kubernetes/pki`):

| Certificate | File | Purpose |
|-------------|------|---------|
| Kubernetes CA | `ca.crt` / `ca.key` | Self-signed CA for all K8s components |
| API Server | `apiserver.crt` / `apiserver.key` | Serving cert for the Kubernetes API |
| API Server-Kubelet Client | `apiserver-kubelet-client.crt` / `.key` | API server → kubelet connection |
| Front Proxy CA | `front-proxy-ca.crt` / `front-proxy-ca.key` | Self-signed CA for front proxy |
| Front Proxy Client | `front-proxy-client.crt` / `front-proxy-client.key` | Front proxy client cert |
| etcd CA | `etcd/ca.crt` / `etcd/ca.key` | Self-signed CA for etcd |
| etcd Server | `etcd/server.crt` / `etcd/server.key` | Serving cert for etcd |
| etcd Peer | `etcd/peer.crt` / `etcd/peer.key` | etcd node-to-node communication |
| etcd Healthcheck Client | `etcd/healthcheck-client.crt` / `.key` | Liveness probes to healthcheck etcd |
| API Server-etcd Client | `apiserver-etcd-client.crt` / `.key` | API server → etcd connection |
| Service Account Signing Key | `sa.key` / `sa.pub` | Private key for signing SA tokens + public key |

### Kubeconfig Certificates

The kubeconfig files in `/etc/kubernetes/` each embed a client certificate:

| Kubeconfig | Embedded cert |
|------------|---------------|
| `admin.conf` | admin client cert |
| `super-admin.conf` | super-admin client cert |
| `controller-manager.conf` | controller manager client cert |
| `scheduler.conf` | scheduler client cert |
| `kubelet.conf` | kubelet client cert (after TLS bootstrap) |

## `certs renew`

Renews certificates. The renewed certs are placed in the same `/etc/kubernetes/pki` directory. After renewal, you must **restart the affected control plane static Pods** for them to pick up the new certs (static Pods restart when their manifest file changes).

### Renewable Certificates

| Cert name | What it renews |
|-----------|----------------|
| `admin.conf` | Admin kubeconfig certificate |
| `super-admin.conf` | Super-admin kubeconfig certificate |
| `apiserver` | API server serving certificate |
| `apiserver-kubelet-client` | API server → kubelet client cert |
| `apiserver-etcd-client` | API server → etcd client cert |
| `controller-manager.conf` | Controller manager kubeconfig cert |
| `etcd-healthcheck-client` | etcd healthcheck client cert |
| `etcd-peer` | etcd peer certificate |
| `etcd-server` | etcd server certificate |
| `front-proxy-client` | Front proxy client certificate |
| `scheduler.conf` | Scheduler kubeconfig cert |
| `all` | All of the above |

### Flags

| Flag | Description |
|------|-------------|
| `--cert-dir` | Path to certificates (default: `/etc/kubernetes/pki`) |
| `--config` | Path to a kubeadm config file |
| `--kubeconfig` | Kubeconfig file (default: `/etc/kubernetes/admin.conf`) |

### Example

```bash
# Renew all certificates
kubeadm certs renew all

# Renew just the API server cert
kubeadm certs renew apiserver

# Check what was renewed
kubeadm certs check-expiration
```

### After Renewal

1. Restart the control plane static Pods (they'll auto-restart when manifests are touched)
2. Restart the kubelet: `systemctl restart kubelet`
3. Update your `~/.kube/config` with the new `admin.conf`: `cp /etc/kubernetes/admin.conf ~/.kube/config`
4. For HA clusters, repeat on each control-plane node

## `certs certificate-key`

Generates a new control-plane certificate key — a hex-encoded AES key of 32 bytes (256 bits). Used to encrypt/decrypt certificates in the `kubeadm-certs` Secret for the control-plane join workflow.

```bash
$ kubeadm certs certificate-key
1234abcd5678...
```

This key is used with `--certificate-key` during `kubeadm init --upload-certs` and `kubeadm join --control-plane`. See [init.md](init.md).

## `certs check-expiration`

Checks the expiration dates for all certificates in the local PKI.

```bash
kubeadm certs check-expiration
```

| Flag | Description |
|------|-------------|
| `--cert-dir` | Path to certificates (default: `/etc/kubernetes/pki`) |
| `--config` | Path to a kubeadm config file |
| `--kubeconfig` | Kubeconfig file |
| `-o, --output` | Output format (text, json, yaml, etc.) |

### Output

Shows each certificate's expiration date and remaining time. Certificates renewed via `kubeadm certs renew` get a 1-year validity (configurable via `certificateValidityPeriod` in v1beta4). CA certificates get a 10-year validity (configurable via `caCertificateValidityPeriod`).

## `certs generate-csr`

Generates private keys and Certificate Signing Requests (CSRs) for all control-plane certificates. This is used in **external CA mode** — where you want an external CA to sign the certificates rather than using kubeadm's self-signed CAs.

```bash
kubeadm certs generate-csr --cert-dir /path/to/output
```

| Flag | Description |
|------|-------------|
| `--cert-dir` | Path to save certificates and CSRs |
| `--config` | Path to a kubeadm config file |
| `--kubeconfig-dir` | Path to save kubeconfig files |

### External CA Mode Workflow

1. Run `kubeadm certs generate-csr` to produce keys + CSRs
2. Send the CSRs to your external CA for signing
3. Place the signed certificates back in the cert directory
4. Run `kubeadm init` — kubeadm detects the existing certs and uses them

## Encryption Algorithm

In v1beta4, the `encryptionAlgorithm` field in `ClusterConfiguration` controls the key algorithm:

| Value | Algorithm |
|-------|-----------|
| `RSA-2048` | RSA 2048-bit (default) |
| `RSA-3072` | RSA 3072-bit |
| `RSA-4096` | RSA 4096-bit |
| `ECDSA-P256` | ECDSA P-256 |
| `ECDSA-P384` | ECDSA P-384 |

This replaces the deprecated `PublicKeysECDSA` feature gate.

## Edge Cases

- **Renewing on HA clusters**: You must renew on **each** control-plane node individually. The certs are local to each node.
- **CA renewal**: `kubeadm certs renew` does **not** renew CA certificates. CA rotation requires manual intervention.
- **Static Pod restart**: After renewal, static Pods need to be restarted to pick up new certs. Touching the manifest files or restarting kubelet triggers this.
- **`sa.key` is not renewable**: The service account signing key is not rotated by `certs renew`. Rotating it requires manual steps and is disruptive (all existing SA tokens become invalid).

## References

- [kubeadm certs docs](https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm-certs/)
- [Certificate management with kubeadm](https://kubernetes.io/docs/tasks/administer-cluster/kubeadm/kubeadm-certs/)

## Related

- [init.md](init.md) — certificate generation during init
- [concepts.md](concepts.md) — the PKI and filesystem layout
- [config.md](config.md) — `certificateValidityPeriod` and `encryptionAlgorithm` fields
