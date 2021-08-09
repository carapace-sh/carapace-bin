package git

import (
	"strings"

	"github.com/rsteube/carapace"
)

func ActionConfigs() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		// TODO support different git folder
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

func ActionColors() carapace.Action {
	return carapace.ActionValues("normal", "black", "red", "green", "yellow", "blue", "magenta", "cyan", "white")
}

func ActionTextAttributes() carapace.Action {
	return carapace.ActionValues("bold", "dim", "ul", "blink", "reverse")
}

func ActionColorConfigs() carapace.Action {
	return carapace.ActionMultiParts(" ", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return ActionColors()
		case 1:
			return ActionColors()
		case 2:
			return ActionTextAttributes()
		default:
			return carapace.ActionValues()
		}
	})
}

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

func ActionConfigTypeOptions(t string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		switch t {
		case "bool":
			return carapace.ActionValues("true", "false")
		case "bool-or-int":
			integers := carapace.ActionValues("1", "2", "3", "4", "5", "6", "7", "8", "9", "0").Invoke(c).Prefix(c.CallbackValue)
			return integers.Merge(carapace.ActionValues("true", "false").Invoke(c)).ToA()
		case "int":
			return carapace.ActionValues("1", "2", "3", "4", "5", "6", "7", "8", "9", "0").Invoke(c).Prefix(c.CallbackValue).ToA()
		case "path":
			return carapace.ActionFiles()
		case "color":
			return ActionColorConfigs()
		default:
			return carapace.ActionValues()
		}
	})
}
