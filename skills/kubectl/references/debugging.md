# Debugging and Troubleshooting Commands

Commands for inspecting and debugging resources — `describe`, `logs`, `exec`, `attach`, `cp`, `port-forward`, `proxy`, `top`, `events`, `debug`, and `wait`.

> **Source of truth**: `kubectl <command> --help`, <https://kubernetes.io/docs/reference/kubectl/generated/kubectl_logs/>. For **cluster management** (cordon, drain, taint), see [cluster.md](cluster.md). For **output formatting**, see [output.md](output.md).

## `kubectl describe`

Shows detailed information about a resource, including related events:

```shell
kubectl describe pod my-pod
kubectl describe node worker-1
kubectl describe deployment web
kubectl describe -f pod.yaml              # describe resource from file
kubectl describe pods -l app=nginx        # describe all matching pods
```

`describe` output is human-readable and includes:
- Resource spec and status
- Events for that resource (from the event store)
- Related resources (e.g., for a pod: its node, volumes, events)

## `kubectl logs`

Prints logs for a container in a pod:

```shell
kubectl logs my-pod
kubectl logs my-pod -c my-container
kubectl logs -f my-pod                          # follow (tail)
kubectl logs --previous my-pod                  # previous (crashed) container
kubectl logs my-pod --tail=50                   # last 50 lines
kubectl logs my-pod --since=1h                  # logs from last hour
kubectl logs my-pod --since-time=2024-01-01T00:00:00Z
kubectl logs -l app=nginx --all-containers      # all pods matching label
kubectl logs -l app=nginx --all-pods --prefix   # all pods, prefix with source
```

### Flags

| Flag | Description |
|------|-------------|
| `-c`, `--container` | Container name (defaults to first or `kubectl.kubernetes.io/default-container` annotation) |
| `-f`, `--follow` | Stream logs in real-time |
| `-p`, `--previous` | Logs from the previous (crashed) container instance |
| `--tail` | Number of lines from the end (default: all) |
| `--since` | Relative duration (e.g., `5s`, `2m`, `1h`) |
| `--since-time` | RFC3339 timestamp |
| `--timestamps` | Prepend timestamp to each line |
| `--limit-bytes` | Maximum bytes of logs |
| `--all-containers` | All containers in the pod(s) |
| `--all-pods` | All pods matching selector (sets `--prefix`) |
| `--prefix` | Prefix each line with pod/container name |
| `-l`, `--selector` | Label selector to filter pods |
| `--max-log-requests` | Max concurrent log requests (default: 5) |
| `--ignore-errors` | Non-fatal errors during streaming |
| `--insecure-skip-tls-verify-backend` | Skip kubelet TLS verification |
| `--pod-running-timeout` | Wait for pod to be running (e.g., `5m`) |

### Multi-Container Pods

If a pod has multiple containers, specify `-c` or use the `kubectl.kubernetes.io/default-container` annotation:

```shell
kubectl logs my-pod -c sidecar
kubectl logs my-pod --all-containers
```

### Previous Container Logs

When a container crashes and restarts, the previous instance's logs are still available:

```shell
kubectl logs --previous my-pod
kubectl logs -p my-pod -c my-container
```

## `kubectl exec`

Executes a command in a container:

```shell
kubectl exec my-pod -- ls /app
kubectl exec my-pod -c my-container -- ls /app
kubectl exec -it my-pod -- /bin/sh           # interactive shell
kubectl exec -it my-pod -- /bin/bash
kubectl exec my-pod -- env                   # print environment
```

### Flags

| Flag | Description |
|------|-------------|
| `-c`, `--container` | Container name |
| `-i`, `--stdin` | Pass stdin to the container |
| `-t`, `--tty` | Stdin is a TTY (allocate a pseudo-terminal) |
| `-q`, `--quiet` | Only print output from the remote session |
| `-f`, `--filename` | Resource to exec into, identified by file |
| `--pod-running-timeout` | Wait for pod to be running |
| `--` | Separator: everything after is the command and its args |

### Interactive Shell

```shell
kubectl exec -it my-pod -- /bin/sh
kubectl exec -it my-pod -- /bin/bash
```

