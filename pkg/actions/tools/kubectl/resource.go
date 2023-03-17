package kubectl

import (
	"strings"

	"github.com/rsteube/carapace"
)

type ResourceOpts struct {
	Namespace string
	Types     string
}

// ActionResources completes resources
// TODO example
func ActionResources(opts ResourceOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("kubectl", "--namespace", opts.Namespace, "get", "-o", "go-template={{range .items}}{{.metadata.name}}\n{{.kind}}\n{{end}}", opts.Types)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(lines[:len(lines)-1]...)
		})
	})
}
