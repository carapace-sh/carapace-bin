package rex

import (
	"time"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/cache/key"
)

// ActionClientActionArgs completes argument names of a client action
//
//	direction
//	ratio
func ActionClientActionArgs(name, client string) carapace.Action {
	args := []string{"actions", "--json"}
	if client != "" {
		args = append(args, "--client", client)
	}
	return carapace.ActionExecCommand("rex", args...)(func(output []byte) carapace.Action {
		actions, err := unmarshalActions(output)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}
		vals := make([]string, 0)
		for _, action := range actions {
			if action.Name != name {
				continue
			}
			for property, schema := range action.Args.Properties {
				description := schemaType(schema.Type)
				for _, required := range action.Args.Required {
					if required == property {
						description = "required"
					}
				}
				vals = append(vals, property, description)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Cache(24*time.Hour, key.String(client, name)).Tag("action args")
}
