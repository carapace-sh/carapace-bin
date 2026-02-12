package atuin

import (
	"maps"
	"slices"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionKeys completes keys
//
//	one
//	default.two
func ActionKeys(namespace string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"kv", "list"}
		switch namespace {
		case "":
			args = append(args, "--all-namespaces")
		default:
			args = append(args, "--namespace", namespace)
		}
		return carapace.ActionExecCommand("atuin", args...)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValues(lines[len(lines)-1])
		})
	}).Tag("keys")
}

// ActionNamespaces completes namespaces
//
//	default
//	custom
func ActionNamespaces() carapace.Action {
	return carapace.ActionExecCommand("atuin", "kv", "list", "--all-namespaces")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		unique := make(map[string]bool)
		for _, line := range lines {
			if namespace, _, ok := strings.Cut(line, "."); ok {
				unique[namespace] = true
			}
		}
		return carapace.ActionValues(slices.Collect(maps.Keys(unique))...)
	}).Tag("namespaces")
}
