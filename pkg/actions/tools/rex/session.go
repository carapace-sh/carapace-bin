package rex

import (
	"encoding/json"
	"time"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/cache/key"
)

// ActionSessions completes session targets
//
//	session:01a10cf3-2655-7e48-9b9c-43dd6612bef5 (Research)
//	session:01a10d97-414b-79ee-9988-4d96d804ef90 (SL Website)
func ActionSessions(server string) carapace.Action {
	args := []string{"ls", "--json"}
	if server != "" {
		args = append(args, "--server", server)
	}
	return carapace.ActionExecCommand("rex", args...)(func(output []byte) carapace.Action {
		var sessions struct {
			Sessions []struct {
				SessionId string `json:"session_id"`
				Label     string `json:"label"`
			}
		}
		if err := json.Unmarshal(output, &sessions); err != nil {
			return carapace.ActionMessage(err.Error())
		}
		vals := make([]string, 0, len(sessions.Sessions)*2)
		for _, session := range sessions.Sessions {
			if session.Label != "" {
				vals = append(vals, session.SessionId, session.Label)
			} else {
				vals = append(vals, session.SessionId)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Cache(time.Minute, key.String(server)).Tag("sessions")
}
