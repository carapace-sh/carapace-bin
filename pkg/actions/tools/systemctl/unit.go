package systemctl

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

type UnitOpts struct {
	User     bool
	Active   bool
	Inactive bool
}

func (o UnitOpts) Default() UnitOpts {
	o.Active = true
	o.Inactive = true
	return o
}

// ActionUnits completes units
//
//	NetworkManager.service (Network Manager)
//	apparmor.service (Load AppArmor profiles)
func ActionUnits(opts UnitOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"list-units", "--all"}
		if opts.User {
			args = append(args, "--user")
		}
		return carapace.ActionExecCommand("systemctl", args...)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")

			vals := make([]string, 0)
			for _, line := range lines[1:] {
				if line == "" {
					break
				}
				fields := strings.Fields(line)
				switch state := fields[2]; {
				case opts.Active && state == "active":
					vals = append(vals, fields[0], strings.Join(fields[4:], " "), style.Carapace.KeywordPositive)
				case opts.Inactive && state == "inactive":
					vals = append(vals, fields[0], strings.Join(fields[4:], " "), style.Carapace.KeywordNegative)
				}
			}
			return carapace.ActionStyledValuesDescribed(vals...)
		})
	})
}
