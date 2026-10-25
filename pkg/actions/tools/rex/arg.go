package rex

import (
	"time"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/cache/key"
)

// ClientActionArgsOpts provides target filters for action argument completion
type ClientActionArgsOpts struct {
	Action string // action name
	Client string // client target
}

func (o ClientActionArgsOpts) cacheKey() key.Key {
	return key.String(o.Client, o.Action)
}

// ActionClientActionArgs completes argument names of a client action
//
//	direction
//	ratio
func ActionClientActionArgs(opts ClientActionArgsOpts) carapace.Action {
	args := []string{"actions", "--json"}
	if opts.Client != "" {
		args = append(args, "--client", opts.Client)
	}
	return carapace.ActionExecCommand("rex", args...)(func(output []byte) carapace.Action {
		actions, err := unmarshalActions(output)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}
		vals := make([]string, 0)
		for _, action := range actions {
			if action.Name != opts.Action {
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
	}).Cache(24*time.Hour, opts.cacheKey()).Tag("action args")
}
