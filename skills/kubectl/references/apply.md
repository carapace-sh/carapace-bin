# Declarative Resource Management

How `kubectl apply`, `diff`, `patch`, `replace`, `create`, and `delete` manage resources — the declarative vs imperative models, patch types, server-side apply, pruning, and kustomize integration.

> **Source of truth**: <https://kubernetes.io/docs/reference/kubectl/generated/kubectl_apply/>, <https://kubernetes.io/docs/reference/using-api/server-side-apply/>. For **core concepts** (spec vs status, reconciliation), see [concepts.md](concepts.md). For **resource types**, see [resources.md](resources.md).

## `kubectl apply`

The primary declarative command — applies a configuration to a resource, creating it if it doesn't exist or updating it to match the configuration.

```shell
kubectl apply -f deployment.yaml
kubectl apply -f manifests/ -R               # recursive directory
kubectl apply -k overlays/prod               # kustomize directory
kubectl apply -f https://raw.githubusercontent.com/.../manifest.yaml  # from URL
```

### How Apply Works (Client-Side)

1. kubectl reads the file and the live object from the cluster
2. Computes a **three-way diff**: last-applied-configuration vs live vs new config
3. Sends a **strategic merge patch** to the API server
4. Stores the new config in the `kubectl.kubernetes.io/last-applied-configuration` annotation

The three-way merge ensures that:
- Fields added by other actors (controllers, other kubectl commands) are preserved
- Fields removed from the config are removed from the live object
- Fields present in the config overwrite the live object

### Server-Side Apply

```shell
kubectl apply -f deployment.yaml --server-side
kubectl apply -f deployment.yaml --server-side --field-manager=my-controller
```

Server-side apply moves the merge logic to the API server:
- **Field ownership tracking**: The server tracks which manager owns each field
- **Conflict detection**: If another manager owns a field you're trying to set, the apply fails with a conflict
- **No `last-applied-configuration` annotation**: The server tracks ownership internally via `managedFields`

```shell
# Force overwrite conflicts
kubectl apply -f deployment.yaml --server-side --force-conflicts

# View field ownership
kubectl get deployment my-app -o jsonpath='{.metadata.managedFields}'
```

### Apply Flags

| Flag | Description |
|------|-------------|
| `-f`, `--filename` | File, directory, or URL with config |
| `-k`, `--kustomize` | Kustomization directory |
| `-R`, `--recursive` | Process directory recursively |
| `--server-side` | Run apply on the server |
| `--field-manager` | Name of field manager (default: `kubectl-client-side-apply` or `kubectl-edit`) |
| `--force-conflicts` | Force changes against conflicts (server-side only) |
| `--validate` | Validation mode: `strict`/`true`, `warn`, `ignore`/`false` |
| `--dry-run` | `none`, `server`, or `client` |
| `--prune` | Delete resources not in the config |
| `--prune-allowlist` | Override default prune allowlist |
| `--all` | Select all resources (with `--prune`) |
| `-l`, `--selector` | Label selector (with `--prune`) |
| `--cascade` | Deletion cascade: `background`, `orphan`, `foreground` |
| `--force` | Immediately remove (bypass graceful deletion) |
| `--grace-period` | Grace period in seconds for deletion |
| `--timeout` | Time to wait before giving up on delete |
| `--openapi-patch` | Use OpenAPI to calculate diff |
| `--overwrite` | Automatically resolve conflicts (client-side) |
| `--wait` | Wait for resources to be gone before returning |
| `--subresource` | Operate on a subresource (server-side only) |

### Pruning

`--prune` deletes resources that were previously applied but are no longer in the config:

```shell
kubectl apply -f manifests/ --prune --all
kubectl apply -f manifests/ --prune -l app=myapp
kubectl apply -f manifests/ --prune --prune-allowlist=apps/v1/Deployment,v1/Service
```

Pruning uses labels or `--all` to scope which resources to consider for deletion. The `--prune-allowlist` restricts pruning to specific resource types.

### Validation

```shell
kubectl apply -f deployment.yaml --validate=strict   # default: full schema validation
kubectl apply -f deployment.yaml --validate=warn      # warn on unknown fields
kubectl apply -f deployment.yaml --validate=ignore    # skip validation
```

Server-side validation requires `ServerSideFieldValidation` to be enabled on the API server. Otherwise, kubectl falls back to client-side validation.

### Dry Run

```shell
kubectl apply -f deployment.yaml --dry-run=client    # print what would be sent
kubectl apply -f deployment.yaml --dry-run=server    # server validates without persisting
kubectl apply -f deployment.yaml --dry-run=none       # actually apply (default)
```

## `kubectl diff`

Preview what `apply` would change without applying:

```shell
kubectl diff -f deployment.yaml
kubectl diff -f manifests/ -R
kubectl diff -k overlays/prod
```

Shows a unified diff between the live state and what would be applied. Uses server-side dry run.

