# Resource Types, API Groups, and Selectors

How Kubernetes organizes resources into API groups, the short names and scopes for each type, and the resource specification syntax kubectl accepts on the command line.

> **Source of truth**: `kubectl api-resources`, `kubectl api-versions`, <https://kubernetes.io/docs/reference/kubectl/>. For **core concepts**, see [concepts.md](concepts.md). For **output formatting and field selectors**, see [output.md](output.md).

## Resource Specification Syntax

kubectl accepts resources in several forms on the command line:

| Form | Example | Description |
|------|---------|-------------|
| `TYPE` | `kubectl get pods` | All resources of that type |
| `TYPE NAME` | `kubectl get pod my-pod` | Single resource (space-separated) |
| `TYPE/NAME` | `kubectl get pod/my-pod` | Single resource (slash-separated) |
| `TYPE1/NAME1 TYPE2/NAME2` | `kubectl get pod/web svc/web-svc` | Multiple resource types in one command |
| `TYPE1 name1 name2` | `kubectl get pod web1 web2` | Multiple resources of the same type |
| `-f FILENAME` | `kubectl get -f pod.yaml` | From a YAML/JSON file |
| `-f DIR -R` | `kubectl apply -f manifests/ -R` | All files in a directory, recursively |
| `-k DIR` | `kubectl apply -k overlays/prod` | From a kustomization directory |

### Rules

- Resource **types are case-insensitive**: `pod`, `Pod`, `POD` are all valid
- **Names are case-sensitive**: `my-pod` ≠ `My-Pod`
- **Singular, plural, and abbreviated** forms are interchangeable: `pod`, `pods`, `po`
- Resource types can include the **API group**: `deployments.apps`, `cronjobs.batch`
- The **fully-qualified** form is `TYPE.VERSION.GROUP` (e.g., `deployments.v1.apps`)

### Multi-Resource Operations

```shell
# Different types in one command
kubectl get pod/web-0 svc/web-svc deployment/web

# From multiple files
kubectl apply -f deployment.yaml -f service.yaml

# All YAML files in a directory
kubectl apply -f manifests/ -R
```

## API Groups

Kubernetes resources are organized into **API groups**, each with a group name and one or more versions. The core API has no group name (just `v1`).

```shell
kubectl api-versions        # list all API groups and versions
kubectl api-resources       # list all resource types with their API groups
kubectl api-resources --cached  # faster, uses cached discovery
```

### API Group Structure

```
GROUP/VERSION
─────────────
v1                              (core group, no group name)
apps/v1
batch/v1
networking.k8s.io/v1
rbac.authorization.k8s.io/v1
storage.k8s.io/v1
```

The `apiVersion` field in a manifest is `<group>/<version>` for named groups, or just `<version>` for the core group (`v1`).

## Common Resource Types

### Core API Group (`v1`)

| Resource | Short Name | Namespaced | Description |
|----------|------------|------------|-------------|
| `pods` | `po` | Yes | Smallest deployable unit |
| `services` | `svc` | Yes | Network abstraction over pods |
| `endpoints` | `ep` | Yes | IP addresses backing a service |
| `configmaps` | `cm` | Yes | Configuration data |
| `secrets` | — | Yes | Sensitive data (base64) |
| `namespaces` | `ns` | No | Virtual cluster boundary |
| `nodes` | `no` | No | Worker machine |
| `persistentvolumes` | `pv` | No | Cluster storage |
| `persistentvolumeclaims` | `pvc` | Yes | User's storage request |
| `replicationcontrollers` | `rc` | Yes | Legacy pod replication |
| `serviceaccounts` | `sa` | Yes | Pod identity |
| `events` | `ev` | Yes | Cluster events |
| `resourcequotas` | `quota` | Yes | Resource limits per namespace |
| `limitranges` | `limits` | Yes | Min/max constraints per namespace |
| `componentstatuses` | `cs` | No | Control plane health (deprecated) |
| `podtemplates` | — | Yes | Template for creating pods |

### `apps/v1` (Workloads)

| Resource | Short Name | Namespaced | Description |
|----------|------------|------------|-------------|
| `deployments` | `deploy` | Yes | Stateless workload with rolling updates |
| `replicasets` | `rs` | Yes | Maintains a stable set of pods |
| `daemonsets` | `ds` | Yes | One pod per node |
| `statefulsets` | `sts` | Yes | Stateful workload with stable identity |
| `controllerrevisions` | — | Yes | Rollout history for DaemonSets/StatefulSets |

