package bridge

import (
	"fmt"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/xdg"
)

// ActionFish bridges completions registered in fish shell
// (uses custom `config.fish` in “~/.config/carapace/bridge/fish`)
func ActionFish(command ...string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if len(command) == 0 {
			return carapace.ActionMessage("missing argument [ActionFish]")
		}

		configDir, err := xdg.UserConfigDir()
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		args := append(command[1:], c.Args...)
		args = append(args, c.CallbackValue)
		for index, arg := range args {
			arg = strings.Replace(arg, `\`, `\\`, -1)
			arg = strings.Replace(arg, `'`, `\'`, -1)
			arg = strings.Replace(arg, ` `, `\ `, -1)
			args[index] = arg
		}

		c.Setenv("XDG_CONFIG_HOME", fmt.Sprintf("%v/carapace/bridge", configDir)) // use custom config.fish for bridge
		return carapace.ActionExecCommand("fish", "--command", fmt.Sprintf(`complete '--do-complete=%v %v'`, command[0], strings.Join(args, " ")))(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			nospace := false

			vals := make([]string, 0)
			for _, line := range lines[:len(lines)-1] {
				splitted := strings.SplitN(line, "\t", 2)

				if len(splitted) > 1 {
					vals = append(vals, splitted...)
				} else {
					vals = append(vals, splitted[0], "")
				}
				if value := splitted[0]; !nospace && len(value) > 0 && strings.ContainsAny(value[len(value)-1:], `/=@:.,-`) {
					nospace = true
				}

			}
			a := carapace.ActionValuesDescribed(vals...)
			if nospace {
				a = a.NoSpace()
			}
			return a
		}).Invoke(c).ToA()
	}).Tag("fish bridge")
}
