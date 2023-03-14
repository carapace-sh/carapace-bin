package bridge

import (
	"fmt"
	"os"

	"github.com/rsteube/carapace"
	spec "github.com/rsteube/carapace-spec"
	"github.com/rsteube/carapace/pkg/xdg"
)

// ActionCarapace bridges https://github.com/rsteube/carapace
func ActionCarapace(command ...string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if len(command) == 0 {
			return carapace.ActionMessage("missing argument [ActionCarapace]")
		}

		args := []string{"_carapace", "export", ""}
		args = append(args, command[1:]...)
		args = append(args, c.Args...)
		args = append(args, c.CallbackValue)
		return carapace.ActionExecCommand(command[0], args...)(func(output []byte) carapace.Action {
			if string(output) == "" {
				return carapace.ActionValues()
			}
			return carapace.ActionImport(output)
		})
	})
}

// ActionCarapaceBin bridges completions registered in carapace-bin
func ActionCarapaceBin(command ...string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if len(command) == 0 {
			return carapace.ActionMessage("missing argument [ActionCarapaceBin]")
		}

		c.Args = append(command[1:], c.Args...)
		if a, err := actionSpec(command[0]); err == nil {
			return a.Invoke(c).ToA()
		} else if !os.IsNotExist(err) {
			return carapace.ActionMessage(err.Error())
		}
		return actionEmbedded(command[0]).Invoke(c).ToA()
	})
}

func actionEmbedded(completer string) carapace.Action {
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

func actionSpec(completer string) (carapace.Action, error) {
	configDir, err := xdg.UserConfigDir()
	if err != nil {
		return carapace.ActionValues(), err
	}
	path := fmt.Sprintf("%v/carapace/specs/%v.yaml", configDir, completer)
	if _, err := os.Stat(path); err != nil {
		return carapace.ActionValues(), err
	}
	return spec.ActionSpec(path), nil
}
