package kubectl

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

type AnnotationOpts struct {
	Namespace string
	Resource  string
}

// ActionAnnotations completes annotations
//
//	key (value)
//	another (value2)
func ActionAnnotations(opts AnnotationOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("kubectl", "--namespace", opts.Namespace, "get", "-o", "go-template={{range $key, $value := .metadata.annotations}}{{$key}}\n{{$value}}\n{{end}}", opts.Resource)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(lines[:len(lines)-1]...)
		})
	})
}
