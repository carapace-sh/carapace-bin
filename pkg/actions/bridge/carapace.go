package bridge

import "github.com/rsteube/carapace"

// ActionCarapace bridges rsteube/carapace completion
func ActionCarapace(cmd string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"_carapace", "export", ""}
		args = append(args, c.Args...)
		args = append(args, c.CallbackValue)
		return carapace.ActionExecCommand(cmd, args...)(func(output []byte) carapace.Action {
			if string(output) == "" {
				return carapace.ActionValues()
			}
			return carapace.ActionImport(output)
		})
	})
}

// ActionCarapaceBin bridges a completer provided by carapace-bin
func ActionCarapaceBin(completer string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{completer, "export", ""}
		args = append(args, c.Args...)
		args = append(args, c.CallbackValue)
		return carapace.ActionExecCommand("carapace", args...)(func(output []byte) carapace.Action {
			if string(output) == "" {
				return carapace.ActionValues()
			}
			return carapace.ActionImport(output)
		})
	})
}
