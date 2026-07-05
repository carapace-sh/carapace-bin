# Output Formats and Templates

How to control kubectl output with `-o`/`--output`, `--sort-by`, and custom columns. Covers JSONPath syntax, Go templates, and custom-columns in depth.

> **Source of truth**: <https://kubernetes.io/docs/reference/kubectl/jsonpath/>, <https://kubernetes.io/docs/reference/kubectl/#output-options>. For **field selectors** (filtering, not formatting), see [resources.md](resources.md).

## Output Formats

The `-o` / `--output` flag controls how kubectl formats results:

| Format | Description |
|--------|-------------|
| (default) | Plain-text table with server-defined columns |
| `wide` | Plain-text table with extra columns (e.g., node name, IP for pods) |
| `json` | Full API object as JSON |
| `yaml` | Full API object as YAML |
| `kyaml` | Kubernetes-specific YAML dialect (beta) |
| `name` | Only `resource/name` (e.g., `pod/nginx`) |
| `go-template` | Go `text/template` formatting |
| `go-template-file` | Go template from a file |
| `template` | Alias for `go-template` |
| `templatefile` | Alias for `go-template-file` |
| `jsonpath` | JSONPath expression |
| `jsonpath-as-json` | JSONPath result as JSON objects |
| `jsonpath-file` | JSONPath expression from a file |
| `custom-columns` | User-defined table columns |
| `custom-columns-file` | Custom columns from a file |

### Basic Formats

```shell
kubectl get pods                          # default table
kubectl get pods -o wide                  # extra columns
kubectl get pods -o json                  # JSON
kubectl get pods -o yaml                  # YAML
kubectl get pods -o name                  # pod/nginx, pod/redis
```

### `name` Format

Prints only the resource type and name — useful for scripting:

```shell
kubectl get pods -o name
# pod/nginx-76b4fb4578-sv42x
# pod/nginx-76b4fb4578-zx8p9

# Delete all pods selected by label
kubectl delete pods -l app=old-app -o name
```

### `wide` Format

Like the default table but with additional server-defined columns:

```shell
kubectl get pods -o wide
# NAME                    READY   STATUS    RESTARTS  AGE   IP           NODE       NOMINATED NODE
# nginx-76b4fb4578-sv42x  1/1     Running   0         5m    10.244.1.5   worker-1   <none>
```

## JSONPath

JSONPath extracts specific fields from the JSON representation of resources. Expressions are wrapped in `{` and `}`; text outside expressions is literal.

### Syntax

```shell
kubectl get pods -o jsonpath='{.items[0].metadata.name}'
kubectl get pods -o jsonpath='{.items[*].metadata.name}'
kubectl get pods -o jsonpath="{range .items[*]}{.metadata.name}{\"\t\"}{.status.startTime}{\"\n\"}{end}"
```

### Operators

| Operator | Description | Example |
|----------|-------------|---------|
| `$` | Root object (optional) | `{$.items}` or `{.items}` |
| `@` | Current object | `{@}` |
| `.` or `[]` | Child access | `{.kind}` or `{['kind']}` |
| `..` | Recursive descent | `{..name}` |
| `*` | Wildcard (all items) | `{.items[*].metadata.name}` |
| `[n]` | Array index | `{.items[0]}` |
| `[start:end]` | Slice | `{.items[0:2]}` |
| `[start:end:step]` | Slice with step | `{.items[0:4:2]}` |
| `[-1]` | Last element | `{.items[-1]}` |
| `[,]` | Union (multiple fields) | `{.items[*]['metadata.name','status.phase']}` |
| `?(@.field)` | Filter | `{.items[?(@.status.phase=="Running")]}` |
| `range ... end` | Iterate list | `{range .items[*]}{.metadata.name}{end}` |
| `'text'` | Quoted string | `{'\t'}` for tab, `{'\n'}` for newline |
| `\` | Escape character | `{.metadata.labels.kubernetes\.io/hostname}` |

### Key Details

- `$` is optional — expressions always start from root
- Double quotes are used inside expressions for string literals
- Regex is **not supported** — use `jq` for regex filtering
- Negative indices work (`[-1]`) but don't wrap around

### Practical Examples

```shell
# All pod names
kubectl get pods -o jsonpath='{.items[*].metadata.name}'

# First pod's name
kubectl get pods -o jsonpath='{.items[0].metadata.name}'

# Pod names and their phases
kubectl get pods -o jsonpath="{.items[*]['metadata.name', 'status.phase']}"

# Tab-separated name and start time (table-like)
kubectl get pods -o jsonpath='{range .items[*]}{.metadata.name}{"\t"}{.status.startTime}{"\n"}{end}'

