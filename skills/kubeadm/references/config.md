# kubeadm config — Configuration Management

The `kubeadm config` command manages kubeadm configuration: printing defaults, migrating old config versions, validating config files, and interacting with container images.

> **Source of truth**: <https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm-config/> and the [v1beta4 API reference](https://kubernetes.io/docs/reference/config-api/kubeadm-config.v1beta4/).

## Subcommands

| Subcommand | Purpose |
|------------|---------|
| `kubeadm config print` | Print configuration (parent) |
| `kubeadm config print init-defaults` | Print default init configuration |
| `kubeadm config print join-defaults` | Print default join configuration |
| `kubeadm config print reset-defaults` | Print default reset configuration |
| `kubeadm config print upgrade-defaults` | Print default upgrade configuration |
| `kubeadm config migrate` | Convert old config to latest API version |
| `kubeadm config validate` | Validate a config file |
| `kubeadm config images list` | List images kubeadm will use |
| `kubeadm config images pull` | Pull images used by kubeadm |

## The `--config` Flag

The preferred way to configure kubeadm is a YAML file passed via `--config`. A single file can contain **multiple configuration types separated by `---`**:

```yaml
apiVersion: kubeadm.k8s.io/v1beta4
kind: InitConfiguration
bootstrapTokens:
  - token: abcdef.0123456789abcdef
    ttl: 24h0m0s
nodeRegistration:
  name: control-plane-1
  criSocket: unix:///var/run/containerd/containerd.sock
---
apiVersion: kubeadm.k8s.io/v1beta4
kind: ClusterConfiguration
networking:
  podSubnet: 10.244.0.0/16
  serviceSubnet: 10.96.0.0/12
  dnsDomain: cluster.local
kubernetesVersion: stable-1
controlPlaneEndpoint: "10.0.0.1:6443"
```

If a configuration type is not provided, kubeadm uses defaults. If an unexpected type is provided, kubeadm ignores it and prints a warning.

## The v1beta4 API

API version: `kubeadm.k8s.io/v1beta4` (introduced in kubeadm v1.31). Use `kubeadm config migrate` to convert from `v1beta3`.

### Supported Configuration Types

| Type | Used with | Purpose |
|------|-----------|---------|
| `InitConfiguration` | `kubeadm init` | Runtime settings for init (bootstrap tokens, node registration) |
| `ClusterConfiguration` | `kubeadm init` | Cluster-wide settings (networking, etcd, control plane, certs) |
| `JoinConfiguration` | `kubeadm join` | Runtime settings for join (discovery, node registration, control plane) |
| `ResetConfiguration` | `kubeadm reset` | Runtime settings for reset |
| `UpgradeConfiguration` | `kubeadm upgrade` | Runtime settings for upgrade |
| `KubeletConfiguration` | `kubeadm init` / `join` | Kubelet component config (uploaded to ConfigMap) |
| `KubeProxyConfiguration` | `kubeadm init` | kube-proxy component config |

## InitConfiguration

Runtime settings specific to `kubeadm init`. These fields are **not** uploaded to the `kubeadm-config` ConfigMap.

| Field | Type | Description |
|-------|------|-------------|
| `bootstrapTokens` | `[]BootstrapToken` | Bootstrap tokens to create |
| `dryRun` | `bool` | Dry run mode |
| `nodeRegistration` | `NodeRegistrationOptions` | Node registration settings |
| `localAPIEndpoint` | `APIEndpoint` | API server endpoint on this node (`advertiseAddress` + `bindPort`) |
| `certificateKey` | `string` | AES key (32 bytes, hex) for encrypting uploaded certs |
| `skipPhases` | `[]string` | Phases to skip |
| `patches` | `Patches` | Patch options for components |
| `timeouts` | `Timeouts` | Various timeouts |

## ClusterConfiguration

Cluster-wide settings — networking, etcd, control plane components, certificates.

| Field | Type | Default | Description |
|-------|------|---------|-------------|
| `etcd` | `Etcd` | — | etcd config (local or external) |
| `networking` | `Networking` | — | `serviceSubnet`, `podSubnet`, `dnsDomain` |
| `kubernetesVersion` | `string` | `stable-1` | Target control plane version |
| `controlPlaneEndpoint` | `string` | — | Stable IP/DNS for the control plane |
| `apiServer` | `APIServer` | — | Extra API server settings (`extraArgs`, `extraVolumes`, `certSANs`) |
| `controllerManager` | `ControlPlaneComponent` | — | Extra controller manager settings |
| `scheduler` | `ControlPlaneComponent` | — | Extra scheduler settings |
| `dns` | `DNS` | — | CoreDNS options (`disabled: true` to skip) |
| `proxy` | `Proxy` | — | kube-proxy options (`disabled: true` to skip) |
| `certificatesDir` | `string` | `/etc/kubernetes/pki` | Certificate storage path |
| `imageRepository` | `string` | `registry.k8s.io` | Container registry |
| `featureGates` | `map[string]bool` | — | kubeadm feature gates |
| `clusterName` | `string` | `kubernetes` | Cluster name |
| `encryptionAlgorithm` | `string` | `RSA-2048` | `RSA-2048`, `RSA-3072`, `RSA-4096`, `ECDSA-P256`, `ECDSA-P384` |
| `certificateValidityPeriod` | `Duration` | `8760h` (1 year) | Validity for non-CA certs |
| `caCertificateValidityPeriod` | `Duration` | `87600h` (10 years) | Validity for CA certs |

## JoinConfiguration

| Field | Type | Description |
|-------|------|-------------|
| `dryRun` | `bool` | Dry run mode |
| `nodeRegistration` | `NodeRegistrationOptions` | Node registration settings |
| `caCertPath` | `string` | Path to CA cert (default: `/etc/kubernetes/pki/ca.crt`) |
| `discovery` | `Discovery` | **Required.** Discovery method (bootstrapToken or file) |
| `controlPlane` | `JoinControlPlane` | Control plane instance to deploy (nil = worker only) |
| `skipPhases` | `[]string` | Phases to skip |
| `patches` | `Patches` | Patch options |
| `timeouts` | `Timeouts` | Various timeouts |

### Discovery

```yaml
discovery:
  bootstrapToken:
    apiServerEndpoint: 10.0.0.1:6443
    token: abcdef.0123456789abcdef
    caCertHashes:
      - sha256:1234...
  # OR
  file:
    kubeConfigPath: /path/to/cluster-info.yaml
```

## ResetConfiguration

| Field | Type | Description |
|-------|------|-------------|
| `cleanupTmpDir` | `bool` | Clean `/etc/kubernetes/tmp` |
| `certificatesDir` | `string` | Cert directory to clean |
| `criSocket` | `string` | CRI socket path |
| `dryRun` | `bool` | Dry run mode |
| `force` | `bool` | Reset without prompting |
| `ignorePreflightErrors` | `[]string` | Preflight errors to ignore |
| `skipPhases` | `[]string` | Phases to skip |
| `unmountFlags` | `[]string` | `unmount2()` flags: `MNT_FORCE`, `MNT_DETACH`, `MNT_EXPIRE`, `UMOUNT_NOFOLLOW` |
| `timeouts` | `Timeouts` | Various timeouts |

## NodeRegistrationOptions

Appears in `InitConfiguration` and `JoinConfiguration`.

| Field | Type | Default | Description |
|-------|------|---------|-------------|
| `name` | `string` | hostname | Node name (also kubelet cert CN) |
| `criSocket` | `string` | auto-detect | Container runtime socket path |
| `taints` | `[]Taint` | control-plane taint | Taints for the Node. `[]` for none. |
| `kubeletExtraArgs` | `[]Arg` | — | Extra kubelet CLI args (structured `name`+`value`, supports duplicates) |
| `ignorePreflightErrors` | `[]string` | — | Preflight errors to ignore |
| `imagePullPolicy` | `string` | `IfNotPresent` | `Always`, `IfNotPresent`, or `Never` |
| `imagePullSerial` | `bool` | `true` | Pull images serially vs. in parallel |

## `config print`

### `print init-defaults`

Prints a default `InitConfiguration` + `ClusterConfiguration` that can be used as a starting point.

```bash
kubeadm config print init-defaults > kubeadm-init.yaml
```

| Flag | Description |
|------|-------------|
| `--component-configs` | Comma-separated list: `KubeProxyConfiguration`, `KubeletConfiguration` |

### `print join-defaults`

Prints a default `JoinConfiguration` for `kubeadm join`.

### `print reset-defaults` / `print upgrade-defaults`

Print defaults for `kubeadm reset` and `kubeadm upgrade` respectively (v1beta4).

## `config migrate`

Converts an old API version config to the latest (`v1beta4`).

```bash
kubeadm config migrate --old-config v1beta3.yaml --new-config v1beta4.yaml
```

| Flag | Description |
|------|-------------|
| `--old-config` | Path to old config file (mandatory) |
| `--new-config` | Path for new config output (default: stdout) |
| `--allow-experimental-api` | Allow migration to unreleased APIs |

## `config validate`

Validates a config file and reports any problems.

```bash
kubeadm config validate --config kubeadm-init.yaml
```

| Flag | Description |
|------|-------------|
| `--config` | Path to config file |
| `--allow-deprecated-api` | Allow deprecated APIs |
| `--allow-experimental-api` | Allow experimental APIs |

## `config images`

### `images list`

Prints the container images kubeadm will use for the control plane.

```bash
kubeadm config images list --kubernetes-version v1.30.0
```

| Flag | Description |
|------|-------------|
| `--config` | Path to config file |
| `--feature-gates` | Feature gates |
| `--image-repository` | Container registry (default: `registry.k8s.io`) |
| `--kubernetes-version` | Kubernetes version (default: `stable-1`) |
| `-o, --output` | Output format (text, json, yaml, etc.) |

### `images pull`

Pulls the images kubeadm will use. Useful for air-gapped or pre-populated environments.

```bash
kubeadm config images pull --image-repository my.registry.com
```

| Flag | Description |
|------|-------------|
| `--config` | Path to config file |
| `--cri-socket` | Path to CRI socket |
| `--feature-gates` | Feature gates |
| `--image-repository` | Container registry |
| `--kubernetes-version` | Kubernetes version |

## v1beta4 Changes from v1beta3

- `ResetConfiguration` and `UpgradeConfiguration` now supported via `--config`
- `extraArgs` / `kubeletExtraArgs` changed from `map[string]string` to structured `[]Arg` (supports duplicate keys)
- `Timeouts` structure added to init/join/reset/upgrade configs
- `encryptionAlgorithm` field added (replaces `PublicKeysECDSA` feature gate)
- `dns.disabled` and `proxy.disabled` added
- `nodeRegistration.imagePullSerial` added
- `certificateValidityPeriod` and `caCertificateValidityPeriod` added
- `extraEnvs` added for control plane components
- `httpEndpoints` added to external etcd config

## References

- [kubeadm config CLI](https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm-config/)
- [v1beta4 API reference](https://kubernetes.io/docs/reference/config-api/kubeadm-config.v1beta4/)

## Related

- [init.md](init.md) — init command flags
- [join.md](join.md) — join command flags
- [reset.md](reset.md) — reset command
- [upgrade.md](upgrade.md) — upgrade command
