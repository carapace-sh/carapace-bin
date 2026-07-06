package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

// ActionMachines completes available machine types by querying the qemu-system binary.
func ActionMachines(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		binaryName := cmd.Name()
		return carapace.ActionExecCommand(binaryName, "-machine", "help")(func(output []byte) carapace.Action {
			lines := string(output)
			values := make([]string, 0)
			for _, line := range splitLines(lines) {
				if name, desc, ok := parseMachineHelpLine(line); ok {
					values = append(values, name, desc)
				}
			}
			if len(values) == 0 {
				return carapace.ActionValues()
			}
			return carapace.ActionValuesDescribed(values...).Tag("machine types")
		})
	})
}

// ActionCpuModels completes available CPU models by querying the qemu-system binary.
func ActionCpuModels(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		binaryName := cmd.Name()
		return carapace.ActionExecCommand(binaryName, "-cpu", "help")(func(output []byte) carapace.Action {
			lines := string(output)
			values := make([]string, 0)
			for _, line := range splitLines(lines) {
				if name, desc, ok := parseCpuHelpLine(line); ok {
					values = append(values, name, desc)
				}
			}
			if len(values) == 0 {
				return carapace.ActionValues()
			}
			return carapace.ActionValuesDescribed(values...).Tag("cpu models")
		})
	})
}

func splitLines(s string) []string {
	var lines []string
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			lines = append(lines, s[start:i])
			start = i + 1
		}
	}
	if start < len(s) {
		lines = append(lines, s[start:])
	}
	return lines
}

// parseMachineHelpLine parses a line like "pc-i440fx-9.2        Standard PC (i440FX + PIIX, 1996)"
func parseMachineHelpLine(line string) (name, desc string, ok bool) {
	trimmed := trimSpaces(line)
	if trimmed == "" || trimmed == "Supported machines are:" {
		return "", "", false
	}
	for i := 0; i < len(trimmed)-1; i++ {
		if i > 0 && trimmed[i] == ' ' && trimmed[i+1] == ' ' {
			name = trimSpaces(trimmed[:i])
			desc = trimSpaces(trimmed[i:])
			if name != "" {
				return name, desc, true
			}
		}
	}
	if trimmed != "" {
		return trimmed, "", true
	}
	return "", "", false
}

// parseCpuHelpLine parses a line like "  486                   (alias configured by machine type)"
func parseCpuHelpLine(line string) (name, desc string, ok bool) {
	trimmed := trimSpaces(line)
	if trimmed == "" || trimmed == "Available CPUs:" {
		return "", "", false
	}
	for i := 0; i < len(trimmed)-1; i++ {
		if i > 0 && trimmed[i] == ' ' && trimmed[i+1] == ' ' {
			name = trimSpaces(trimmed[:i])
			desc = trimSpaces(trimmed[i:])
			if name != "" {
				return name, desc, true
			}
		}
	}
	if trimmed != "" {
		return trimmed, "", true
	}
	return "", "", false
}

func trimSpaces(s string) string {
	start := 0
	end := len(s)
	for start < end && s[start] == ' ' {
		start++
	}
	for end > start && s[end-1] == ' ' {
		end--
	}
	return s[start:end]
}
