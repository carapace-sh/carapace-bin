package gh

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionAgents completes known agents
//
//	github-copilot
//	crush
func ActionAgents() carapace.Action {
	return carapace.Batch(
		carapace.ActionValuesDescribed(
			"antigravity", "Antigravity",
			"claude-code", "Claude Code",
			"codex", "Codex",
			"cursor", "Cursor",
			"gemini-cli", "Gemini CLI",
			"github-copilot", "GitHub Copilot",
		).Style(style.Blue),
		carapace.ActionValuesDescribed(
			"adal", "AdaL",
			"amp", "Amp",
			"augment", "Augment",
			"bob", "IBM Bob",
			"cline", "Cline",
			"codebuddy", "CodeBuddy",
			"command-code", "Command Code",
			"continue", "Continue",
			"cortex", "Cortex Code",
			"crush", "Crush",
			"deepagents", "Deep Agents",
			"droid", "Droid",
			"firebender", "Firebender",
			"goose", "Goose",
			"iflow-cli", "iFlow CLI",
			"junie", "Junie",
			"kilo", "Kilo Code",
			"kimi-cli", "Kimi Code CLI",
			"kiro-cli", "Kiro CLI",
			"kode", "Kode",
			"mcpjam", "MCPJam",
			"mistral-vibe", "Mistral Vibe",
			"mux", "Mux",
			"neovate", "Neovate",
			"openclaw", "OpenClaw",
			"opencode", "OpenCode",
			"openhands", "OpenHands",
			"pi", "Pi",
			"pochi", "Pochi",
			"qoder", "Qoder",
			"qwen-code", "Qwen Code",
			"replit", "Replit",
			"roo", "Roo Code",
			"trae", "Trae",
			"trae-cn", "Trae CN",
			"universal", "Universal",
			"warp", "Warp",
			"windsurf", "Windsurf",
			"zencoder", "Zencoder",
		),
	).ToA().Tag("agents")
}
