package tldr

import (
	"encoding/json"
	"strings"

	"github.com/carapace-sh/carapace"
)

func ActionCommands() carapace.Action {
	return carapace.ActionExecCommand("tldr", "--help")(func(output []byte) carapace.Action {
		switch {
		case strings.Contains(string(output), "Python"):
			return actionCommandsPython()
		case strings.Contains(string(output), "tealdeer"):
			return actionCommandsTealdeer()
		default:
			return carapace.ActionValues()
		}
	})
}

func actionCommandsPython() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("tldr", "--list")(func(output []byte) carapace.Action {
			var commands []string
			fixedOutput := strings.Replace(string(output), `'`, `"`, -1)
			if err := json.Unmarshal([]byte(fixedOutput), &commands); err != nil {
				return carapace.ActionMessage(err.Error())
			} else {
				if len(commands) == 0 {
					return carapace.ActionMessage("execute `tldr -u` first")
				} else {
					return carapace.ActionValues(commands...)
				}
			}
		})
	})
}

func actionCommandsTealdeer() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("tldr", "--list")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValues(lines...)
		})
	})
}