# Escape dots in label keys
kubectl get pods -o jsonpath='{.items[0].metadata.labels.kubernetes\.io/hostname}'

# Filter: only Running pods
kubectl get pods -o jsonpath='{.items[?(@.status.phase=="Running")].metadata.name}'

# All container images in a pod
kubectl get pod my-pod -o jsonpath='{.spec.containers[*].image}'

# List pod names as a JSON array (structured output)
kubectl get pods -o jsonpath-as-json='{.items[*].metadata.name}'
```

### `jsonpath-as-json`

Outputs the JSONPath result as proper JSON objects/arrays instead of concatenated strings:

```shell
kubectl get pods -o jsonpath-as-json='{.items[*].metadata.name}'
# ["nginx-76b4fb4578-sv42x","nginx-76b4fb4578-zx8p9"]
```

### `jsonpath-file`

Reads the JSONPath template from a file — useful for complex expressions:

```shell
kubectl get pods -o jsonpath-file=template.txt
```

### Windows Quoting

On Windows, use double quotes for templates containing spaces, and escape literal double quotes:

```cmd
kubectl get pods -o jsonpath="{range .items[*]}{.metadata.name}{'\t'}{.status.startTime}{'\n'}{end}"
```

## Go Templates

Go templates use Go's [`text/template`](https://pkg.go.dev/text/template) package. Actions are wrapped in `{{ }}`, with `.` as the root object.

### Basic Syntax

```shell
kubectl get pods -o go-template='{{.metadata.name}}'
kubectl get pods -o go-template='{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}'
```

### Actions

| Action | Description |
|--------|-------------|
| `{{.field}}` | Field access |
| `{{range .items}}...{{end}}` | Iterate list |
| `{{if eq .x "y"}}...{{end}}` | Conditional |
| `{{printf "%s" .name}}` | Format string |
| `{{len .items}}` | Length of list |
| `{{index .items 0}}` | Index into list/map |
| `{{or .field "default"}}` | First non-empty value |
| `{{$var := .x}}` | Variable assignment |
| `{{- -}}` | Trim surrounding whitespace |

### Examples

```shell
# All pod names, one per line
kubectl get pods -o go-template='{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}'

# With index
kubectl get pods -o go-template='{{range $i, $p := .items}}{{$i}}: {{$p.metadata.name}}{{"\n"}}{{end}}'

# Formatted columns
kubectl get pods -A -o go-template='{{range .items}}{{printf "%-40s %20s\n" .metadata.name .metadata.namespace}}{{end}}'

# Conditional: only pods in default namespace
kubectl get pods -A -o go-template='{{range .items}}{{if eq .metadata.namespace "default"}}{{.metadata.name}}{{"\n"}}{{end}}{{end}}'

# Pod and container images
kubectl get pods -o go-template='{{range .items}}{{.metadata.name}}{{range .spec.containers}} {{.name}}:{{.image}}{{end}}{{"\n"}}{{end}}'

# Count pods
kubectl get pods -o go-template='{{printf "%d\n" (len .items)}}'

# Handle missing fields with or
kubectl get pods -o go-template='{{range .items}}{{.metadata.name}} {{or .spec.containers[0].resources.requests.cpu "unspecified"}}{{"\n"}}{{end}}'
```

### Comparison Operators

`eq`, `ne`, `lt`, `le`, `gt`, `ge` — for use in `{{if}}` conditions.

### `go-template-file`

For complex templates, store in a file:

```shell
kubectl get pods -o go-template-file=pod-template.gotemplate
```

## Custom Columns

Custom columns produce a clean table with user-defined columns using JSONPath expressions:

### Syntax

```shell
kubectl get pods -o custom-columns=HEADER:.jsonpath,HEADER2:.jsonpath2
```

### Examples

```shell
# Single column
kubectl get pods -o custom-columns="POD-NAME":.metadata.name

# Multiple columns
kubectl get pods -o custom-columns=NAME:.metadata.name,NAMESPACE:.metadata.namespace,IMAGES:.spec.containers[*].image

# CPU and memory requests
kubectl get pods -o custom-columns='POD:.metadata.name,CPU:.spec.containers[*].resources.requests.cpu,MEM:.spec.containers[*].resources.requests.memory'

