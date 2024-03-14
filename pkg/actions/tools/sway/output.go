package sway

import (
	"encoding/json"
	"fmt"

	"github.com/carapace-sh/carapace"
)

type swayOutput struct {
	Name  string
	Make  string
	Modes []struct {
		Width   int
		Height  int
		Refresh int
	}
}

// ActionOutputs completes outputs
func ActionOutputs() carapace.Action {
	return carapace.ActionExecCommand("swaymsg", "-t", "get_outputs")(func(output []byte) carapace.Action {
		var outputs []swayOutput
		if err := json.Unmarshal(output, &outputs); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, o := range outputs {
			description := o.Make
			if len(o.Modes) > 0 {
				description += fmt.Sprintf(" %vx%v @ %v Hz", o.Modes[0].Width, o.Modes[0].Height, float64(o.Modes[0].Refresh)/1000)
			}
			vals = append(vals, o.Name, description)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
