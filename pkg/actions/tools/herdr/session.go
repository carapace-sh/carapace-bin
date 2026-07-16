package herdr

import (
	"encoding/json"

	"github.com/carapace-sh/carapace"
)

type session struct {
	Name       string `json:"name"`
	Running    bool   `json:"running"`
	Default    bool   `json:"default"`
	SocketPath string `json:"socket_path"`
}

// ActionSessions completes sessions
//
//	default (running (default))
//	other (stopped)
func ActionSessions() carapace.Action {
	return carapace.ActionExecCommand("herdr", "session", "list", "--json")(func(output []byte) carapace.Action {
		wrapper := struct {
			Sessions []session `json:"sessions"`
		}{}
		if err := json.Unmarshal(output, &wrapper); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, s := range wrapper.Sessions {
			desc := "stopped"
			if s.Running {
				desc = "running"
			}
			if s.Default {
				desc += " (default)"
			}
			vals = append(vals, s.Name, desc)
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("sessions")
}