The `--` separates kubectl flags from the command. It's required when the command has flags that could be confused with kubectl flags:

```shell
kubectl exec my-pod -- ls -la /app
```

### Multi-Container Pods

```shell
kubectl exec my-pod -c sidecar -- ls /app
# Or set default-container annotation:
kubectl annotate pod my-pod kubectl.kubernetes.io/default-container=sidecar
kubectl exec my-pod -- ls /app     # now defaults to sidecar
```

## `kubectl attach`

Attaches to a process that is already running inside a container:

```shell
kubectl attach my-pod -it
kubectl attach my-pod -c my-container -it
```

Unlike `exec` which starts a new process, `attach` connects to the container's primary process (PID 1 or the entrypoint). Useful for interactive processes.

## `kubectl cp`

Copies files and directories between a container and the local filesystem:

```shell
kubectl cp my-pod:/app/logs.txt ./logs.txt
kubectl cp ./config.yaml my-pod:/app/config.yaml
kubectl cp my-pod:/app/logs ./logs -c my-container
kubectl cp ./data my-pod:/app/data --no-preserve
```

### File-Spec Syntax

```
[NAMESPACE/]POD_NAME:PATH   (container source)
LOCAL_PATH                  (local source/destination)
```

For pods in other namespaces: `kubectl cp dev/my-pod:/app/data ./data`

### Flags

| Flag | Description |
|------|-------------|
| `-c`, `--container` | Container name |
| `--no-preserve` | Don't preserve ownership/permissions |
| `--retries` | Number of retries (default: 0) |

### How It Works

`cp` uses `tar` under the hood — it runs `tar` inside the container to create/extract an archive streamed over the API. The container must have `tar` installed.

## `kubectl port-forward`

Forwards local ports to a pod — useful for accessing services without exposing them:

```shell
kubectl port-forward pod/my-pod 8080:80
kubectl port-forward pod/my-pod 8080                # same local and remote port
kubectl port-forward svc/my-svc 8080:80             # forwards to a service endpoint
kubectl port-forward deployment/web 8080:80         # forwards to a pod in the deployment
kubectl port-forward pod/my-pod 8080 8081 8082      # multiple ports
kubectl port-forward pod/my-pod 8080:80 9090:9090   # multiple port mappings
kubectl port-forward pod/my-pod 8080:80 --address=0.0.0.0  # bind to all interfaces
kubectl port-forward pod/my-pod 8080:80 --address=0.0.0.0,192.168.1.100
```

### Flags

| Flag | Description |
|------|-------------|
| `--address` | Addresses to listen on (comma-separated, IPs or `localhost`) |
| `--pod-running-timeout` | Wait for pod to be running |

### Port Syntax

```
LOCAL_PORT:REMOTE_PORT    # map specific ports
LOCAL_PORT                # same port locally and remotely
```

### Limitations

- Port-forward connects to a **single pod**, not a load-balanced set
- For Services, it picks one endpoint pod
- The connection is tunneled through the API server — not a direct connection
- If the pod restarts, the port-forward breaks

## `kubectl proxy`

Runs a proxy server to the Kubernetes API server, making the API available on localhost:

```shell
kubectl proxy                                    # http://localhost:8001
kubectl proxy --port=9090                        # custom port
kubectl proxy --api-prefix=/custom               # custom API prefix
kubectl proxy --www=./static/                    # serve static files
kubectl proxy --address=0.0.0.0 --accept-hosts=.* # allow external access
```

The proxy allows accessing the API via HTTP without authentication (kubectl handles auth). Useful for:
- Browsing the API with a browser
- Accessing service proxy endpoints (`/api/v1/namespaces/{ns}/services/{svc}:port/proxy/`)
- Debugging API responses

## `kubectl top`

Shows CPU, memory, and storage usage for nodes or pods. Requires **metrics-server** to be installed:

```shell
kubectl top nodes
kubectl top nodes --sort-by=memory
kubectl top pods
kubectl top pods -A                              # all namespaces
kubectl top pods --sort-by=cpu
kubectl top pod my-pod -c                        # per-container breakdown
kubectl top pods -l app=nginx                    # filtered by label
kubectl top pod my-pod --containers
```

