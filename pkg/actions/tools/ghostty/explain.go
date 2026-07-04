package ghostty

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionKeybindActions completes keybind actions
//
//	ignore
//	unbind
func ActionKeybindActions() carapace.Action {
	return carapace.ActionExecCommand("ghostty", "+list-actions")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[:len(lines)-1]...).
			Tag("keybind actions").
			UidF(Uid("keybind-action"))
	})
}

// ActionConfigKeys completes config keys
//
//	font-size
//	background
func ActionConfigKeys() carapace.Action {
	return carapace.ActionExecCommand("ghostty", "+show-config", "--default", "--docs")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		keys := make([]string, 0)
		for _, line := range lines {
			if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, " ") {
				continue
			}
			parts := strings.SplitN(line, " ", 2)
			keys = append(keys, parts[0])
		}
		return carapace.ActionValues(keys...).
			Tag("config keys").
			UidF(Uid("config-key"))
	})
}
