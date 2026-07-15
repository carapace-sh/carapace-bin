package action

import (
	"encoding/json"

	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
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

func ActionPanes(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"pane", "list"}
		if flag := cmd.Flag("workspace"); flag != nil && flag.Changed {
			args = append(args, "--workspace", flag.Value.String())
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
