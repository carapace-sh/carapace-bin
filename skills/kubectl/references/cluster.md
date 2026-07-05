# Cluster Management Commands

Commands for managing the cluster infrastructure — nodes (cordon, uncordon, drain, taint), cluster info, certificates, RBAC authorization, and resource metrics.

> **Source of truth**: `kubectl <command> --help`, <https://kubernetes.io/docs/reference/kubectl/generated/kubectl_cordon/>. For **debugging workloads** (logs, exec, describe), see [debugging.md](debugging.md). For **node concepts**, see [concepts.md](concepts.md).

## Node Management

### `kubectl cordon`

Marks a node as **unschedulable** — no new pods will be placed on it. Existing pods are not affected:

```shell
kubectl cordon worker-1
kubectl cordon node/worker-1
```

This sets `spec.unschedulable: true` on the node. The scheduler skips cordoned nodes when placing new pods.

### `kubectl uncordon`

Marks a node as **schedulable** again — reverses `cordon`:

```shell
kubectl uncordon worker-1
```

### `kubectl drain`

Safely **evicts all pods** from a node in preparation for maintenance. Combines cordon + eviction:

```shell
kubectl drain worker-1
kubectl drain worker-1 --ignore-daemonsets
kubectl drain worker-1 --ignore-daemonsets --delete-emptydir-data
kubectl drain worker-1 --force                         # evict unmanaged pods too
kubectl drain worker-1 --grace-period=30
kubectl drain worker-1 --timeout=5m
kubectl drain worker-1 --dry-run=client               # preview
kubectl drain worker-1 --disable-eviction              # use delete instead of evict
```

### How Drain Works

1. **Cordons** the node (marks unschedulable)
2. **Evicts** or **deletes** each pod on the node:
   - Respects `PodDisruptionBudgets` (unless `--disable-eviction`)
   - Skips DaemonSet pods (unless `--ignore-daemonsets` is not set, in which case it errors)
   - Skips pods with `emptyDir` volumes (unless `--delete-emptydir-data` or `--force`)

### Flags

| Flag | Description |
|------|-------------|
| `--ignore-daemonsets` | Skip DaemonSet-managed pods |
| `--delete-emptydir-data` | Continue even if pods use emptyDir (data will be lost) |
| `--force` | Continue even if pods don't have a controller (standalone pods) |
| `--disable-eviction` | Use delete instead of eviction (bypasses PDB) |
| `--grace-period` | Grace period per pod (seconds) |
| `--timeout` | Max time to wait for eviction |
| `--dry-run` | `client` or `server` |
| `--pod-selector` | Only evict pods matching this label selector |
| `-l`, `--selector` | Label selector (for the pods to evict) |
| `--skip-wait-for-delete-timeout` | Skip waiting for pods with old deletion timestamp |

### Drain vs Cordon

| Action | What it does |
|--------|-------------|
| `cordon` | Mark unschedulable (no eviction) |
| `drain` | Cordon + evict all pods |
| `uncordon` | Mark schedulable again |

Typical maintenance workflow:

```shell
kubectl drain worker-1 --ignore-daemonsets --delete-emptydir-data
# ... perform maintenance on worker-1 ...
kubectl uncordon worker-1
```

### `kubectl taint`

Updates taints on nodes. Taints repel pods — a pod is only scheduled on a tainted node if it has a matching **toleration**:

```shell
kubectl taint nodes worker-1 key=value:NoSchedule
kubectl taint nodes worker-1 key:NoSchedule-                  # remove taint
kubectl taint nodes worker-1 key=value:NoSchedule --overwrite  # update taint
kubectl taint nodes worker-1 dedicated=gpu:PreferNoSchedule    # soft taint
```

### Taint Format

```
KEY=VALUE:EFFECT
```

### Taint Effects

| Effect | Description |
|--------|-------------|
| `NoSchedule` | No new pods will be scheduled (existing pods stay) |
| `NoExecute` | Evict existing pods that don't tolerate the taint |
| `PreferNoSchedule` | Try not to schedule, but don't strictly prevent |

### Removing Taints

Append `-` to the key/effect:

```shell
kubectl taint nodes worker-1 key:NoSchedule-       # remove specific taint
kubectl taint nodes worker-1 key-                   # remove all taints with key
```

### Common Use Cases