# Pod name and volumes
kubectl get pods -o custom-columns='POD:.metadata.name,VOLUMES:.spec.volumes[*].name'
```

### `custom-columns-file`

Store column definitions in a file (space or tab separated):

```
NAME          RSRC
metadata.name metadata.resourceVersion
```

```shell
kubectl get pods -o custom-columns-file=columns.txt
```

### Discovering Fields with `explain`

Use `kubectl explain` to understand the YAML structure and find fields for your expressions:

```shell
kubectl explain pods.spec.containers.resources.requests
kubectl explain pods --recursive | grep -A5 resources
```

## `--sort-by`

Sorts output by a field expressed as a JSONPath (note: **no curly braces** for `--sort-by`):

```shell
kubectl get pods --sort-by=.metadata.name
kubectl get pods --sort-by=.status.containerStatuses[0].restartCount
kubectl get pods --sort-by=.metadata.creationTimestamp
kubectl get pods -o wide --sort-by=.spec.nodeName
kubectl get pods --sort-by=.spec.containers[0].resources.requests.cpu
```

The field must be a string or numeric type. Combine with any output format.

## `--server-print`

By default, kubectl uses **server-side printing** — the API server returns a table format with server-defined columns. This supports extension APIs and CRDs.

```shell
kubectl get pods --server-print=false   # disable server-side table printing
```

When `--server-print=false`, kubectl falls back to its built-in column definitions, which may differ from the server's.

## Combining with Selectors

Output formatting combines with label selectors and field selectors:

```shell
kubectl get pods -l app=nginx -o wide --sort-by=.metadata.creationTimestamp
kubectl get pods --field-selector=status.phase=Running -o custom-columns=NAME:.metadata.name,NODE:.spec.nodeName
kubectl get pods -A -o jsonpath='{range .items[*]}{.metadata.namespace}{"/"}{.metadata.name}{"\n"}{end}'
```

## Common Patterns

### Extract Pod IPs

```shell
kubectl get pods -o jsonpath='{range .items[*]}{.metadata.name}{"\t"}{.status.podIP}{"\n"}{end}'
```

### List All Container Images in a Namespace

```shell
kubectl get pods -o jsonpath='{range .items[*]}{range .spec.containers[*]}{.image}{"\n"}{end}{end}' | sort -u
```

### Show Node Capacity

```shell
kubectl get nodes -o custom-columns="NODE:.metadata.name,CPU:.status.capacity.cpu,MEM:.status.capacity.memory"
```

### Get Pod Resource Requests and Limits

```shell
kubectl get pods -o custom-columns='NAME:.metadata.name,CPU_REQ:.spec.containers[*].resources.requests.cpu,CPU_LIM:.spec.containers[*].resources.limits.cpu'
```

### Extract All Labels as Key=Value

```shell
kubectl get pods -o jsonpath='{range .items[*]}{.metadata.name}{"\t"}{.metadata.labels}{"\n"}{end}'
```

## JSONPath vs Go Template vs jq

| Feature | JSONPath | Go Template | `jq` (external) |
|---------|----------|-------------|-----------------|
| Field extraction | Yes | Yes | Yes |
| Filtering | `?(@.field=="value")` | `{{if eq .field "value"}}` | `select(.field == "value")` |
| Regex | No | No | Yes |
| Iteration | `range/end` | `range/end` | `| .items[]` |
| Formatting | Limited | Full `printf` | Full |
| Conditionals | Filter only | `if/else` | `select()` |
| External dependency | No | No | Yes (must install `jq`) |

When JSONPath is insufficient (e.g., regex, complex filtering), pipe to `jq`:

```shell
kubectl get pods -o json | jq -r '.items[] | select(.metadata.name | test("web-")).metadata.name'
```

## Edge Cases and Known Issues

- **JSONPath regex is not supported**: kubectl's JSONPath implementation doesn't support regex. Use `jq` instead.
- **`--sort-by` uses no braces**: `--sort-by=.metadata.name`, not `--sort-by={.metadata.name}`.
- **Escaping dots in keys**: Label keys with dots (e.g., `kubernetes.io/hostname`) must be escaped: `{.metadata.labels.kubernetes\.io/hostname}`.
- **`custom-columns` and `custom-columns-file`**: The file format uses whitespace between header and expression — don't use commas.
- **Server-side printing may change columns**: Different API server versions may define different default columns. `--server-print=false` forces client-side columns.
- **`-o name` output is `type/name`**: Not just the name. Use jsonpath if you need only the name.
- **Empty results**: JSONPath with no matches produces empty output (no error). `go-template` may produce an error if a field is missing and `--allow-missing-template-keys=false`.
- **Quoting on shells**: Single quotes preserve the template literally on bash/zsh. Use double quotes when the template contains single quotes or for Windows compatibility.

## References

- [JSONPath Support](https://kubernetes.io/docs/reference/kubectl/jsonpath/)
- [kubectl Output Options](https://kubernetes.io/docs/reference/kubectl/#output-options)
- [Go text/template](https://pkg.go.dev/text/template)

## Related Skills

- For resource types and field selector details, see [resources.md](resources.md).
- For the `get` command and its flags, see [cli.md](cli.md).