### `batch/v1` (Jobs)

| Resource | Short Name | Namespaced | Description |
|----------|------------|------------|-------------|
| `jobs` | — | Yes | Runs pods to completion |
| `cronjobs` | `cj` | Yes | Scheduled jobs |

### `networking.k8s.io/v1`

| Resource | Short Name | Namespaced | Description |
|----------|------------|------------|-------------|
| `ingresses` | `ing` | Yes | HTTP routing to services |
| `ingressclasses` | — | No | Ingress controller type |
| `networkpolicies` | `netpol` | Yes | Pod network traffic rules |

### `rbac.authorization.k8s.io/v1`

| Resource | Short Name | Namespaced | Description |
|----------|------------|------------|-------------|
| `roles` | — | Yes | Namespace-scoped permissions |
| `rolebindings` | — | Yes | Binds roles to subjects |
| `clusterroles` | — | No | Cluster-wide permissions |
| `clusterrolebindings` | — | No | Binds cluster roles to subjects |

### `storage.k8s.io/v1`

| Resource | Short Name | Namespaced | Description |
|----------|------------|------------|-------------|
| `storageclasses` | `sc` | No | Storage provisioner type |
| `volumeattachments` | — | No | Volume attached to a node |
| `csidrivers` | — | No | CSI driver specification |
| `csinodes` | — | No | CSI node info |

### `autoscaling/v2`

| Resource | Short Name | Namespaced | Description |
|----------|------------|------------|-------------|
| `horizontalpodautoscalers` | `hpa` | Yes | Auto-scaling based on metrics |

### `policy/v1`

| Resource | Short Name | Namespaced | Description |
|----------|------------|------------|-------------|
| `poddisruptionbudgets` | `pdb` | Yes | Min available pods during disruptions |

### Other Notable Groups

| Resource | Short Name | API Group | Namespaced |
|----------|------------|-----------|------------|
| `customresourcedefinitions` | `crd`, `crds` | `apiextensions.k8s.io/v1` | No |
| `apiservices` | — | `apiregistration.k8s.io/v1` | No |
| `certificatesigningrequests` | `csr` | `certificates.k8s.io/v1` | No |
| `priorityclasses` | `pc` | `scheduling.k8s.io/v1` | No |
| `leases` | — | `coordination.k8s.io/v1` | Yes |
| `endpointslices` | — | `discovery.k8s.io/v1` | Yes |
| `runtimeclasses` | — | `node.k8s.io/v1` | No |

## Discovering Resources

### `kubectl api-resources`

Lists all resource types the API server knows about, including CRDs:

```shell
kubectl api-resources
kubectl api-resources --namespaced=true      # only namespaced resources
kubectl api-resources --namespaced=false     # only cluster-scoped
kubectl api-resources --api-group=apps       # filter by API group
kubectl api-resources --verbs=list           # resources supporting a verb
kubectl api-resources --cached               # use cached discovery (fast)
```

Output columns: `NAME`, `SHORTNAMES`, `APIVERSION`, `NAMESPACED`, `KIND`.

### `kubectl api-versions`

Lists all API group/version pairs:

```shell
kubectl api-versions
# Output:
# apps/v1
# batch/v1
# networking.k8s.io/v1
# ...
```

### `kubectl explain`

Shows the OpenAPI schema for a resource type — the field reference for writing manifests:

```shell
kubectl explain pods
kubectl explain pods.spec
kubectl explain pods.spec.containers
kubectl explain pods --recursive              # full field tree
kubectl explain deployment --api-version=apps/v1
kubectl explain pods --output=plaintext-openapiv2
```

`explain` reads from the API server's OpenAPI spec, so it includes CRDs and reflects the actual server version.

## Label Selectors

Label selectors filter resources by their `metadata.labels`. Used with `-l` / `--selector`:

```shell
kubectl get pods -l app=nginx
kubectl get pods -l 'app in (nginx, redis)'
kubectl get pods -l app!=nginx
kubectl get pods -l 'app,environment'          # both labels must exist
kubectl get pods -l app=nginx,env=prod          # AND logic (comma)
```

### Operators

| Operator | Example | Meaning |
|----------|---------|---------|
| `=` / `==` | `app=nginx` | Label equals value |
| `!=` | `app!=nginx` | Label does not equal value |
| `in` | `'app in (a, b)'` | Label value is in set |
| `notin` | `'app notin (a, b)'` | Label value is not in set |
| `exists` | `app` (key only) | Label key exists |
| `!` | `'!app'` | Label key does not exist |

