package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "claude",
	Short: "Claude Code",
	Long:  "https://code.claude.com/docs/en/cli-reference",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringSlice("add-dir", nil, "Additional directories to allow tool access to")
	rootCmd.Flags().String("advisor", "", "Enable the server-side advisor tool for this session")
	rootCmd.Flags().String("agent", "", "Agent for the current session")
	rootCmd.Flags().String("agents", "", "JSON object defining custom agents, or with --print the path to a file that holds one…")
	rootCmd.Flags().Bool("allow-dangerously-skip-permissions", false, "Enable bypassing all permission checks as an option, without it being enabled by default")
	rootCmd.Flags().StringSlice("allowed-tools", nil, "Comma or space-separated list of tool names to allow (e.g. \"Bash(git *) Edit\")")
	rootCmd.Flags().StringSlice("allowedTools", nil, "Comma or space-separated list of tool names to allow (e.g. \"Bash(git *) Edit\")")
	rootCmd.Flags().String("append-subagent-system-prompt", "", "Append text to every subagent system prompt (print mode)")
	rootCmd.Flags().String("append-subagent-system-prompt-file", "", "Append the contents of a file to subagent system prompts")
	rootCmd.Flags().String("append-system-prompt", "", "Append a system prompt to the default system prompt")
	rootCmd.Flags().String("append-system-prompt-file", "", "Append the contents of a file to the default system prompt")
	rootCmd.Flags().String("autocompact", "", "Auto-compact window size (auto, or 100k–1M tokens)")
	rootCmd.Flags().Bool("ax-screen-reader", false, "Render screen-reader friendly output (flat text, no decorative borders or animations)")
	rootCmd.Flags().Bool("background", false, "Start the session in the background and return immediately")
	rootCmd.Flags().Bool("bare", false, "Minimal mode: skip hooks…")
	rootCmd.Flags().StringSlice("betas", nil, "Beta headers to include in API requests (API key users only)")
	rootCmd.Flags().Bool("bg", false, "Start the session in the background and return immediately")
	rootCmd.Flags().Bool("brief", false, "Enable SendUserMessage tool for agent-to-user communication")
	rootCmd.Flags().StringSlice("channels", nil, "MCP servers whose channel notifications to listen for (research preview)")
	rootCmd.Flags().Bool("chrome", false, "Enable Claude in Chrome integration")
	rootCmd.Flags().String("cloud", "", "Create a cloud session with the given description, or attach to an existing one by session ID or…")
	rootCmd.Flags().BoolP("continue", "c", false, "Continue the most recent conversation in the current directory")
	rootCmd.Flags().StringSlice("dangerously-load-development-channels", nil, "Enable channels not on the approved allowlist")
	rootCmd.Flags().Bool("dangerously-skip-permissions", false, "Bypass all permission checks")
	rootCmd.Flags().StringP("debug", "d", "", "Enable debug mode with optional category filtering…")
	rootCmd.Flags().String("debug-file", "", "Write debug logs to a specific file path (implicitly enables debug mode)")
	rootCmd.Flags().Bool("desktop", false, "Open in the Claude Desktop app instead of the terminal…")
	rootCmd.Flags().Bool("disable-slash-commands", false, "Disable all skills and commands")
	rootCmd.Flags().StringSlice("disallowed-tools", nil, "Comma or space-separated list of tool names to deny (e.g. \"Bash(git *) Edit\")")
	rootCmd.Flags().StringSlice("disallowedTools", nil, "Comma or space-separated list of tool names to deny (e.g. \"Bash(git *) Edit\")")
	rootCmd.Flags().String("effort", "", "Effort level for the current session (low, medium, high, xhigh, max)")
	rootCmd.Flags().String("environment", "", "Create a new cloud session that runs on the given self-hosted environment (ccpool_...)")
	rootCmd.Flags().Bool("exclude-dynamic-system-prompt-sections", false, "Move per-machine sections (cwd, env info, memory paths, git status) from the system prompt into…")
	rootCmd.Flags().Bool("exec", false, "Run a shell command as a background job (with --bg)")
	rootCmd.Flags().String("fallback-model", "", "Enable automatic fallback to specified model(s) when the default model is overloaded or not…")
	rootCmd.Flags().StringSlice("file", nil, "File resources to download at startup")
	rootCmd.Flags().Bool("fork-session", false, "When resuming, create a new session ID instead of reusing the original…")
	rootCmd.Flags().Bool("forward-subagent-text", false, "Forward subagent text and thinking blocks as assistant/user messages with parent_tool_use_id set…")
	rootCmd.Flags().String("from-pr", "", "Resume a session linked to a PR by PR number/URL, or open interactive picker with optional search…")
	rootCmd.Flags().BoolP("help", "h", false, "Display help for command")
	rootCmd.Flags().Bool("ide", false, "Automatically connect to IDE on startup if exactly one valid IDE is available")
	rootCmd.Flags().Bool("include-hook-events", false, "Include all hook lifecycle events in the output stream (only works with --output-format=stream-json)")
	rootCmd.Flags().Bool("include-partial-messages", false, "Include partial message chunks as they arrive…")
	rootCmd.Flags().Bool("init", false, "Run Setup hooks with the init matcher before the session")
	rootCmd.Flags().Bool("init-only", false, "Run Setup and SessionStart hooks, then exit")
	rootCmd.Flags().String("input-format", "", "Input format (only works with --print): \"text\" (default), or \"stream-json\"…")
	rootCmd.Flags().String("json-schema", "", "JSON Schema for structured output validation (only works with --print)")
	rootCmd.Flags().Bool("maintenance", false, "Run Setup hooks with the maintenance matcher before the session")
	rootCmd.Flags().String("max-budget-usd", "", "Maximum dollar amount to spend on API calls (only works with --print)")
	rootCmd.Flags().String("max-turns", "", "Limit the number of agentic turns (print mode)")
	rootCmd.Flags().StringSlice("mcp-config", nil, "Load MCP servers from JSON files or strings (space-separated)")
	rootCmd.Flags().String("model", "", "Model for the current session")
	rootCmd.Flags().StringP("name", "n", "", "Set a display name for this session (shown in the prompt box, /resume picker, and terminal title)")
	rootCmd.Flags().Bool("no-chrome", false, "Disable Claude in Chrome integration")
	rootCmd.Flags().Bool("no-session-persistence", false, "Disable session persistence - sessions will not be saved to disk and cannot be resumed…")
	rootCmd.Flags().String("output-format", "", "Output format (only works with --print): \"text\" (default), \"json\" (single result), or…")
	rootCmd.Flags().String("permission-mode", "", "Permission mode to use for the session")
	rootCmd.Flags().String("permission-prompt-tool", "", "MCP tool to handle permission prompts in non-interactive mode")
	rootCmd.Flags().String("permission-prompts", "", "Who answers permission prompts with --print: \"host\" (the SDK host or --permission-prompt-tool) or…")
	rootCmd.Flags().String("plugin-dir", "", "Load a plugin from a directory or .zip for this session only; a folder of plugins loads each child…")
	rootCmd.Flags().String("plugin-url", "", "Fetch a plugin .zip from a URL for this session only (repeatable: --plugin-url A --plugin-url B)")
	rootCmd.Flags().BoolP("print", "p", false, "Print response and exit (useful for pipes)")
	rootCmd.Flags().String("prompt-suggestions", "", "Enable prompt suggestions")
	rootCmd.Flags().String("rc", "", "Alias for --remote-control")
	rootCmd.Flags().String("ref", "", "Base the cloud session checkout on a named ref (with --environment)")
	rootCmd.Flags().String("remote", "", "Deprecated alias for --cloud")
	rootCmd.Flags().String("remote-control", "", "Start an interactive session with Remote Control enabled (optionally named)")
	rootCmd.Flags().String("remote-control-session-name-prefix", "", "Prefix for auto-generated Remote Control session names")
	rootCmd.Flags().Bool("replay-user-messages", false, "Re-emit user messages from stdin back on stdout for acknowledgment…")
	rootCmd.Flags().Bool("restricted", false, "Restricted mode: removes the built-in tools that run commands or code…")
	rootCmd.Flags().StringP("resume", "r", "", "Resume a conversation by session ID or name, or open interactive picker with optional search term")
	rootCmd.Flags().Bool("safe-mode", false, "Start with all customizations…")
	rootCmd.Flags().String("session-id", "", "Use a specific session ID for the conversation (must be a valid UUID)")
	rootCmd.Flags().String("setting-sources", "", "Comma-separated list of setting sources to load (user, project, local)")
	rootCmd.Flags().String("settings", "", "Path to a settings JSON file or a JSON string to load additional settings from")
	rootCmd.Flags().Bool("strict-mcp-config", false, "Only use MCP servers from --mcp-config, ignoring all other MCP configurations")
	rootCmd.Flags().String("system-prompt", "", "System prompt to use for the session")
	rootCmd.Flags().String("system-prompt-file", "", "Replace the system prompt with the contents of a file")
	rootCmd.Flags().String("system-prompt-snapshot", "", "Record the system prompt once per conversation and reuse it verbatim on every request and resume")
	rootCmd.Flags().String("teammate-mode", "", "How agent team teammates display")
	rootCmd.Flags().String("teleport", "", "Resume a teleport session, optionally specify session ID")
	rootCmd.Flags().Bool("tmux", false, "Create a tmux session for the worktree (requires --worktree)")
	rootCmd.Flags().StringSlice("tools", nil, "Specify the list of available tools from the built-in set")
	rootCmd.Flags().Bool("verbose", false, "Override verbose mode setting from config")
	rootCmd.Flags().BoolP("version", "v", false, "Output the version number")
	rootCmd.Flags().StringP("worktree", "w", "", "Create a new git worktree for this session (optionally specify a name)")
	rootCmd.Flag("allowedTools").Hidden = true
	rootCmd.Flag("bg").Hidden = true
	rootCmd.Flag("cloud").NoOptDefVal = " "
	rootCmd.Flag("debug").NoOptDefVal = " "
	rootCmd.Flag("disallowedTools").Hidden = true
	rootCmd.Flag("from-pr").NoOptDefVal = " "
	rootCmd.Flag("prompt-suggestions").NoOptDefVal = " "
	rootCmd.Flag("rc").NoOptDefVal = " "
	rootCmd.Flag("remote").NoOptDefVal = " "
	rootCmd.Flag("remote").Hidden = true
	rootCmd.Flag("remote-control").NoOptDefVal = " "
	rootCmd.Flag("resume").NoOptDefVal = " "
	rootCmd.Flag("teleport").NoOptDefVal = " "
	rootCmd.Flag("worktree").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"add-dir":                            carapace.ActionDirectories(),
		"advisor":                            carapace.ActionValues("fable", "haiku", "opus", "sonnet"),
		"agents":                             carapace.ActionFiles(),
		"append-subagent-system-prompt-file": carapace.ActionFiles(),
		"append-system-prompt-file":          carapace.ActionFiles(),
		"debug-file":                         carapace.ActionFiles(),
		"effort":                             carapace.ActionValues("high", "low", "max", "medium", "ultracode", "xhigh"),
		"fallback-model":                     carapace.ActionValues("fable", "haiku", "opus", "sonnet").UniqueList(","),
		"input-format":                       carapace.ActionValues("stream-json", "text"),
		"mcp-config":                         carapace.ActionFiles(),
		"model":                              carapace.ActionValues("fable", "haiku", "opus", "sonnet"),
		"output-format":                      carapace.ActionValues("json", "stream-json", "text"),
		"permission-mode":                    carapace.ActionValues("acceptEdits", "auto", "bypassPermissions", "default", "dontAsk", "manual", "plan"),
		"permission-prompts":                 carapace.ActionValues("host", "none"),
		"plugin-dir":                         carapace.ActionFiles(),
		"prompt-suggestions":                 carapace.ActionValues("0", "1", "false", "no", "off", "on", "true", "yes"),
		"setting-sources":                    carapace.ActionValues("local", "project", "user").UniqueList(","),
		"settings":                           carapace.ActionFiles(),
		"system-prompt-file":                 carapace.ActionFiles(),
		"system-prompt-snapshot":             carapace.ActionValues("off", "on"),
		"teammate-mode":                      carapace.ActionValues("auto", "in-process", "iterm2", "tmux"),
	})
}
