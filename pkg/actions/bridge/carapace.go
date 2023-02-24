package bridge

import (
	"fmt"
	"os"

	"github.com/rsteube/carapace"
	spec "github.com/rsteube/carapace-spec"
	"github.com/rsteube/carapace/pkg/xdg"
)

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
		if a, err := actionSpec(completer); err == nil {
			return a
		} else if !os.IsNotExist(err) {
			return carapace.ActionMessage(err.Error())
		}
		return actionEmbedded(completer)
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