Multiple selectors are **AND**-ed together (comma-separated).

### Label Selector in Manifests

Selectors also appear in resource specs:

```yaml
selector:
  matchLabels:
    app: nginx
  matchExpressions:
    - key: environment
      operator: In
      values: [prod, staging]
```

## Field Selectors

Field selectors filter by **resource fields** (not labels). Used with `--field-selector`:

```shell
kubectl get pods --field-selector=status.phase=Running
kubectl get pods --field-selector=spec.nodeName=worker-1
kubectl get pods --field-selector=status.phase!=Running,spec.restartPolicy=Always
kubectl get services --field-selector=spec.type=LoadBalancer
kubectl get nodes --field-selector=spec.unschedulable!=true
```

### Operators

- `=` / `==` — equal to
- `!=` — not equal to

Set-based operators (`in`, `notin`, `exists`) are **not** supported for field selectors.

### Supported Fields by Type

Only certain fields are indexed and support field selection. Unsupported fields produce an error.

| Resource | Supported Fields |
|----------|-----------------|
| Pod | `spec.nodeName`, `spec.restartPolicy`, `spec.schedulerName`, `spec.serviceAccountName`, `spec.hostNetwork`, `status.phase`, `status.podIP`, `status.nominatedNodeName` |
| Event | `involvedObject.kind`, `involvedObject.namespace`, `involvedObject.name`, `involvedObject.uid`, `reason`, `type`, `source` |
| Service | `spec.clusterIP`, `spec.type` |
| Namespace | `status.phase` |
| Node | `spec.unschedulable` |
| Secret | `type` |
| ReplicaSet / ReplicationController | `status.replicas` |
| Job | `status.successful` |
| All types | `metadata.name`, `metadata.namespace` |

### Cross-Type Field Selectors

Field selectors work across multiple resource types in one command:

```shell
kubectl get statefulsets,services --all-namespaces --field-selector metadata.namespace!=default
```

## Resource Version and Optimistic Concurrency

Every object has a `metadata.resourceVersion` — a string that changes on every write. kubectl and the API server use this for **optimistic concurrency control**:

1. Client reads object (gets `resourceVersion: "123"`)
2. Client sends update with `resourceVersion: "123"`
3. If another client updated first (version is now `"124"`), the server rejects with `409 Conflict`

Commands like `kubectl apply` and `kubectl patch` handle this automatically by retrying. For `kubectl replace`, you may need to handle conflicts manually.

## Custom Resource Definitions (CRDs)

CRDs extend the Kubernetes API with custom resource types. Once installed, they appear in `kubectl api-resources` and can be managed like built-in resources:

```shell
kubectl get crd                          # list CRDs
kubectl get certificates.cert-manager.io # list instances of a CRD
kubectl apply -f my-certificate.yaml     # create a CRD instance
kubectl explain certificates.cert-manager.io  # schema from CRD
```

CRDs follow the same `apiVersion: <group>/<version>` and `kind` conventions. Their short names (if defined) appear in `kubectl api-resources`.

## Edge Cases and Known Issues

- **Short name collisions**: Some short names can match multiple resources across API groups. Use the fully-qualified `TYPE.GROUP` form to disambiguate (e.g., `deployments.apps`).
- **`api-resources --cached` may be stale**: If a CRD was recently added, the cache may not reflect it. Run without `--cached` to force fresh discovery.
- **Field selectors are limited**: Only a small set of fields per resource type support field selection. You cannot field-select arbitrary spec fields.
- **`explain` depends on OpenAPI**: If the API server doesn't expose OpenAPI v2/v3, `explain` may fail or produce incomplete output.
- **Deleting a resource doesn't delete dependents by default**: Use `--cascade` to control dependent deletion (see [apply.md](apply.md)).
- **`kubectl get` without a type lists nothing**: Unlike some tools, `kubectl get` with no arguments does not list all resources — you must specify a type.

## References

- [kubectl Overview](https://kubernetes.io/docs/reference/kubectl/)
- [Labels and Selectors](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/)
- [Field Selectors](https://kubernetes.io/docs/concepts/overview/working-with-objects/field-selectors/)

## Related Skills

- For core Kubernetes concepts (pods, services, controllers), see [concepts.md](concepts.md).
- For output formatting including sort-by and custom columns, see [output.md](output.md).
- For declarative apply and patch types, see [apply.md](apply.md).