### Flags

| Flag | Description |
|------|-------------|
| `-A`, `--all-namespaces` | All namespaces (pods only) |
| `--sort-by` | Sort by `cpu` or `memory` |
| `-l`, `--selector` | Label selector |
| `--containers` | Show per-container usage (pods) |
| `--no-headers` | Don't print headers |

### Units

- CPU: millicores (`m`) — `500m` = 0.5 cores
- Memory: mebibytes (`Mi`), gibibytes (`Gi`)
- Values are current usage, not percentage of limit

## `kubectl events`

Lists cluster events (similar to `kubectl get events` but with better default formatting):

```shell
kubectl events
kubectl events -A
kubectl events --for pod/my-pod                  # events for a specific resource
kubectl events --types=Warning                   # filter by type
kubectl events --watch
kubectl events --sort-by=.lastTimestamp
```

## `kubectl debug`

Creates debugging sessions for troubleshooting workloads and nodes. Adds an **ephemeral container** to a running pod, or creates a copy of a pod with debugging tools:

```shell
kubectl debug my-pod --image=busybox --target=my-container
kubectl debug my-pod -it --image=ubuntu --target=my-container
kubectl debug my-pod --image=busybox --copy-to=my-pod-debug --share-processes
kubectl debug node/worker-1 -it --image=ubuntu
kubectl debug my-pod --profile=sysadmin --image=debian
```

### Flags

| Flag | Description |
|------|-------------|
| `-c`, `--container` | Name for the debug container |
| `--image` | Container image for debugging |
| `--image-pull-policy` | Pull policy: `Always`, `Never`, `IfNotPresent` |
| `--target` | Target container to share process namespace with |
| `--copy-to` | Create a copy of the pod with this name |
| `--replace` | Delete the original pod when using `--copy-to` |
| `--same-node` | Schedule the copy on the same node |
| `--share-processes` | Enable process namespace sharing in the copy |
| `--profile` | Debug profile: `general`, `baseline`, `restricted`, `netadmin`, `sysadmin` |
| `--attach` | Wait for container and attach |
| `-i`, `--stdin` | Keep stdin open |
| `-t`, `--tty` | Allocate a TTY |
| `--quiet` | Suppress informational messages |
| `--set-image` | Change container images in the copy (`name=image,...`) |
| `--env` | Environment variables for the debug container |

### Debug Profiles

| Profile | Capabilities |
|---------|-------------|
| `general` | Default — standard debugging tools |
| `baseline` | Minimum privileges |
| `restricted` | Restricted (security-hardened) |
| `netadmin` | Network administration capabilities |
| `sysadmin` | Full system administration capabilities |

### Ephemeral Containers

Ephemeral containers are a special container type for debugging:
- They don't restart if they exit
- They're not part of the pod's resource guarantees
- They can share process namespace with a target container (`--target`)
- They're added to an already-running pod (no restart needed)

## `kubectl wait`

Waits for a specific condition on one or more resources:

```shell
kubectl wait --for=condition=Ready pod/my-pod
kubectl wait --for=condition=Ready pod/my-pod --timeout=5m
kubectl wait --for=condition=Available deployment/web
kubectl wait --for=delete pod/old-pod
kubectl wait --for=create -f pod.yaml
kubectl wait --for=jsonpath='{.status.phase}'=Running pod/my-pod
kubectl wait --for=condition=Ready pods --all --timeout=10m
kubectl wait --for=condition=Ready pod -l app=web
```

### `--for` Syntax

| Form | Description |
|------|-------------|
| `--for=condition=<name>` | Wait for a condition to be true |
| `--for=condition=<name>=<value>` | Wait for condition with specific value |
| `--for=create` | Wait for resource to be created |
| `--for=delete` | Wait for resource to be deleted |
| `--for=jsonpath='{.path}'=value` | Wait for a JSONPath field to equal a value |

### Flags

