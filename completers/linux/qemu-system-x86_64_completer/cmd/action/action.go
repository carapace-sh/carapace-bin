package action

import (
	"os/exec"

	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

// ActionMachines completes available machine types by querying the qemu-system binary.
func ActionMachines(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		binaryName := cmd.Name()
		return carapace.ActionExecCommandE(binaryName, "-machine", "help")(func(output []byte, err error) carapace.Action {
			if err != nil {
				if exitErr, ok := err.(*exec.ExitError); !ok || exitErr.ExitCode() != 1 {
					return carapace.ActionMessage(err.Error())
				}
			}
			return parseHelpOutput(output, "Supported machines are:", "machine types")
		})
	})
}

// ActionCpuModels completes available CPU models by querying the qemu-system binary.
func ActionCpuModels(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		binaryName := cmd.Name()
		return carapace.ActionExecCommandE(binaryName, "-cpu", "help")(func(output []byte, err error) carapace.Action {
			if err != nil {
				if exitErr, ok := err.(*exec.ExitError); !ok || exitErr.ExitCode() != 1 {
					return carapace.ActionMessage(err.Error())
				}
			}
			return parseHelpOutput(output, "Available CPUs:", "cpu models")
		})
	})
}

func parseHelpOutput(output []byte, skipLine string, tag string) carapace.Action {
	lines := string(output)
	values := make([]string, 0)
	for _, line := range splitLines(lines) {
		if name, desc, ok := parseHelpLine(line, skipLine); ok {
			values = append(values, name, desc)
		}
	}
	if len(values) == 0 {
		return carapace.ActionValues()
	}
	return carapace.ActionValuesDescribed(values...).Tag(tag)
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

// parseHelpLine parses a line like "pc-i440fx-9.2        Standard PC (i440FX + PIIX, 1996)"
// or "  486                   (alias configured by machine type)"
func parseHelpLine(line string, skipLine string) (name, desc string, ok bool) {
	trimmed := trimSpaces(line)
	if trimmed == "" || trimmed == skipLine {
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
