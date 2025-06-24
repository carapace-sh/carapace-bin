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

type UnitFileOpts struct {
	User bool

	// type
	Automount bool
	Mount     bool
	Path      bool
	Scope     bool
	Service   bool
	Slice     bool
	Socket    bool
	Target    bool
	Timer     bool

	// state
	Alias          bool
	Disabled       bool
	Enabled        bool
	EnabledRuntime bool
	Generated      bool
	Indirect       bool
	Static         bool
	Transient      bool
}

func (o UnitFileOpts) Default() UnitFileOpts {
	o.Automount = true
	o.Mount = true
	o.Path = true
	o.Scope = true
	o.Service = true
	o.Slice = true
	o.Socket = true
	o.Target = true
	o.Timer = true

	o.Alias = true
	o.Disabled = true
	o.Enabled = true
	o.EnabledRuntime = true
	o.Generated = true
	o.Indirect = true
	o.Static = true
	o.Transient = true
	return o
}

func (o UnitFileOpts) types() string {
	types := make([]string, 0)
	if o.Automount {
		types = append(types, "automount")
	}
	if o.Mount {
		types = append(types, "mount")
	}
	if o.Path {
		types = append(types, "path")
	}
	if o.Scope {
		types = append(types, "scope")
	}
	if o.Service {
		types = append(types, "service")
	}
	if o.Slice {
		types = append(types, "slice")
	}
	if o.Socket {
		types = append(types, "socket")
	}
	if o.Target {
		types = append(types, "target")
	}
	if o.Timer {
		types = append(types, "timer")
	}
	return strings.Join(types, ",")
}

func (o UnitFileOpts) states() string {
	states := make([]string, 0)
	if o.Alias {
		states = append(states, "alias")
	}
	if o.Disabled {
		states = append(states, "disabled")
	}
	if o.Enabled {
		states = append(states, "enabled")
	}
	if o.EnabledRuntime {
		states = append(states, "enabled-runtime")
	}
	if o.Generated {
		states = append(states, "generated")
	}
	if o.Indirect {
		states = append(states, "indirect")
	}
	if o.Static {
		states = append(states, "static")
	}
	if o.Transient {
		states = append(states, "transient")
	}
	return strings.Join(states, ",")
}

// ActionUnitFiles completes unit files
//
//	NetworkManager.service
//	apparmor.service
func ActionUnitFiles(opts UnitFileOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"list-unit-files", "--type", opts.types(), "--state", opts.states()}
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
				switch state := fields[1]; state {
				case "enabled", "alias": // TODO alias correct?
					vals = append(vals, fields[0], style.Carapace.KeywordPositive)
				case "disabled":
					vals = append(vals, fields[0], style.Carapace.KeywordNegative)
				default:
					vals = append(vals, fields[0], style.Default)
				}
			}
			return carapace.ActionStyledValues(vals...)
		})
	})
}