| Flag | Description |
|------|-------------|
| `--for` | The condition to wait on (repeatable, AND-ed) |
| `--timeout` | Max wait time (`0` = check once, negative = 1 week) |
| `--local` | Don't contact API server (local only) |
| `-f`, `--filename` | Resource identified by file |
| `-l`, `--selector` | Label selector |
| `--all` | All resources of the type |
| `-A`, `--all-namespaces` | Across all namespaces |
| `--field-selector` | Field selector |

### Multiple Conditions

```shell
kubectl wait --for=condition=Ready --for=condition=Initialized pod/my-pod
```

Conditions are AND-ed and checked in order. `--for=create` is always checked first.

## Debugging Workflow

### Pod Not Starting

```shell
kubectl get pod my-pod                          # check status
kubectl describe pod my-pod                     # events, conditions, scheduling
kubectl get events --for pod/my-pod --sort-by=.lastTimestamp
```

Common issues visible in `describe`:
- `Pending` — not scheduled (check node resources, taints, affinity)
- `ImagePullBackOff` — image can't be pulled (check image name, registry auth)
- `CrashLoopBackOff` — container keeps crashing (check `logs --previous`)

### Container Crashing

```shell
kubectl logs my-pod                            # current container logs
kubectl logs --previous my-pod                  # previous instance before crash
kubectl logs -f my-pod                          # follow in real-time
kubectl exec my-pod -- ls /app                  # inspect filesystem
kubectl exec my-pod -- env                      # check environment variables
kubectl debug my-pod -it --image=busybox --target=my-container  # debug container
```

### Pod Running but Not Accessible

```shell
kubectl get endpoints my-svc                    # check service endpoints
kubectl describe svc my-svc                     # check selector and ports
kubectl exec my-pod -- wget -qO- http://service-name:port  # test from inside
kubectl port-forward svc/my-svc 8080:80         # test from local machine
kubectl get networkpolicy                        # check network policies
```

### Node Issues

```shell
kubectl describe node worker-1                  # conditions, capacity, events
kubectl top node worker-1                        # resource usage
kubectl get events --field-selector involvedObject.kind=Node
kubectl get pods --field-selector spec.nodeName=worker-1   # pods on the node
kubectl debug node/worker-1 -it --image=ubuntu   # get a shell on the node
```

## Edge Cases and Known Issues

- **`logs` requires a running pod**: If the pod hasn't started, `logs` returns an error. Use `--pod-running-timeout` to wait.
- **`logs --previous` only works for restarted containers**: If the container exited cleanly and wasn't restarted, `--previous` may not have logs.
- **`exec` into a container requires `tar` for `cp`**: `kubectl cp` runs `tar` inside the container. Minimal images (distroless, scratch) may not have `tar`.
- **`port-forward` to a service picks one pod**: It doesn't load-balance. If the pod dies, the forward breaks.
- **`top` requires metrics-server**: Without the metrics-server addon, `top` returns an error. Install it separately if needed.
- **`debug` ephemeral containers version requirements**: Ephemeral containers were alpha until Kubernetes 1.22, beta from 1.23 (enabled by default), and GA from 1.28. Clusters older than 1.23 may not support `kubectl debug`.
- **`wait --for=jsonpath` needs exact match**: The JSONPath value must match exactly (after Unicode case folding). No regex.
- **`wait --timeout=0` checks once**: A zero timeout means "check once and return immediately" — it does not wait.
- **`describe` events are limited**: The events section shows recent events only. Old events may have been garbage-collected (default TTL: 1 hour).
- **`cp` to/from namespaces**: Use `namespace/pod:path` syntax: `kubectl cp dev/my-pod:/app/data ./data`.

## References

- [kubectl logs](https://kubernetes.io/docs/reference/kubectl/generated/kubectl_logs/)
- [kubectl exec](https://kubernetes.io/docs/reference/kubectl/generated/kubectl_exec/)
- [kubectl debug](https://kubernetes.io/docs/reference/kubectl/generated/kubectl_debug/)
- [Ephemeral Containers](https://kubernetes.io/docs/concepts/workloads/pods/ephemeral-containers/)

## Related Skills

- For cluster management (cordon, drain, taint), see [cluster.md](cluster.md).
- For output formatting (jsonpath for inspecting status), see [output.md](output.md).
- For workload management (rollout status, scale), see [workloads.md](workloads.md).
