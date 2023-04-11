package git

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

// ActionConfigs completes configs
func ActionConfigs() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("git", "help", "--config")(func(output []byte) carapace.Action {
			vals := make([]string, 0)
			for _, line := range strings.Split(string(output), "\n") {
				if line != "" && !strings.ContainsAny(line, "<>*") && !strings.Contains(line, "git help config") {
					vals = append(vals, line)
				}
			}
			return carapace.ActionValues(vals...).Invoke(c).ToMultiPartsA(".")
		})
	})
}

// ActionColors completes colors
func ActionColors() carapace.Action {
	return carapace.ActionStyledValues(
		"normal", style.Default,
		"black", style.Black,
		"red", style.Red,
		"green", style.Green,
		"yellow", style.Yellow,
		"blue", style.Blue,
		"magenta", style.Magenta,
		"cyan", style.Cyan,
		"white", style.White,
	)
}

// ActionTextAttributes completes test attributes
func ActionTextAttributes() carapace.Action {
	return carapace.ActionStyledValues(
		"bold", style.Bold,
		"dim", style.Dim,
		"ul", style.Underlined,
		"blink", style.Blink,
		"reverse", style.Default,
	)
}

// ActionColorConfigs completes color config
func ActionColorConfigs() carapace.Action {
	return carapace.ActionMultiParts(" ", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return ActionColors().NoSpace()
		case 1:
			return ActionColors().NoSpace()
		case 2:
			return ActionTextAttributes()
		default:
			return carapace.ActionValues()
		}
	})
}

// ActionConfigTypes completes config types
func ActionConfigTypes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"bool", "canonicalize values as either \"true\" or \"false\"",
		"int", "canonicalize values as simple decimal numbers",
		"bool-or-int", "canonicalize according to either bool or int",
		"path", "canonicalize by adding a leading ~ to the value of $HOME",
		"expiry-date", "canonicalize by converting from a fixed or relative date-string to a timestamp",
		"color", "When getting a value, canonicalize by converting to an ANSI color escape sequence",
	)
}

// ActionConfigTypeOptions completes options for a config type
func ActionConfigTypeOptions(t string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		switch t {
		case "bool":
			return carapace.ActionValues("true", "false")
		case "bool-or-int":
			integers := carapace.ActionValues("1", "2", "3", "4", "5", "6", "7", "8", "9", "0").Invoke(c).Prefix(c.Value)
			return integers.Merge(carapace.ActionValues("true", "false").Invoke(c)).ToA()
		case "int":
			return carapace.ActionValues("1", "2", "3", "4", "5", "6", "7", "8", "9", "0").Invoke(c).Prefix(c.Value).ToA()
		case "path":
			return carapace.ActionFiles()
		case "color":
			return ActionColorConfigs()
		default:
			return carapace.ActionValues()
		}
	})
}
