package kubectl

import (
	"strings"

	"github.com/rsteube/carapace"
)

// ActionApiGroups completes api groups
//
//	authorization.k8s.io
//	autoscaling
func ActionApiGroups() carapace.Action {
	return carapace.ActionValues(
		"admissionregistration.k8s.io",
		"apiextensions.k8s.io",
		"apiregistration.k8s.io",
		"apps",
		"authentication.k8s.io",
		"authorization.k8s.io",
		"autoscaling",
		"batch",
		"certificates.k8s.io",
		"coordination.k8s.io",
		"core",
		"discovery.k8s.io",
		"events.k8s.io",
		"extensions",
		"flowcontrol.apiserver.k8s.io",
		"networking.k8s.io",
		"node.k8s.io",
		"policy",
		"rbac.authorization.k8s.io",
		"scheduling.k8s.io",
		"settings.k8s.io",
		"storage.k8s.io",
	)
}

// ActionApiResourceResources completes api resources and resources separately
//
//	apiservices/v1.admissionregistration.k8s.io
//	endpoints/kubernetes
func ActionApiResourceResources() carapace.Action {
	return carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return ActionApiResources().Invoke(c).Suffix("/").ToA()
		case 1:
			return ActionResources(ResourceOpts{Namespace: "", Types: c.Parts[0]})
		default:
			return carapace.ActionValues()
		}
	})
}

// ActionApiResources completes api resources
//
//	apiservices
//	endpoints
func ActionApiResources() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("kubectl", "api-resources", "--output=name", "--cached")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			for index, line := range lines {
				lines[index] = strings.SplitN(line, ".", 2)[0]
			}
			return carapace.ActionValues(lines[:len(lines)-1]...)
		})
	})
}
