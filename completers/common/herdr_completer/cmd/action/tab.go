package action

import (
	"encoding/json"

	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

type tab struct {
	TabID       string `json:"tab_id"`
	Label       string `json:"label"`
	Number      int    `json:"number"`
	WorkspaceID string `json:"workspace_id"`
	PaneCount   int    `json:"pane_count"`
	AgentStatus string `json:"agent_status"`
	Focused     bool   `json:"focused"`
}

type tabListResult struct {
	Tabs []tab  `json:"tabs"`
	Type string `json:"type"`
}

func ActionTabs(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"tab", "list"}
		if flag := cmd.Flag("workspace"); flag != nil && flag.Changed {
			args = append(args, "--workspace", flag.Value.String())
		}
		return carapace.ActionExecCommand("herdr", args...)(func(output []byte) carapace.Action {
			wrapper := struct {
				Result tabListResult `json:"result"`
			}{}
			if err := json.Unmarshal(output, &wrapper); err != nil {
				return carapace.ActionMessage(err.Error())
			}

			vals := make([]string, 0)
			for _, t := range wrapper.Result.Tabs {
				desc := t.Label
				if t.Focused {
					desc += " (focused)"
				}
				vals = append(vals, t.TabID, desc)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	}).Tag("tabs")
}
