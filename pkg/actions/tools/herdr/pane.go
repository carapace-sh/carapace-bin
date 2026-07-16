package herdr

import (
	"encoding/json"

	"github.com/carapace-sh/carapace"
)

type pane struct {
	PaneID      string `json:"pane_id"`
	TabID       string `json:"tab_id"`
	WorkspaceID string `json:"workspace_id"`
	Agent       string `json:"agent"`
	AgentStatus string `json:"agent_status"`
	Cwd         string `json:"cwd"`
	Focused     bool   `json:"focused"`
}

type paneListResult struct {
	Panes []pane `json:"panes"`
	Type  string `json:"type"`
}

// PaneOpts specifies options for completing panes
type PaneOpts struct {
	// Workspace filters panes by workspace id
	Workspace string
}

// ActionPanes completes panes
//
//	pane-1 (agent-name)
//	pane-2 (running)
func ActionPanes(opts PaneOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"pane", "list"}
		if opts.Workspace != "" {
			args = append(args, "--workspace", opts.Workspace)
		}
		return carapace.ActionExecCommand("herdr", args...)(func(output []byte) carapace.Action {
			wrapper := struct {
				Result paneListResult `json:"result"`
			}{}
			if err := json.Unmarshal(output, &wrapper); err != nil {
				return carapace.ActionMessage(err.Error())
			}

			vals := make([]string, 0)
			for _, p := range wrapper.Result.Panes {
				desc := p.Agent
				if desc == "" {
					desc = p.AgentStatus
				}
				if p.Focused {
					desc += " (focused)"
				}
				vals = append(vals, p.PaneID, desc)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	}).Tag("panes")
}
