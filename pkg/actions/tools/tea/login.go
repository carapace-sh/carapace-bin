package tea

import (
	"fmt"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"gopkg.in/yaml.v3"
)

// ActionLogins completes logins
//
//	codeberg (user@codeberg.org)
//	gitea (user@gitea.com)
func ActionLogins() carapace.Action {
	return carapace.ActionExecCommand("tea", "logins", "list", "--output", "yaml")(func(output []byte) carapace.Action {
		var logins []struct {
			Name    string `yaml:"Name"`
			SSHHost string `yaml:"SSHHost"`
			User    string `yaml:"User"`
			Default string `yaml:"Default"`
		}

		if err := yaml.Unmarshal(output, &logins); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, login := range logins {
			s := style.Default
			if login.Default == "true" {
				s = style.Blue
			}
			vals = append(vals, login.Name, fmt.Sprintf("%v@%v", login.User, login.SSHHost), s)
		}
		return carapace.ActionStyledValuesDescribed(vals...)

	})
}
