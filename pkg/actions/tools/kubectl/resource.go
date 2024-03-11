package kubectl

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

type ResourceOpts struct {
	Context   string
	Namespace string
	Types     string
}

// ActionResources completes resources
// TODO example
func ActionResources(opts ResourceOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("kubectl", "--context", opts.Context, "--namespace", opts.Namespace, "get", "-o", "go-template={{range .items}}{{.metadata.name}}\n{{.kind}}\n{{end}}", opts.Types)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(lines[:len(lines)-1]...)
		})
	})
}
