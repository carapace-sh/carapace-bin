package but

import (
	"encoding/json"
	"os/exec"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

type AliasOpts struct {
	User     bool `yaml:"user"`
	Default_ bool `yaml:"default"`
}

func (o AliasOpts) Default() AliasOpts {
	o.User = true
	o.Default_ = true
	return o
}

type butAliases struct {
	User []struct {
		Name  string
		Value string
		Scope string
	}
	Default map[string]string
}

func Aliases() (*butAliases, error) {
	output, err := exec.Command("but", "alias", "list", "--json").Output() // TODO support context?
	if err != nil {
		return nil, err
	}

	var aliases butAliases
	if err := json.Unmarshal(output, &aliases); err != nil {
		return nil, err
	}
	return &aliases, nil
}

// ActionAliases completes aliases
//
//	st (status)
//	stf (status --files)
func ActionAliases(opts AliasOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		aliases, err := Aliases()
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		if opts.User {
			for _, alias := range aliases.User {
				s := style.Blue
				if alias.Scope == "global" {
					s = style.Yellow
				}
				vals = append(vals, alias.Name, alias.Value, s)
			}
		}
		if opts.Default_ {
			for name, value := range aliases.Default {
				vals = append(vals, name, value, style.Default)
			}
		}
		return carapace.ActionStyledValuesDescribed(vals...)
	}).Tag("aliases")
}
