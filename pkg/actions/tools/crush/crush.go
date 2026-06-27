// Package crush provides actions for Crush AI assistant
package crush

import (
	"encoding/json"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionPlatforms completes Crush platforms
//
//	catwalk
//	hyper
func ActionPlatforms() carapace.Action {
	return carapace.ActionValues("catwalk", "github", "github-copilot", "hyper").Tag("platforms")
}

// ActionProviderSources completes provider sources
//
//	catwalk
//	hyper
func ActionProviderSources() carapace.Action {
	return carapace.ActionValues("catwalk", "hyper").Tag("provider sources")
}

type sessionEntry struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// ActionSessions completes session IDs
//
//	abc123 (my-session)
//	def456 (another-one)
func ActionSessions() carapace.Action {
	return carapace.ActionExecCommand("crush", "session", "list", "--json")(func(output []byte) carapace.Action {
		var sessions []sessionEntry
		if err := json.Unmarshal(output, &sessions); err != nil {
			return carapace.ActionMessage("failed to parse sessions: " + err.Error())
		}

		vals := make([]string, 0)
		for _, s := range sessions {
			vals = append(vals, s.ID, s.Title)
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("sessions")
}

// ActionModels completes known models from `crush models`
//
//	anthropic/claude-sonnet-4-20250514 (Anthropic)
//	openai/gpt-4o (OpenAI)
func ActionModels() carapace.Action {
	return carapace.ActionExecCommand("crush", "models")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines {
			if trimmed := strings.TrimSpace(line); trimmed != "" {
				vals = append(vals, trimmed)
			}
		}
		return carapace.ActionValues(vals...)
	}).Tag("models").UidF(Uid("model"))
}
