package mcp

type mcpTool struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	InputSchema map[string]any `json:"inputSchema"`
}

var tools = []mcpTool{
	{
		Name:        "complete_command",
		Description: "Return context‑aware, dynamic completions for shell commands.",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"args": map[string]any{
					"type": "array",
					"items": map[string]any{
						"type": "string",
					},
					"description": "command line arguments",
				},
				"executable": map[string]any{
					"type":        "string",
					"description": "path to the executable providing the completion (requires bridge; differs from args[0] which is the command name)",
				},
				"bridge": map[string]any{
					"type":        "string",
					"description": "bridge providing the completion (e.g. carapace-bin, cobra, zsh, fish, bash, argcomplete, click)",
				},
			},
			"required":             []string{"args"},
			"additionalProperties": false,
		},
	},
	{
		Name:        "complete_macro",
		Description: "Return context‑aware, dynamic completions for a carapace macro.",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"macro": map[string]any{
					"type":        "string",
					"description": "macro name (e.g. tools.git.Refs)",
				},
				"args": map[string]any{
					"type": "array",
					"items": map[string]any{
						"type": "string",
					},
					"description": "arguments to complete",
				},
				"executable": map[string]any{
					"type":        "string",
					"description": "path to the carapace executable providing the macro",
				},
			},
			"required":             []string{"macro", "args"},
			"additionalProperties": false,
		},
	},
	{
		Name:        "list_macros",
		Description: "List available macros and their signatures.",
		InputSchema: map[string]any{
			"type":                 "object",
			"properties":           map[string]any{},
			"additionalProperties": false,
		},
	},
	{
		Name:        "codegen",
		Description: "Generate Go code from a YAML spec file.",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"path": map[string]any{
					"type":        "string",
					"description": "Path to the YAML spec file",
				},
			},
			"required":             []string{"path"},
			"additionalProperties": false,
		},
	},
}
