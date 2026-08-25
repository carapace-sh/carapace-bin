package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pi"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pi",
	Short: "AI coding agent",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	carapace.Gen(rootCmd).Standalone()
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().String("api-key", "", "API key (defaults to env vars)")
	rootCmd.Flags().String("append-system-prompt", "", "Append text or file contents to the system prompt")
	rootCmd.Flags().BoolP("approve", "a", false, "Trust project-local files for this run")
	rootCmd.Flags().BoolP("continue", "c", false, "Continue previous session")
	rootCmd.Flags().StringP("exclude-tools", "xt", "", "Comma-separated denylist of tool names to disable")
	rootCmd.Flags().String("export", "", "Export session file to HTML and exit")
	rootCmd.Flags().StringP("extension", "e", "", "Load an extension file")
	rootCmd.Flags().String("fork", "", "Fork specific session into a new session")
	rootCmd.Flags().BoolP("help", "h", false, "Show help")
	rootCmd.Flags().String("list-models", "", "List available models (with optional fuzzy search)")
	rootCmd.Flags().String("mode", "", "Output mode")
	rootCmd.Flags().String("model", "", "Model pattern or ID")
	rootCmd.Flags().String("models", "", "Comma-separated model patterns for Ctrl+P cycling")
	rootCmd.Flags().StringP("name", "n", "", "Set session display name")
	rootCmd.Flags().BoolP("no-approve", "na", false, "Ignore project-local files for this run")
	rootCmd.Flags().BoolP("no-builtin-tools", "nbt", false, "Disable built-in tools by default but keep extension/custom tools")
	rootCmd.Flags().BoolP("no-context-files", "nc", false, "Disable AGENTS.md and CLAUDE.md discovery")
	rootCmd.Flags().BoolP("no-extensions", "ne", false, "Disable extension discovery")
	rootCmd.Flags().BoolP("no-prompt-templates", "np", false, "Disable prompt template discovery")
	rootCmd.Flags().Bool("no-session", false, "Don't save session (ephemeral)")
	rootCmd.Flags().BoolP("no-skills", "ns", false, "Disable skills discovery and loading")
	rootCmd.Flags().Bool("no-themes", false, "Disable theme discovery and loading")
	rootCmd.Flags().BoolP("no-tools", "nt", false, "Disable all tools by default")
	rootCmd.Flags().Bool("offline", false, "Disable startup network operations")
	rootCmd.Flags().BoolP("print", "p", false, "Non-interactive mode: process prompt and exit")
	rootCmd.Flags().String("prompt-template", "", "Load a prompt template file or directory")
	rootCmd.Flags().String("provider", "", "Provider name (default: google)")
	rootCmd.Flags().BoolP("resume", "r", false, "Select a session to resume")
	rootCmd.Flags().String("session", "", "Use specific session file or partial UUID")
	rootCmd.Flags().String("session-dir", "", "Directory for session storage and lookup")
	rootCmd.Flags().String("session-id", "", "Use exact project session ID")
	rootCmd.Flags().String("skill", "", "Load a skill file or directory")
	rootCmd.Flags().String("system-prompt", "", "System prompt")
	rootCmd.Flags().String("theme", "", "Load a theme file or directory")
	rootCmd.Flags().String("thinking", "", "Set thinking level")
	rootCmd.Flags().StringP("tools", "t", "", "Comma-separated allowlist of tool names to enable")
	rootCmd.Flags().String("tui-mode", "", "TUI mode")
	rootCmd.Flags().String("use-theme", "", "Set the initial interactive theme")
	rootCmd.Flags().Bool("verbose", false, "Force verbose startup")
	rootCmd.Flags().BoolP("version", "v", false, "Show version number")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"append-system-prompt": carapace.ActionFiles(),
		"exclude-tools":        pi.ActionTools().UniqueList(","),
		"export":               carapace.ActionFiles(),
		"extension":            carapace.ActionFiles(),
		"fork":                 pi.ActionSessions(),
		"list-models":          pi.ActionModels(),
		"mode": carapace.ActionValues(
			"text",
			"json",
			"rpc",
		),
		"model":           pi.ActionModels(),
		"models":          pi.ActionModels().UniqueList(","),
		"prompt-template": carapace.ActionFiles(),
		"provider":        pi.ActionProviders(),
		"session":         pi.ActionSessions(),
		"session-dir":     carapace.ActionDirectories(),
		"skill":           carapace.ActionFiles(),
		"theme":           carapace.ActionFiles(),
		"thinking": carapace.ActionValues(
			"off",
			"minimal",
			"low",
			"medium",
			"high",
			"xhigh",
			"max",
		),
		"tools": pi.ActionTools().UniqueList(","),
		"tui-mode": carapace.ActionValues(
			"regular",
			"fullscreen",
		),
		"use-theme": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles().Tag("files"),
	)
}