## `kubectl create`

Imperatively creates a resource. Fails if the resource already exists (unlike `apply`).

```shell
kubectl create -f deployment.yaml
kubectl create -f deployment.yaml --dry-run=client -o yaml
kubectl create -f deployment.yaml --save-config        # store config for future apply
```

### `--save-config`

The `--save-config` flag stores the full configuration in the `kubectl.kubernetes.io/last-applied-configuration` annotation. This allows transitioning from `create` to `apply` without conflicts:

```shell
kubectl create -f deployment.yaml --save-config
# Later:
kubectl apply -f deployment.yaml     # works smoothly because last-applied exists
```

### `create` Subcommands

kubectl can generate resources without a manifest file:

```shell
kubectl create namespace dev
kubectl create configmap app-config --from-literal=key1=val1 --from-literal=key2=val2
kubectl create configmap app-config --from-file=config.properties
kubectl create secret generic db-secret --from-literal=password=s3cret
kubectl create secret tls tls-secret --cert=path/to/tls.crt --key=path/to/tls.key
kubectl create secret docker-registry regcred \
  --docker-server=ghcr.io \
  --docker-username=user \
  --docker-password=pass \
  --docker-email=email@example.com
kubectl create deployment web --image=nginx --replicas=3
kubectl create service clusterip my-svc --tcp=80:80
kubectl create serviceaccount my-sa
kubectl create role pod-reader --verb=get,list,watch --resource=pods
kubectl create rolebinding read-pods --role=pod-reader --serviceaccount=default:my-sa
kubectl create clusterrole node-reader --verb=get,list --resource=nodes
kubectl create clusterrolebinding read-nodes --clusterrole=node-reader --serviceaccount=default:my-sa
kubectl create job pi --image=perl:5.34 -- perl -Mbignum=bpi -wle 'print bpi(2000)'
kubectl create cronjob hourly-job --image=busybox --schedule="0 * * * *" -- /bin/sh -c "date"
kubectl create ingress web-ingress --rule="example.com/=web-svc:80"
kubectl create quota compute-quota --hard=cpu=2,memory=4Gi,pods=10
kubectl create priorityclass high-priority --value=1000
kubectl create poddisruptionbudget my-pdb --min-available=2 --selector=app=myapp
kubectl create token my-sa                          # service account token
```

## `kubectl delete`

Deletes resources. Supports deletion by name, file, label selector, or `--all`.

```shell
kubectl delete pod my-pod
kubectl delete -f deployment.yaml
kubectl delete pods -l app=old-app
kubectl delete pods --all -n dev
kubectl delete deployment web service web-svc   # multiple by type+name
kubectl delete namespace dev                     # deletes namespace and all contents
```

### Cascade Deletion

| Strategy | Description |
|----------|-------------|
| `background` (default) | API server deletes the object, then controllers garbage-collect dependents |
| `foreground` | API server marks the object for deletion, waits for dependents to be deleted, then deletes the object |
| `orphan` | Object is deleted but dependents are left orphaned |

```shell
kubectl delete deployment web --cascade=orphan
kubectl delete deployment web --cascade=foreground
kubectl delete deployment web --force --grace-period=0
```

### Force Deletion

```shell
kubectl delete pod stuck-pod --force --grace-period=0
```

Bypasses graceful deletion. May leave finalizers or cause inconsistency — use as a last resort.

## `kubectl replace`

Replaces a resource entirely from a file. Unlike `apply`, this replaces the full `spec` — fields not in the file are reset.

```shell
kubectl replace -f deployment.yaml
kubectl replace --force -f deployment.yaml       # delete then recreate
```

`replace` fails if the resource doesn't exist. It sends the full object, not a patch.

## `kubectl patch`

Updates specific fields of a resource without replacing the entire object. Three patch types:

```shell
kubectl patch deployment web -p '{"spec":{"replicas":5}}'
kubectl patch deployment web --type=merge -p '{"spec":{"replicas":5}}'
kubectl patch deployment web --type=json -p='[{"op":"replace","path":"/spec/replicas","value":5}]'
kubectl patch deployment web --patch-file patch.yaml
```

### Patch Types

| Type | Description | Format |
|------|-------------|--------|
| `strategic` (default) | Kubernetes-specific merge with field-type-aware merging | Partial JSON/YAML |
| `merge` | Standard JSON merge patch (RFC 7386) | Partial JSON/YAML |
| `json` | JSON Patch (RFC 6902) | Array of operations |

### Strategic Merge Patch

Kubernetes' default patch type. Uses field tags in the API type definitions to determine merge behavior:

- **Scalar fields**: Replaced
- **List fields with `patchStrategy: merge`**: Merged by a merge key (e.g., `name` for containers)
- **List fields with `patchStrategy: replace`**: Replaced entirely

