# Workload Management Commands

Commands for managing Kubernetes workloads — `run`, `expose`, `scale`, `autoscale`, `rollout`, and the `set` family. Covers deployment lifecycle, rolling updates, and rollback.

> **Source of truth**: `kubectl <command> --help`, <https://kubernetes.io/docs/reference/kubectl/generated/kubectl_rollout/>. For **declarative apply** (alternative to these imperative commands), see [apply.md](apply.md). For **core workload concepts**, see [concepts.md](concepts.md).

## `kubectl run`

Creates and runs a particular image, potentially replicating it. Creates a Pod or Deployment depending on flags.

```shell
kubectl run nginx --image=nginx
kubectl run nginx --image=nginx --replicas=3            # create Deployment (deprecated for run)
kubectl run nginx --image=nginx --port=80
kubectl run nginx --image=nginx --env=KEY=value
kubectl run nginx --image=nginx --dry-run=client -o yaml > pod.yaml
kubectl run nginx --image=nginx --restart=OnFailure     # create a Job
kubectl run nginx --image=nginx --schedule="0 * * * *"  # create a CronJob
kubectl run nginx --image=nginx --command -- nginx -g "daemon off;"
```

### Common Flags

| Flag | Description |
|------|-------------|
| `--image` | Container image to run |
| `--port` | Container port to expose |
| `--replicas` | Number of replicas (creates Deployment) |
| `--restart` | Restart policy: `Always`, `OnFailure`, `Never` |
| `--schedule` | Cron schedule (creates CronJob) |
| `--command` | Use args as command, not args to the image entrypoint |
| `--env` | Environment variables (`KEY=value`, repeatable) |
| `--overrides` | JSON override for the generated spec |
| `--dry-run` | `client` or `server` |
| `-o`, `--output` | Output format (useful with `--dry-run=client`) |
| `--labels`, `-l` | Labels to apply |
| `--serviceaccount` | Service account to use |
| `--limits` | Resource limits (e.g., `cpu=200m,memory=512Mi`) |
| `--requests` | Resource requests |

## `kubectl expose`

Exposes a resource as a new Kubernetes Service. The resource must have a selector.

```shell
kubectl expose deployment web --port=80 --target-port=80
kubectl expose deployment web --port=80 --target-port=80 --type=LoadBalancer
kubectl expose deployment web --port=80 --target-port=8080 --name=web-svc
kubectl expose pod nginx --port=80 --type=NodePort
kubectl expose rs frontend --port=80 --target-port=8080
```

### Common Flags