```shell
# Dedicated node for GPU workloads
kubectl taint nodes gpu-node-1 nvidia.com/gpu=:NoSchedule

# Maintenance taint (evicts non-tolerating pods)
kubectl taint nodes worker-1 node.kubernetes.io/under-maintenance:NoExecute

# Node not ready (set automatically by node controller)
kubectl taint nodes worker-1 node.kubernetes.io/not-ready:NoSchedule
```

## `kubectl cluster-info`

Displays endpoint information about the control plane and services:

```shell
kubectl cluster-info
kubectl cluster-info dump                          # dump all cluster info to stdout
kubectl cluster-info dump --output-directory=./cluster-state
kubectl cluster-info dump --namespaces=kube-system,default
```

### `cluster-info dump`

Collects diagnostic information from the cluster — useful for bug reports and troubleshooting:

```shell
kubectl cluster-info dump --output-directory=./debug
kubectl cluster-info dump --namespaces=default,kube-system
```

Dumps include:
- Node descriptions and conditions
- Pod descriptions across namespaces
- Service and endpoint info
- Events

## `kubectl certificate`

Manages CertificateSigningRequest (CSR) resources:

```shell
kubectl certificate approve my-csr
kubectl certificate deny my-csr
kubectl get csr
kubectl describe csr my-csr
```

### Approve / Deny

```shell
kubectl certificate approve my-csr
kubectl certificate deny my-csr --reason="key usage not allowed"
```

The signing process:
1. A client creates a CSR resource with a certificate request
2. A controller (or human) approves or denies it
3. If approved, a signer controller signs the certificate and updates the status

## `kubectl auth`

Inspects authorization — what a user or service account can do:

### Subcommands

| Subcommand | Description |
|------------|-------------|
| `auth can-i` | Check if you can perform an action |
| `auth reconcile` | Reconcile RBAC rules from a file |
| `auth whoami` | Show your current identity |

### `auth can-i`

```shell
kubectl auth can-i list pods
kubectl auth can-i create deployments
kubectl auth can-i '*' "*"
kubectl auth can-i list pods --all-namespaces
kubectl auth can-i get pods --as=system:serviceaccount:default:my-sa
kubectl auth can-i get pods --as=alice
kubectl auth can-i get pods -n dev
kubectl auth can-i list pods --subresource=status
```

