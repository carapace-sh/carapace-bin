---
name: kubeadm
description: >
  Use when working with the kubeadm CLI — bootstrapping Kubernetes clusters, init/join/reset/upgrade
  workflows, the phase system, certificate management, bootstrap tokens, kubeconfig generation,
  and the v1beta4 configuration API. Triggers on: "kubeadm", "kubeadm init", "kubeadm join",
  "kubeadm reset", "kubeadm upgrade", "kubeadm config", "kubeadm token", "kubeadm certs",
  "kubeadm kubeconfig", "kubeadm version", "kubeadm alpha", "bootstrap token",
  "InitConfiguration", "ClusterConfiguration", "JoinConfiguration", "ResetConfiguration",
  "NodeRegistrationOptions", "control-plane-endpoint", "discovery-token", "certificate-key",
  "feature-gates", "skip-phases", "static pod manifest", "TLS bootstrap", "kubeadm-config".
user-invocable: true
---

# kubeadm In-Depth Reference

A structured reference for the `kubeadm` CLI — the Kubernetes cluster bootstrapping tool. Covers the command surface, the phase workflow system, the v1beta4 configuration API, certificate and token management, and the conceptual model behind cluster init/join/upgrade/reset. See <https://kubernetes.io/docs/reference/setup-tools/kubeadm/>.

## Data Flow

```
kubeadm init (control-plane node)
  → preflight checks
    → PKI/cert generation        (/etc/kubernetes/pki)
      → kubeconfig generation    (/etc/kubernetes/*.conf)
        → static Pod manifests   (/etc/kubernetes/manifests)
          → kubelet starts control plane Pods
            → upload-config + bootstrap-token
              → addons (CoreDNS, kube-proxy)
                → join command printed

kubeadm join (worker or control-plane node)
  → discovery (token or file)
    → TLS bootstrap (CSR auto-approved)
      → kubelet definitive identity
        → [control-plane] download certs + etcd join
```

## Sub-Resources

Load the reference that matches your task. When in doubt, load multiple references.

| Keywords | Reference |
|----------|----------|
| concepts, bootstrapping, control plane, etcd, static pods, bootstrap tokens, TLS bootstrap, discovery, kubeconfig files, feature gates | [references/concepts.md](references/concepts.md) |
| init, kubeadm init, control-plane init, init phases, init flags, pod-network-cidr, service-cidr, upload-certs, certificate-key | [references/init.md](references/init.md) |
| join, kubeadm join, worker node, control-plane join, discovery-token, discovery-file, ca-cert-hash, tls-bootstrap-token | [references/join.md](references/join.md) |
| phases, skip-phases, phase tree, init phase, join phase, reset phase, upgrade phase, phase invocation | [references/phases.md](references/phases.md) |
| config, kubeadm config, InitConfiguration, ClusterConfiguration, JoinConfiguration, ResetConfiguration, UpgradeConfiguration, NodeRegistrationOptions, v1beta4, print init-defaults, migrate, validate, images | [references/config.md](references/config.md) |
| certs, kubeadm certs, renew, certificate-key, check-expiration, generate-csr, certificate inventory, CA, external CA | [references/certs.md](references/certs.md) |
| token, kubeadm token, bootstrap token, create, delete, generate, list, token-ttl, token usages | [references/token.md](references/token.md) |
| upgrade, kubeadm upgrade, apply, plan, node, diff, upgrade phases, certificate-renewal, etcd-upgrade | [references/upgrade.md](references/upgrade.md) |
| reset, kubeadm reset, reset phases, cleanup-node, remove-etcd-member, force reset | [references/reset.md](references/reset.md) |
| kubeconfig, kubeadm kubeconfig, kubeconfig user, version, alpha | [references/misc.md](references/misc.md) |
| global flags, rootfs, v, vmodule, kubeconfig flag, dry-run, ignore-preflight-errors, patches, cri-socket, preflight checks | [references/flags.md](references/flags.md) |

## Quick Guide

- **How do I bootstrap a control plane?** → [references/init.md](references/init.md)
- **How do I join a node to a cluster?** → [references/join.md](references/join.md)
- **How does the phase system work?** → [references/phases.md](references/phases.md)
- **How do I write a kubeadm config file?** → [references/config.md](references/config.md)
- **How do I renew certificates?** → [references/certs.md](references/certs.md)
- **How do I manage bootstrap tokens?** → [references/token.md](references/token.md)
- **How do I upgrade a cluster?** → [references/upgrade.md](references/upgrade.md)
- **How do I tear down a node?** → [references/reset.md](references/reset.md)
- **What does kubeadm actually do under the hood?** → [references/concepts.md](references/concepts.md)
- **What are the common flags across commands?** → [references/flags.md](references/flags.md)
- **How do I generate a kubeconfig for an extra user?** → [references/misc.md](references/misc.md)

## Cross-Project References

- For Kubernetes API concepts (Pods, Services, ConfigMaps, Secrets, RBAC) in general, consult the Kubernetes documentation at <https://kubernetes.io/docs/concepts/>.
- For `kubectl` usage, see the kubectl reference at <https://kubernetes.io/docs/reference/kubectl/>.
- For the kubeadm completer implementation in this repo, see `completers/common/kubeadm_completer/`.