| Flag | Description |
|------|-------------|
| `--port` | Port the service will serve on |
| `--target-port` | Port on the pod the service forwards to (defaults to `--port`) |
| `--type` | Service type: `ClusterIP`, `NodePort`, `LoadBalancer`, `ExternalName` |
| `--name` | Name of the new service |
| `--protocol` | IP protocol: `TCP` (default), `UDP`, `SCTP` |
| `--selector` | Label selector (defaults to the resource's selector) |
| `--session-affinity` | `None` (default) or `ClientIP` |
| `--external-ip` | External IP to accept traffic on |
| `--load-balancer-ip` | Cloud load balancer IP |
| `--cluster-ip` | Specific ClusterIP to request |
| `--labels`, `-l` | Labels for the service |

## `kubectl scale`

Changes the number of replicas for a Deployment, ReplicaSet, ReplicationController, or StatefulSet:

```shell
kubectl scale deployment web --replicas=5
kubectl scale deployment/web --replicas=3
kubectl scale rs frontend --replicas=0
kubectl scale statefulset db --replicas=3
kubectl scale deployment web --replicas=5 --current-replicas=3   # only if currently 3
```

### Conditional Scaling

`--current-replicas` makes scaling conditional — the scale only happens if the current replica count matches. This enables safe concurrent scaling:

```shell
kubectl scale deployment web --replicas=5 --current-replicas=3
# Scales to 5 only if currently at 3
```

### Flags

| Flag | Description |
|------|-------------|
| `--replicas` | Target replica count (required) |
| `--current-replicas` | Only scale if current count matches (precondition) |
| `--resource-version` | Only scale if resource version matches (optimistic concurrency) |
| `-l`, `--selector` | Label selector to find resources to scale |
| `--all` | Scale all resources of the type in the namespace |

## `kubectl autoscale`

Creates a HorizontalPodAutoscaler (HPA) that automatically scales a workload based on CPU/memory usage or custom metrics:

```shell
kubectl autoscale deployment web --min=2 --max=10
kubectl autoscale deployment web --min=2 --max=10 --cpu-percent=80
kubectl autoscale deployment web --min=2 --max=10 --cpu-percent=80 --name=web-hpa
```

### Flags

| Flag | Description |
|------|-------------|
| `--min` | Minimum number of replicas |
| `--max` | Maximum number of replicas (required) |
| `--cpu-percent` | Target CPU utilization percentage (default: 80) |
| `--name` | Name of the HPA (defaults to the workload name) |
| `--dry-run` | `client` or `server` |

The HPA is a separate resource — manage it with `kubectl get hpa`, `kubectl describe hpa`, `kubectl delete hpa`.

## `kubectl rollout`

Manages the rollout of a resource — view history, undo, pause, resume, restart, and check status.

### Subcommands

| Subcommand | Description |
|------------|-------------|
| `rollout history` | View rollout revision history |
| `rollout pause` | Pause a rollout (freeze updates) |
| `rollout resume` | Resume a paused rollout |
| `rollout undo` | Roll back to a previous revision |
| `rollout status` | Watch rollout status until complete |
| `rollout restart` | Restart the rollout (trigger pod recreation) |

### `rollout history`

```shell
kubectl rollout history deployment/web
kubectl rollout history deployment/web --revision=2
```

Shows revision numbers and change causes. Use `--record` (deprecated) or server-side apply with field managers to track change causes.

### `rollout undo`

```shell
kubectl rollout undo deployment/web                    # roll back to previous revision
kubectl rollout undo deployment/web --to-revision=2    # roll back to specific revision
```

Creates a new ReplicaSet with the old spec. The previous ReplicaSet is preserved for future rollbacks.

### `rollout status`

```shell
kubectl rollout status deployment/web
kubectl rollout status deployment/web --timeout=5m
```

Watches the rollout until it completes or fails. Blocks until the rollout is done (or timeout).

### `rollout pause` / `rollout resume`

```shell
kubectl rollout pause deployment/web
# Make multiple changes without triggering a rollout each time:
kubectl set image deployment/web nginx=nginx:1.25
kubectl set resources deployment/web -c=nginx --limits=cpu=500m
kubectl rollout resume deployment/web   # now triggers one rollout
```

Pausing allows making multiple changes that trigger only a single rollout when resumed.

### `rollout restart`

```shell
kubectl rollout restart deployment/web
kubectl rollout restart daemonset/log-agent
```

Triggers a rolling restart by modifying the pod template (adds an annotation with a timestamp). Useful for picking up config changes, secret rotations, or forced pod recreation.

## `kubectl set`

Modifies specific properties of workload resources:

### Subcommands

| Subcommand | Description |
|------------|-------------|
| `set env` | Update environment variables on a pod template |
| `set image` | Update container images |
| `set resources` | Update resource requests/limits |
| `set selector` | Set the label selector on a resource |
| `set serviceaccount` | Update the service account |
| `set subject` | Update subjects of a RoleBinding or ClusterRoleBinding |

### `set image`

```shell
kubectl set image deployment/web nginx=nginx:1.25
kubectl set image deployment/web nginx=nginx:1.25 redis=redis:7.0
kubectl set image deployment/web *=nginx:1.25               # all containers
kubectl set image deployment/web -c=nginx --local -o yaml   # preview
```

Triggers a rolling update for Deployments, StatefulSets, and DaemonSets.

### `set env`

```shell
kubectl set env deployment/web DEBUG=true
kubectl set env deployment/web DEBUG-                      # remove env var
kubectl set env deployment/web --from=configmap/app-config # from ConfigMap
kubectl set env deployment/web --from=secret/db-secret      # from Secret
kubectl set env deployment/web --list                       # list current env
kubectl set env deployment/web KEY1=val1 KEY2=val2          # multiple
```

### `set resources`

```shell
kubectl set resources deployment/web -c=nginx --limits=cpu=500m,memory=1Gi
kubectl set resources deployment/web -c=nginx --requests=cpu=100m,memory=256Mi
kubectl set resources deployment/web -c=nginx --limits=cpu=500m --requests=cpu=200m
```

### `set selector`

```shell
kubectl set selector service my-svc app=myapp
```

Only works on Services (selector is immutable for most other resources after creation).

### `set serviceaccount`

```shell
kubectl set serviceaccount deployment web my-sa
kubectl set serviceaccount deployment web my-sa --all
```

### `set subject`

```shell
kubectl set subject rolebinding my-rb --user=alice --group=devs
kubectl set subject clusterrolebinding my-crb --serviceaccount=default:my-sa
```

## `kubectl label`

Add, update, or remove labels on resources:

```shell
kubectl label pods my-pod environment=production
kubectl label pods my-pod environment=staging --overwrite        # update existing
kubectl label pods my-pod environment-                            # remove label
kubectl label pods --all environment=production                   # all pods
kubectl label pods -l app=nginx tier=frontend                     # label matching pods
kubectl label deployment web app.kubernetes.io/version=v2
```

### Flags

| Flag | Description |
|------|-------------|
| `--overwrite` | Allow updating existing label value |
| `--all` | Select all resources of the type |
| `-l`, `--selector` | Label selector for resources to modify |
| `--resource-version` | Only update if resource version matches |
| `--list` | List current labels (read-only) |

## `kubectl annotate`

Add, update, or remove annotations on resources (same syntax as `label`):

```shell
kubectl annotate pods my-pod description="my test pod"
kubectl annotate pods my-pod description="updated" --overwrite
kubectl annotate pods my-pod description-
kubectl annotate deployment web kubernetes.io/change-cause="update to v2"
```

## Workload Lifecycle

### Deployment Rolling Update Flow

```
kubectl set image deployment/web nginx=nginx:1.25
  → Deployment spec.template updated
    → Deployment controller creates new ReplicaSet (v2)
      → v2 ReplicaSet scales up (maxSurge)
      → v1 ReplicaSet scales down (maxUnavailable)
        → Rolling update completes
          → v1 ReplicaSet scaled to 0 (preserved for rollback)
```

### Rolling Update Strategy

Controlled by `spec.strategy.rollingUpdate`:

```yaml
strategy:
  type: RollingUpdate
  rollingUpdate:
    maxUnavailable: 25%       # max pods that can be unavailable during update
    maxSurge: 25%             # max pods above desired count during update
```

### Rollback Flow

```
kubectl rollout undo deployment/web
  → Deployment spec.template reverted to previous ReplicaSet
    → New ReplicaSet (old spec) scales up
      → Current ReplicaSet scales down
```

## Edge Cases and Known Issues

- **`kubectl run --replicas` is deprecated**: Using `run` with `--replicas` creates a Deployment, but this behavior is deprecated. Use `kubectl create deployment` instead.
- **`rollout pause` blocks all template changes**: While paused, changes to the pod template accumulate but don't trigger a rollout. Resume with `rollout resume`.
- **`set image` triggers immediate rollout**: Unlike `set` under `rollout pause`, a standalone `set image` triggers a rolling update immediately.
- **`scale` doesn't change the spec**: Scaling changes `spec.replicas`, not the pod template. It doesn't trigger a rolling update.
- **StatefulSet rollout is ordered**: StatefulSets update pods in reverse ordinal order (N-1, N-2, ..., 0) by default. Use `partition` for canary rollouts.
- **`rollout undo` to a deleted revision fails**: If a ReplicaSet was garbage-collected, its revision is gone. Keep enough revision history limit (`spec.revisionHistoryLimit`).
- **`autoscale` requires metrics-server**: HPA based on CPU/memory requires the metrics-server to be installed. Custom metrics require a custom metrics adapter.
- **`expose` requires a selector**: The target resource must have labels that the service can select. `expose` without a selector on the source resource creates a service with no endpoints.
- **`label --overwrite` is required for existing labels**: Without `--overwrite`, updating an existing label value fails.
- **`rollout restart` adds an annotation**: It works by adding `kubectl.kubernetes.io/restartedAt` to the pod template, which triggers a rolling update.

## References

- [kubectl rollout](https://kubernetes.io/docs/reference/kubectl/generated/kubectl_rollout/)
- [Deployments](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/)
- [Horizontal Pod Autoscaling](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/)

## Related Skills

- For declarative apply (alternative to imperative commands), see [apply.md](apply.md).
- For core workload concepts, see [concepts.md](concepts.md).
- For debugging workloads (logs, exec, describe), see [debugging.md](debugging.md).
