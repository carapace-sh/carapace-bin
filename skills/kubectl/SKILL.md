---
name: kubectl
description: >
  Use when working with the kubectl CLI — Kubernetes command-line tool for managing cluster
  resources. Covers the command structure, resource types and API groups, kubeconfig and
  contexts, declarative apply, workload management (deployments, rollout, scale, autoscale,
  expose), debugging (logs, exec, port-forward, debug, top), cluster management (cordon,
  drain, taint), output formats (jsonpath, go-template, custom-columns), patches, labels,
  annotations, kustomize, and the plugin system. Triggers on: "kubectl", "kubectl get",
  "kubectl apply", "kubectl create", "kubectl delete", "kubectl describe", "kubectl logs",
  "kubectl exec", "kubectl config", "kubectl rollout", "kubectl scale", "kubectl expose",
  "kubectl port-forward", "kubectl top", "kubectl drain", "kubectl cordon", "kubectl taint",
  "kubectl patch", "kubectl label", "kubectl annotate", "kubectl kustomize", "kubectl wait",
  "kubectl debug", "kubectl explain", "kubeconfig", "kubectl context", "kubectl plugin".
user-invocable: true
---

# kubectl — Kubernetes Command-Line Tool In-Depth Reference

Comprehensive reference for [kubectl](https://kubernetes.io/docs/reference/kubectl/), the CLI for communicating with the Kubernetes control plane. Covers command structure, resource types, kubeconfig, declarative management, workload operations, debugging, cluster administration, and output formatting.

## Data Flow

```
kubectl command
  → parse global flags (--context, --namespace, --kubeconfig, --server)
    → resolve kubeconfig (KUBECONFIG env, ~/.kube/config, in-cluster)
      → authenticate (client cert, token, in-cluster service account)
        → REST request to API server
          → response marshalling (table, json, yaml, jsonpath, go-template)
```

## Sub-Resources

Load the reference that matches your task. When in doubt, load multiple references.

| Keywords | Reference |
|----------|----------|
| concept, model, control plane, API server, pod, container, deployment, replicaSet, service, namespace, node, kubeconfig, context, cluster, declarative, imperative, reconciliation, controller, reconciliation loop, label, annotation, selector, service account, RBAC, workloads, control plane vs data plane | [references/concepts.md](references/concepts.md) |
| resource type, API group, core/v1, apps/v1, batch/v1, networking.k8s.io, rbac, short name, abbreviation, namespaced, cluster-scoped, resource specification, TYPE/NAME, TYPE NAME, api-resources, api-versions, field selector, label selector, resource version, kind, crd, custom resource | [references/resources.md](references/resources.md) |
| cli, command, subcommand, command groups, basic commands, intermediate commands, deploy commands, cluster management, troubleshooting, advanced commands, settings commands, global flags, --context, --namespace, --kubeconfig, --server, --as, --as-group, --token, --certificate-authority, verbosity, -v, completion, options, version | [references/cli.md](references/cli.md) |
| config, kubeconfig, context, cluster, user, current-context, config view, config use-context, config set-context, config set-cluster, config set-credentials, config get-contexts, config get-clusters, KUBECONFIG, kubeconfig file, merge kubeconfig, named context, default namespace, namespace override | [references/config.md](references/config.md) |
| output format, json, yaml, kyaml, name, wide, go-template, go-template-file, jsonpath, jsonpath-as-json, jsonpath-file, custom-columns, custom-columns-file, sort-by, field-selector, template, --output, -o, jsonpath syntax, go template syntax, custom columns syntax, table output, server-print | [references/output.md](references/output.md) |
| apply, declarative, server-side apply, client-side apply, last-applied-configuration, prune, dry-run, validate, field-manager, managed fields, conflict, force-conflicts, diff, replace, create, delete, cascade, graceful deletion, save-config, kustomize, kustomization, edit, patch, strategic merge patch, json patch, merge patch, patch type | [references/apply.md](references/apply.md) |
| workload, deployment, statefulset, daemonset, replicaset, replicationcontroller, rollout, rollout history, rollout undo, rollout pause, rollout resume, rollout restart, rollout status, scale, autoscale, expose, service, run, set, set image, set env, set resources, set selector, set serviceaccount, set subject, job, cronjob, rolling update, canary | [references/workloads.md](references/workloads.md) |
| debugging, logs, exec, attach, cp, port-forward, proxy, top, top node, top pod, describe, events, debug, ephemeral container, wait, condition, troubleshoot, metrics | [references/debugging.md](references/debugging.md) |
| cluster management, node, cordon, uncordon, drain, taint, cluster-info, cluster-info dump, certificate, certificate approve, certificate deny, CSR, top node, top pod, metrics-server, auth, auth can-i, auth reconcile, auth whoami, RBAC, role, clusterrole, rolebinding, clusterrolebinding | [references/cluster.md](references/cluster.md) |
| plugin, krew, kubectl plugin list, kubectl- prefix, plugin discovery, kubectl plugin, plugin manager, kubectl convert, kubectl alpha, custom plugin, shell plugin, plugin path | [references/plugins.md](references/plugins.md) |

## Quick Guide

- **What is kubectl and how does it connect to a cluster?** → [references/concepts.md](references/concepts.md)
- **What resource types exist and how do I specify them?** → [references/resources.md](references/resources.md)
- **What are the command groups and global flags?** → [references/cli.md](references/cli.md)
- **How do I configure kubeconfig and contexts?** → [references/config.md](references/config.md)
- **How do I format output with jsonpath or go-templates?** → [references/output.md](references/output.md)
- **How does `kubectl apply` work vs `kubectl create`?** → [references/apply.md](references/apply.md)
- **How do I manage deployments and rollouts?** → [references/workloads.md](references/workloads.md)
- **How do I debug a pod (logs, exec, port-forward)?** → [references/debugging.md](references/debugging.md)
- **How do I manage nodes and RBAC?** → [references/cluster.md](references/cluster.md)
- **How do I install and use kubectl plugins?** → [references/plugins.md](references/plugins.md)
- **How do I use `kubectl patch`?** → [references/apply.md](references/apply.md)
- **How do I use label and field selectors?** → [references/resources.md](references/resources.md) and [references/output.md](references/output.md)
- **How do I drain a node for maintenance?** → [references/cluster.md](references/cluster.md)
- **How do I wait for a condition?** → [references/debugging.md](references/debugging.md)

## Cross-Project References

- For **carapace** completion integration for kubectl (the `completers/common/kubectl_completer/` synthetic cobra tree, shared actions in `pkg/actions/tools/kubectl/`, macros like `tools.kubectl.Resources`), see the **carapace** skill → `references/action.md`.
- For **Kustomize** as a standalone tool (advanced template-free customization, components, generators, transformers), see the upstream [Kustomize docs](https://kubectl.docs.kubernetes.io/). This skill covers `kubectl kustomize` integration only.
- For **Helm** (package management for Kubernetes), use the Helm documentation — kubectl does not manage Helm releases directly, though `kubectl kustomize --enable-helm` can inflate Helm charts.
