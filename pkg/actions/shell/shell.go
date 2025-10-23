package shell

import (
	"os"
	"slices"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/carapace-sh/carapace/pkg/style"
)

// TODO support more shells

// ActionBuiltins completes shell builtins
func ActionBuiltins() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		action := carapace.ActionValues(strings.Fields(os.Getenv("CARAPACE_SHELL_BUILTINS"))...)
		switch os.Getenv("CARAPACE_SHELL") {
		case "bash":
			action = action.Filter(".", ":", "[")
		case "fish":
			action = action.Filter("-", ".", ":", "[", "_")
		case "zsh":
			action = action.Filter("-", ".", ":", "[")
		}
		return action
	}).Tag("shell builtins")
}

// ActionFunctions completes shell functions
func ActionFunctions(hidden bool) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		vals := strings.Fields(os.Getenv("CARAPACE_SHELL_FUNCTIONS"))
		if !hidden && !strings.HasPrefix(c.Value, "_") {
			vals = slices.DeleteFunc(vals, func(s string) bool { return strings.HasPrefix(s, "_") })
		}
		return carapace.ActionValues(vals...).StyleF(func(s string, sc style.Context) string {
			if strings.HasPrefix(s, "_") {
				return style.Dim
			}
			return style.Default
		})
	}).Tag("shell functions")
}

// ActionExecutables completes builtins, functions, and commands
func ActionExecutables() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		switch len(c.Args) {
		case 0:
			return carapace.Batch(
				// TODO which order?
				ActionBuiltins(),
				ActionFunctions(false),
				carapace.ActionExecutables(),
				carapace.ActionFiles(),
			).ToA()
		default:
			return bridge.ActionCarapaceBin()
		}
	})
}