Returns `yes` or `no`. Supports `--as` for impersonation (checking another user's permissions).

### `auth reconcile`

Applies RBAC resources (Role, ClusterRole, RoleBinding, ClusterRoleBinding) from a file, removing rules that are no longer present:

```shell
kubectl auth reconcile -f rbac.yaml
kubectl auth reconcile -f rbac.yaml --remove-extra-permissions
kubectl auth reconcile -f rbac.yaml --remove-extra-subjects
```

Unlike `apply`, `reconcile` for RBAC removes old rules/subjects that are no longer in the file. `apply` would leave them (because of the strategic merge behavior on lists).

### `auth whoami`

```shell
kubectl auth whoami
kubectl auth whoami -o yaml
```

Shows the authenticated identity (requires the `SelfSubjectReview` API).

## RBAC Management

### Creating RBAC Resources

```shell
# Role — namespace-scoped permissions
kubectl create role pod-reader --verb=get,list,watch --resource=pods
kubectl create role pod-exec --verb=get --resource=pods --resource-name=my-pod

# RoleBinding — binds a role to subjects in a namespace
kubectl create rolebinding read-pods \
  --role=pod-reader \
  --serviceaccount=default:my-sa

kubectl create rolebinding alice-reads \
  --role=pod-reader \
  --user=alice \
  --group=devs

# ClusterRole — cluster-wide or reusable permissions
kubectl create clusterrole node-reader --verb=get,list --resource=nodes
kubectl create clusterrole pod-reader-global --verb=get,list,watch --resource=pods

# ClusterRoleBinding — binds a ClusterRole cluster-wide
kubectl create clusterrolebinding read-nodes \
  --clusterrole=node-reader \
  --serviceaccount=default:my-sa

kubectl create clusterrolebinding all-authenticated \
  --clusterrole=system:node \
  --group=system:nodes
```

### RBAC Subjects

Subjects in RoleBindings and ClusterRoleBindings can be:

| Subject Type | Example | Scope |
|-------------|---------|-------|
| `User` | `--user=alice` | Cluster-wide identity (from auth) |
| `Group` | `--group=devs` | Cluster-wide group (from auth) |
| `ServiceAccount` | `--serviceaccount=namespace:name` | Namespace-scoped |

### Viewing RBAC

```shell
kubectl get roles,rolebindings -A
kubectl get clusterroles,clusterrolebindings
kubectl describe role pod-reader
kubectl describe clusterrolebinding cluster-admin
kubectl auth can-i list pods --as=system:serviceaccount:default:my-sa
```

### Common ClusterRoles

| ClusterRole | Description |
|-------------|-------------|
| `cluster-admin` | Super-user, all operations on all resources |
| `admin` | All operations on namespaced resources |
| `edit` | Modify most resources, can't modify RBAC |
| `view` | Read-only access to most resources |

## `kubectl top` (Cluster-Level)

For cluster-wide resource usage, see the `top` section in [debugging.md](debugging.md). Key commands for cluster management:

```shell
kubectl top nodes                                 # node CPU/memory usage
kubectl top nodes --sort-by=cpu                   # busiest nodes
kubectl top pods -A                               # pod usage across all namespaces
kubectl top pods -A --sort-by=memory              # memory-hungriest pods
```

## Node Inspection

```shell
kubectl get nodes -o wide
kubectl get nodes -o custom-columns=NAME:.metadata.name,VERSION:.status.nodeInfo.kubeletVersion
kubectl get nodes --show-labels
kubectl get nodes -l node-role.kubernetes.io/worker=
kubectl describe node worker-1                    # full details: capacity, conditions, pods
kubectl get pods --field-selector spec.nodeName=worker-1 --all-namespaces  # pods on a node
```

### Node Conditions

Visible in `describe node` output:

| Condition | Description |
|-----------|-------------|
| `Ready` | Node is healthy and accepting pods |
| `MemoryPressure` | Node memory is low |
| `DiskPressure` | Node disk is low |
| `PIDPressure` | Too many processes on node |
| `NetworkUnavailable` | Node network is misconfigured |

## Edge Cases and Known Issues

- **`drain` may hang**: If a pod has a `PodDisruptionBudget` that prevents eviction, drain will wait. Use `--timeout` to limit the wait, or `--disable-eviction` to bypass PDBs.
- **`drain --ignore-daemonsets` is almost always needed**: DaemonSet pods can't be evicted (they're managed by the DaemonSet controller). Without this flag, drain fails on any node with DaemonSet pods.
- **`taint NoExecute` evicts immediately**: Unlike `NoSchedule` (which only prevents new pods), `NoExecute` evicts existing pods that don't tolerate the taint, with a grace period.
- **`uncordon` doesn't move pods back**: After draining and uncordoning, the evicted pods are gone. New pods may be scheduled, but old ones aren't restored.
- **`certificate approve` requires permission**: The user must have `approve` permission on `certificatesigningrequests/approval`. The built-in approver is `kubernetes.io/kube-apiserver-client` or `kubernetes.io/legacy-unknown`.
- **`auth can-i` checks the API server**: It sends a `SelfSubjectAccessReview` to the API server, so it reflects the server's actual RBAC evaluation, not a client-side guess.
- **`auth reconcile` removes rules not in the file**: This differs from `kubectl apply`, which would preserve old rules. Use `reconcile` for RBAC cleanup.
- **Node conditions are set by controllers**: `Ready`, `MemoryPressure`, etc. are set by the kubelet and node controller. Manually changing them is not supported.
- **`top` requires metrics-server**: Without the metrics-server addon, `top` returns an error. Install it separately if needed.

## References

- [kubectl cordon](https://kubernetes.io/docs/reference/kubectl/generated/kubectl_cordon/)
- [kubectl drain](https://kubernetes.io/docs/reference/kubectl/generated/kubectl_drain/)
- [kubectl taint](https://kubernetes.io/docs/reference/kubectl/generated/kubectl_taint/)
- [Taints and Tolerations](https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/)
- [Using RBAC Authorization](https://kubernetes.io/docs/reference/access-authn-authz/rbac/)

## Related Skills

- For debugging commands (logs, exec, top), see [debugging.md](debugging.md).
- For node and pod concepts, see [concepts.md](concepts.md).
- For workload management (scale, rollout), see [workloads.md](workloads.md).
