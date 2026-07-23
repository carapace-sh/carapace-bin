package herdr

import (
	"encoding/json"

	"github.com/carapace-sh/carapace"
)

type workspace struct {
	WorkspaceID string `json:"workspace_id"`
	Label       string `json:"label"`
	Number      int    `json:"number"`
	TabCount    int    `json:"tab_count"`
	PaneCount   int    `json:"pane_count"`
	AgentStatus string `json:"agent_status"`
	Focused     bool   `json:"focused"`
}

type workspaceListResult struct {
	Workspaces []workspace `json:"workspaces"`
	Type       string      `json:"type"`
}

// ActionWorkspaces completes workspaces
//
//	ws-1 (Label)
//	ws-2 (Label (focused))
func ActionWorkspaces() carapace.Action {
	return carapace.ActionExecCommand("herdr", "workspace", "list")(func(output []byte) carapace.Action {
		wrapper := struct {
			Result workspaceListResult `json:"result"`
		}{}
		if err := json.Unmarshal(output, &wrapper); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, w := range wrapper.Result.Workspaces {
			desc := w.Label
			if w.Focused {
				desc += " (focused)"
			}
			vals = append(vals, w.WorkspaceID, desc)
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("workspaces")
}
