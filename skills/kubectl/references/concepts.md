# Core Concepts

The foundational mental model for understanding kubectl — the control plane, API server, resource model, declarative vs imperative management, and the key abstractions kubectl manipulates.

> **Source of truth**: <https://kubernetes.io/docs/reference/kubectl/>, <https://kubernetes.io/docs/concepts/>. For **resource type listings**, see [resources.md](resources.md). For **kubeconfig and contexts**, see [config.md](config.md). For **command reference**, see [cli.md](cli.md).

## What kubectl Is

`kubectl` is a **client-side CLI** that communicates with the Kubernetes **control plane** via the Kubernetes API. It does not run containers, schedule pods, or manage networking itself — it sends REST requests to the API server, which is the single entry point to the cluster.

```
kubectl (client)
  │  HTTPS REST request
  ▼
API Server ─── etcd (cluster state store)
  │
  ├── Controller Manager (reconciliation loops)
  ├── Scheduler (assigns pods to nodes)
  └── Cloud Controller (cloud-specific integrations)

         Nodes (data plane)
           ├── kubelet (pod lifecycle)
           └── kube-proxy (networking)
```

## How kubectl Authenticates

kubectl locates cluster credentials in this order:

1. **`--kubeconfig` flag** — explicit path to a kubeconfig file
2. **`KUBECONFIG` env var** — one or more kubeconfig paths (colon-separated on Linux/macOS, semicolon on Windows), merged in order
3. **`~/.kube/config`** — default location

### In-Cluster Detection

When running inside a pod, kubectl auto-detects in-cluster auth by checking for:

- `KUBERNETES_SERVICE_HOST` environment variable
- `KUBERNETES_SERVICE_PORT` environment variable
- Service account token at `/var/run/secrets/kubernetes.io/serviceaccount/token`

If all three exist, kubectl authenticates using the pod's service account. The `POD_NAMESPACE` env var overrides the default namespace in this mode. Explicit `--namespace <value>` always overrides.

### Authentication Methods

| Method | Mechanism | Key flags / kubeconfig fields |
|--------|-----------|-------------------------------|
| **Client certificate** | TLS client cert + key | `--client-certificate`, `--client-key`, `--certificate-authority` |
| **Bearer token** | `Authorization: Bearer <token>` header | `--token` |
| **Basic auth** | Username + password (deprecated) | `--username`, `--password` |
| **Service account** | In-cluster token (auto-detected) | Automatic in pods |
| **Cloud provider** | Cloud-specific (GCP, AWS, Azure) | Via cloud auth plugins |
| **Impersonation** | Act as another user/group | `--as`, `--as-group`, `--as-uid` |

## The Kubernetes Resource Model

Everything in Kubernetes is a **resource** — a typed object stored in etcd and exposed via the API server. kubectl commands operate on resources.

### Resource Anatomy

Every Kubernetes object has a standard structure:

```yaml
apiVersion: apps/v1        # API group + version
kind: Deployment           # Resource type
metadata:
  name: my-deployment      # Unique within namespace
  namespace: default       # Isolation boundary
  labels:                  # Key-value for selection
    app: myapp
  annotations:             # Key-value for arbitrary metadata
    managed-by: kubectl
  uid: <uuid>              # Immutable unique ID
  resourceVersion: "123"   # Optimistic concurrency token
spec:                      # Desired state (user-provided)
  replicas: 3
  selector:
    matchLabels:
      app: myapp
  template:
    ...
status:                    # Actual state (controller-managed, read-only)
  readyReplicas: 2
  conditions: [...]
```

### Key Metadata Fields

| Field | Scope | Description |
|-------|-------|-------------|
| `apiVersion` | Object | API group and version (e.g. `apps/v1`, `v1`) |
| `kind` | Object | Resource type name (e.g. `Pod`, `Deployment`) |
| `metadata.name` | Namespace | Unique name within the namespace |
| `metadata.namespace` | Object | Isolation boundary (not for cluster-scoped resources) |
| `metadata.labels` | Object | Key-value pairs for label selectors |
| `metadata.annotations` | Object | Key-value pairs for tooling/metadata (non-identifying) |
| `metadata.uid` | Object | Immutable unique identifier |
| `metadata.resourceVersion` | Object | Optimistic concurrency token — changes on every update |
| `metadata.creationTimestamp` | Object | When the object was created |

