package herdr

import (
	"encoding/json"

	"github.com/carapace-sh/carapace"
)

type agent struct {
	Agent     string `json:"agent"`
	AgentID   string `json:"agent_id"`
	PaneID    string `json:"pane_id"`
	Status    string `json:"agent_status"`
	Cwd       string `json:"cwd"`
	Focused   bool   `json:"focused"`
	Workspace string `json:"workspace_id"`
}

type agentListResult struct {
	Agents []agent `json:"agents"`
	Type   string  `json:"type"`
}

// ActionAgents completes agents
//
//	pane-1 (agent-name)
//	agent-name (pane-1)
func ActionAgents() carapace.Action {
	return carapace.ActionExecCommand("herdr", "agent", "list")(func(output []byte) carapace.Action {
		var result agentListResult
		wrapper := struct {
			Result agentListResult `json:"result"`
		}{}
		if err := json.Unmarshal(output, &wrapper); err != nil {
			return carapace.ActionMessage(err.Error())
		}
		result = wrapper.Result

		vals := make([]string, 0)
		for _, a := range result.Agents {
			if a.Agent != "" {
				vals = append(vals, a.Agent, a.PaneID)
			}
			vals = append(vals, a.PaneID, a.Agent)
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("agents")
}
