package kubectl

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionDryRunModes completes dry run modes
//
//	none
//	server
func ActionDryRunModes() carapace.Action {
	return carapace.ActionValues("client", "none", "server")
}

// ActionOutputFormats completes output formats
//
//	json
//	yaml
func ActionOutputFormats() carapace.Action {
	return carapace.ActionValues(
		"go-template",
		"go-template-file",
		"json",
		"jsonpath",
		"jsonpath-as-json",
		"jsonpath-file",
		"kyaml",
		"name",
		"template",
		"templatefile",
		"yaml",
	)
}

// ActionResourceVerbs completes resource verbs
//
//	get
//	list
func ActionResourceVerbs() carapace.Action {
	return carapace.ActionValues("create", "delete", "deletecollection", "get", "list", "patch", "update", "watch")
}

// ActionNamespaceServiceAccounts completes namespaces and serviceaccounts separately
//
//	default (Namespace)
//	kube-public (Namespace)
func ActionNamespaceServiceAccounts() carapace.Action {
	return carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return ActionResources(ResourceOpts{Namespace: "", Types: "namespaces"}).Invoke(c).Suffix(":").ToA()
		case 1:
			return ActionResources(ResourceOpts{Namespace: c.Parts[0], Types: "serviceaccounts"})
		default:
			return carapace.ActionValues()
		}
	})
}

type ContainerOpts struct {
	Namespace string
	Resource  string
}

// ActionContainers completes containers
//
//	my-container
//	sidecar
func ActionContainers(opts ContainerOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("kubectl", "--namespace", opts.Namespace, "get", "-o", "go-template={{range .spec.containers}}{{.name}}\n{{end}}", opts.Resource)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValues(lines[:len(lines)-1]...)
		})
	})
}

type LabelOpts struct {
	Namespace string
	Resource  string
}

// ActionLabels completes labels
//
//	app=myapp
//	version=v1
func ActionLabels(opts LabelOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("kubectl", "--namespace", opts.Namespace, "get", "-o", "go-template={{range $key, $value := .metadata.labels}}{{$key}}\n{{$value}}\n{{end}}", opts.Resource)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(lines[:len(lines)-1]...)
		})
	})
}