### Spec vs Status

- **`spec`** — the **desired state**, provided by the user (or kubectl). This is what you *want*.
- **`status`** — the **actual state**, managed by controllers. This is what *is*. Read-only from the client's perspective.

Controllers run a **reconciliation loop**: observe current state → compare to desired state → take action to converge.

## Declarative vs Imperative Management

kubectl supports two paradigms for managing resources:

| Paradigm | Commands | How it works |
|----------|----------|--------------|
| **Imperative** | `create`, `delete`, `replace`, `patch`, `set`, `scale`, `expose`, `run` | kubectl sends a direct action (create this, delete that) |
| **Declarative** | `apply`, `diff` | kubectl computes the diff between config and live state, then patches |

### Imperative Example

```shell
kubectl create deployment nginx --image=nginx --replicas=3
kubectl scale deployment nginx --replicas=5
kubectl delete deployment nginx
```

Each command is a discrete action. The cluster state is not derived from a file.

### Declarative Example

```shell
kubectl apply -f deployment.yaml      # create or update
kubectl diff -f deployment.yaml        # preview changes
kubectl apply -f deployment.yaml --prune --all  # remove orphaned resources
```

The cluster state is derived from configuration files. Running `apply` repeatedly converges the cluster to match the files.

### When to Use Which

| Scenario | Recommended |
|----------|-------------|
| Quick one-off operations | Imperative |
| Production/GitOps workflows | Declarative |
| Creating from generated manifest | Declarative (`apply`) |
| Emergency scaling | Imperative (`scale`) |
| Long-term infrastructure | Declarative (version-controlled YAML) |

See [apply.md](apply.md) for the full declarative management reference.

## Core Resource Abstractions

### Pods

The **smallest deployable unit**. A pod contains one or more containers that share network and storage. Pods are ephemeral — they are created, destroyed, and replaced, never healed in place.

```shell
kubectl get pods                    # list pods in current namespace
kubectl get pods -A                 # list across all namespaces
kubectl get pod my-pod -o yaml      # full spec in YAML
kubectl describe pod my-pod         # detailed state + events
```

Short name: `po`. API: `v1` (core group).

### Deployments

Manages a set of identical pods via **ReplicaSets**. Provides rolling updates, rollbacks, and scaling. The most common way to run stateless workloads.

```shell
kubectl create deployment web --image=nginx --replicas=3
kubectl rollout status deployment/web
kubectl rollout undo deployment/web
kubectl scale deployment/web --replicas=5
```

Short name: `deploy`. API: `apps/v1`.

### Services

An abstraction that defines a logical set of pods and a policy to access them. Provides stable networking (DNS name, IP) despite pods being ephemeral.

| Type | Description |
|------|-------------|
| `ClusterIP` | Exposes on cluster-internal IP (default) |
| `NodePort` | Exposes on each node's IP at a static port |
| `LoadBalancer` | Cloud provider load balancer |
| `ExternalName` | DNS CNAME to an external service |

Short name: `svc`. API: `v1` (core group).

```shell
kubectl expose deployment web --port=80 --target-port=80 --type=LoadBalancer
kubectl get services
kubectl get svc web -o wide
```

### Namespaces

Virtual clusters within a physical cluster. Used to isolate resources, apply quotas, and scope RBAC. Not all resources are namespaced — nodes, persistent volumes, and namespaces themselves are cluster-scoped.

```shell
kubectl get namespaces
kubectl get pods -n kube-system
kubectl config set-context --current --namespace=dev
```

Short name: `ns`. API: `v1` (core group).

### Nodes

A node is a worker machine (VM or physical) in the cluster. Each node runs the kubelet and kube-proxy. Nodes are cluster-scoped (not namespaced).

```shell
kubectl get nodes -o wide
kubectl describe node worker-1
kubectl cordon node worker-1      # mark unschedulable
kubectl drain node worker-1       # evict pods for maintenance
kubectl uncordon node worker-1    # mark schedulable again
```

Short name: `no`. API: `v1` (core group).

