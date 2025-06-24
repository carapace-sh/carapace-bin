package action

import (
	"regexp"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/systemctl"
	"github.com/spf13/cobra"
)

func userFlag(cmd *cobra.Command) bool {
	if flag := cmd.Root().Flag("user"); flag != nil && flag.Changed {
		return true
	}
	return false
}

func ActionUnits(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return systemctl.ActionUnits(systemctl.UnitOpts{User: userFlag(cmd), Active: true, Inactive: true})
	})
}

func ActionServices(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return systemctl.ActionServices(userFlag(cmd))
	})
}

func ActionArchitectures() carapace.Action {
	return carapace.ActionExecCommand("systemd-analyze", "architectures")(func(output []byte) carapace.Action {
		re := regexp.MustCompile(`(?m)^(\S+)\s+(\S+)$`)
		matches := re.FindAllSubmatch(output, -1)

		var vals []string
		for _, match := range matches {
			vals = append(vals, string(match[1]), string(match[2]))
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

func ActionCapabilities() carapace.Action {
	return carapace.ActionExecCommand("systemd-analyze", "capability")(func(output []byte) carapace.Action {
		re := regexp.MustCompile(`(?m)^(cap_\S+)\s+(\d+)$`)
		matches := re.FindAllSubmatch(output, -1)

		var vals []string
		for _, match := range matches {
			vals = append(vals, string(match[1]), string(match[2]))
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

func ActionSyscallSets() carapace.Action {
	return carapace.ActionExecCommand("systemd-analyze", "syscall-filter")(func(output []byte) carapace.Action {
		re := regexp.MustCompile(`(?m)^@(\S+)\n\s+#\s+(.*)$`)
		matches := re.FindAllSubmatch(output, -1)

		vals := []string{}
		for _, match := range matches {
			group := "@" + string(match[1])
			description := string(match[2])
			vals = append(vals, group, description)
		}

		return carapace.ActionValuesDescribed(vals...)
	})
}

func ActionFilesystemSets() carapace.Action {
	return carapace.ActionExecCommand("systemd-analyze", "filesystems")(func(output []byte) carapace.Action {
		re := regexp.MustCompile(`(?m)^@(\S+)\n\s+#\s+(.*)$`)
		matches := re.FindAllSubmatch(output, -1)

		vals := []string{}
		for _, match := range matches {
			group := "@" + string(match[1])
			description := string(match[2])
			vals = append(vals, group, description)
		}

		return carapace.ActionValuesDescribed(vals...)
	})
}