```shell
# Update replicas (scalar)
kubectl patch deployment web -p '{"spec":{"replicas":5}}'

# Update container image (list merged by name)
kubectl patch deployment web -p '{"spec":{"template":{"spec":{"containers":[{"name":"nginx","image":"nginx:1.25"}]}}}}'
```

The strategic merge knows that `containers` is a list keyed by `name`, so it updates the container named `nginx` without affecting other containers.

### JSON Merge Patch

Standard RFC 7386 merge. Lists are always replaced entirely (no merge-key awareness):

```shell
kubectl patch deployment web --type=merge -p '{"spec":{"replicas":5}}'
```

Use this when strategic merge's behavior is too "smart" or for resources that don't have strategic merge tags (e.g., CRDs).

### JSON Patch

RFC 6902 — array of operations with explicit paths:

```shell
kubectl patch deployment web --type=json -p='[{"op":"replace","path":"/spec/replicas","value":5}]'
kubectl patch deployment web --type=json -p='[{"op":"add","path":"/spec/template/spec/containers/0/env","value":[{"name":"DEBUG","value":"true"}]}]'
kubectl patch deployment web --type=json -p='[{"op":"remove","path":"/spec/template/spec/containers/0/env/0"}]'
```

Operations: `add`, `remove`, `replace`, `move`, `copy`, `test`.

### Patch Flags

| Flag | Description |
|------|-------------|
| `-p`, `--patch` | Patch string (JSON or YAML) |
| `--patch-file` | Patch from a file |
| `--type` | Patch type: `strategic`, `merge`, `json` |
| `--field-manager` | Field manager name for server-side tracking |
| `--dry-run` | `none`, `server`, `client` |
| `--subresource` | Patch a subresource (e.g., `status`) |
| `--local` | Operate on file content, not server-side resource |
| `-f`, `--filename` | Patch resource identified by file |

## `kubectl edit`

Opens the live resource in your default editor, and applies changes on save:

```shell
kubectl edit deployment web
kubectl edit configmap app-config -o yaml
KUBE_EDITOR=nano kubectl edit deployment web
```

Uses the `KUBE_EDITOR` or `EDITOR` environment variable. Changes are applied via `apply` under the hood.

## `kubectl kustomize`

Builds a kustomization directory and outputs the resulting manifests:

```shell
kubectl kustomize overlays/prod
kubectl kustomize overlays/prod -o rendered.yaml
kubectl kustomize . --enable-helm
```

For standalone Kustomize details, see the [Kustomize docs](https://kubectl.docs.kubernetes.io/).

## `kubectl replace --force`

Deletes and recreates a resource — useful when a field is immutable:

```shell
kubectl replace --force -f service.yaml
```

This is a delete-then-create, not a patch. The resource gets a new `uid` and `resourceVersion`.

## Edge Cases and Known Issues

- **`apply` on imperatively created resources**: If a resource was created with `create` (not `apply`), the `last-applied-configuration` annotation is missing. The first `apply` may not prune fields correctly. Fix: `kubectl apply --force` or create with `--save-config`.
- **Strategic merge on CRDs**: CRDs don't have strategic merge tags, so `strategic` patch behaves like `merge` patch for them. Use `--type=merge` explicitly.
- **`apply` and lists**: Strategic merge can merge lists by merge key (e.g., `containers` by `name`). JSON merge always replaces lists entirely. This is a common source of confusion.
- **`--prune` scope**: Pruning only considers resources matching `--selector` or `--all` and in the `--prune-allowlist`. Resources outside the scope are not pruned.
- **`delete --force` may leave finalizers**: Force deletion bypasses finalizers, which can leave resources stuck in a `Terminating` state or cause controller inconsistency.
- **`patch --local` doesn't contact the server**: It operates on the file content only, useful for testing patch syntax.
- **Server-side apply conflicts**: When multiple managers own overlapping fields, `apply --server-side` fails. Use `--force-conflicts` to take ownership.
- **`replace` resets unspecified fields**: Unlike `apply`, `replace` sends the full object. Fields not in the file are reset to their zero values (or removed).
- **`edit` applies on save**: If the resource changed since you opened the editor, the apply may fail with a conflict. Re-run `edit` to retry.

## References

- [kubectl apply](https://kubernetes.io/docs/reference/kubectl/generated/kubectl_apply/)
- [Server-Side Apply](https://kubernetes.io/docs/reference/using-api/server-side-apply/)
- [Strategic Merge Patch](https://kubernetes.io/docs/tasks/manage-kubernetes-objects/update-api-object-kubectl-patch/)
- [JSON Patch (RFC 6902)](https://datatracker.ietf.org/doc/html/rfc6902)
- [JSON Merge Patch (RFC 7386)](https://datatracker.ietf.org/doc/html/rfc7386)

## Related Skills

- For core concepts (spec vs status, reconciliation), see [concepts.md](concepts.md).
- For resource types and naming, see [resources.md](resources.md).
- For workload-specific commands (scale, rollout), see [workloads.md](workloads.md).
