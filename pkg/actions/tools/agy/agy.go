package agy

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionAgents completes agents
//
//	my-agent
//	default
func ActionAgents() carapace.Action {
	return carapace.ActionExecCommand("agy", "agents")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines {
			if line != "" {
				parts := strings.Fields(line)
				if len(parts) > 0 {
					vals = append(vals, parts[0])
				}
			}
		}
		return carapace.ActionValues(vals...)
	}).Tag("agents")
}

// ActionModels completes models
//
//	gemini-3.5-flash-medium	Gemini 3.5 Flash (Medium)
//	claude-sonnet-4-6	Claude Sonnet 4.6 (Thinking)
func ActionModels() carapace.Action {
	return carapace.ActionExecCommand("agy", "models")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines {
			if line == "" {
				continue
			}
			parts := strings.SplitN(line, "\t", 2)
			if len(parts) == 2 {
				vals = append(vals, parts[0], parts[1])
			} else {
				vals = append(vals, parts[0], "")
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("models")
}

// ActionPlugins completes installed plugins
//
//	my-plugin
//	example-plugin
func ActionPlugins() carapace.Action {
	return carapace.ActionExecCommand("agy", "plugin", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines {
			if line == "" || strings.Contains(line, "No imported plugins") || strings.HasPrefix(line, "NAME") {
				continue
			}
			parts := strings.Fields(line)
			if len(parts) > 0 {
				vals = append(vals, parts[0])
			}
		}
		return carapace.ActionValues(vals...)
	}).Tag("plugins")
}

// ActionMcpServers completes configured MCP servers
//
//	my-server
//	example-server
func ActionMcpServers() carapace.Action {
	return carapace.ActionExecCommand("agy", "mcp", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines {
			if line == "" || strings.Contains(line, "No MCP servers") || strings.HasPrefix(line, "NAME") {
				continue
			}
			parts := strings.Fields(line)
			if len(parts) > 0 {
				vals = append(vals, parts[0])
			}
		}
		return carapace.ActionValues(vals...)
	}).Tag("mcp servers")
}