### ConfigMaps and Secrets

**ConfigMaps** store non-confidential configuration data as key-value pairs. **Secrets** store sensitive data (passwords, tokens, keys) — base64-encoded, and can be mounted as volumes or exposed as environment variables.

```shell
kubectl create configmap app-config --from-literal=key1=value1
kubectl create secret generic db-secret --from-literal=password=s3cret
kubectl get cm,secret
```

ConfigMap short name: `cm`. API: `v1` (core group).

### Service Accounts

An identity for processes running in a pod. Used by pods to authenticate to the API server. Bound to roles via RoleBinding or ClusterRoleBinding for authorization.

```shell
kubectl get serviceaccounts
kubectl create serviceaccount my-sa
kubectl auth can-i list pods --as=system:serviceaccount:default:my-sa
```

Short name: `sa`. API: `v1` (core group).

## Labels and Annotations

### Labels

Key-value pairs for **identifying and selecting** resources. Used by selectors to group resources (e.g., a Service selects pods by label).

```shell
kubectl get pods -l app=nginx
kubectl get pods -l 'app in (nginx, redis)'
kubectl label pod my-pod environment=production
kubectl label pod my-pod environment-        # remove label
```

Label selectors support: `=`, `==`, `!=`, `in`, `notin`, `exists`. See [resources.md](resources.md) for selector syntax details.

### Annotations

Key-value pairs for **arbitrary non-identifying metadata**. Not used for selection. Used by tooling, kubectl itself (e.g., `last-applied-configuration`), and controllers.

```shell
kubectl annotate pod my-pod description="my test pod"
kubectl annotate pod my-pod description-     # remove annotation
```

## Controllers and Reconciliation

Kubernetes runs **controllers** that watch the cluster state and drive it toward the desired state:

| Controller | Manages | Reconciles |
|-----------|---------|------------|
| Deployment controller | Deployments → ReplicaSets | Creates/deletes ReplicaSets to match desired replicas |
| ReplicaSet controller | ReplicaSets → Pods | Creates/deletes Pods to match replica count |
| StatefulSet controller | StatefulSets → Pods | Ordered creation/deletion with stable identity |
| DaemonSet controller | DaemonSets → Pods | One Pod per node |
| Job controller | Jobs → Pods | Runs pods to completion, handles retries |
| Service controller | Services → Endpoints | Maintains endpoint list from pod labels |

kubectl commands like `apply`, `scale`, and `rollout` modify the desired state (`spec`); controllers handle the actual convergence. This is why `kubectl scale deployment --replicas=5` returns immediately — it sets the desired count, and the ReplicaSet controller creates the pods asynchronously.

## Edge Cases and Known Issues

- **Pods are ephemeral**: Never rely on a pod's IP or name persisting. Use Services for stable addressing.
- **`kubectl apply` on resources created imperatively**: Without `--save-config` or `kubectl.kubernetes.io/last-applied-configuration` annotation, the first `apply` may conflict with existing fields. Use `kubectl apply --force` or create with `--save-config`.
- **`kubectl get` for CRDs**: Custom Resource Definitions appear in `kubectl api-resources` but may have non-standard behavior. Use `kubectl get <crd-name>` to list them.
- **Namespace deletion is not instant**: Deleting a namespace triggers cleanup of all resources in it. This can hang if finalizers are pending.
- **`--namespace` does not affect cluster-scoped resources**: `kubectl get nodes -n my-ns` ignores the namespace flag.
- **In-cluster auth uses the pod's SA**: Running kubectl inside a pod uses the pod's service account token. Ensure the SA has the necessary RBAC permissions.
- **`kubectl logs` requires a running pod**: If the pod hasn't started or has crashed, `logs` may return empty or error. Use `--previous` for crashed container logs.

## References

- [kubectl Overview](https://kubernetes.io/docs/reference/kubectl/)
- [Kubernetes Concepts](https://kubernetes.io/docs/concepts/)
- [Labels and Selectors](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/)

## Related Skills

- For resource types and API groups, see [resources.md](resources.md).
- For kubeconfig and contexts, see [config.md](config.md).
- For declarative apply, see [apply.md](apply.md).
