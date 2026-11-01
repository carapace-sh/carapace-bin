package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var agentsCmd = &cobra.Command{
	Use:   "agents",
	Short: "Manage background agents",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(agentsCmd).Standalone()

	agentsCmd.Flags().String("add-dir", "", "Additional directory to allow tool access to in dispatched sessions (repeatable)")
	agentsCmd.Flags().String("agent", "", "Default agent for sessions dispatched from agent view")
	agentsCmd.Flags().Bool("all", false, "With --json: also include completed background sessions")
	agentsCmd.Flags().Bool("allow-dangerously-skip-permissions", false, "Make bypass-permissions mode available to dispatched sessions without defaulting to it")
	agentsCmd.Flags().String("cwd", "", "Show only background sessions started under <path>")
	agentsCmd.Flags().Bool("dangerously-skip-permissions", false, "Alias for --permission-mode bypassPermissions")
	agentsCmd.Flags().String("effort", "", "Default effort level for sessions dispatched from agent view")
	agentsCmd.Flags().BoolP("help", "h", false, "Display help for command")
	agentsCmd.Flags().Bool("json", false, "Print active sessions (interactive and background) as a JSON array and exit…")
	agentsCmd.Flags().String("mcp-config", "", "MCP server configuration to apply to dispatched sessions (repeatable)")
	agentsCmd.Flags().String("model", "", "Default model for sessions dispatched from agent view")
	agentsCmd.Flags().String("permission-mode", "", "Default permission mode for sessions dispatched from agent view")
	agentsCmd.Flags().String("plugin-dir", "", "Load plugins from specified directory for the agent view and dispatched sessions; a folder of…")
	agentsCmd.Flags().Bool("restricted", false, "Start dispatched sessions in restricted mode")
	agentsCmd.Flags().String("setting-sources", "", "Comma-separated list of setting sources to load (user, project, local)")
	agentsCmd.Flags().String("settings", "", "Settings file or JSON string to apply to the agent view and dispatched sessions")
	agentsCmd.Flags().Bool("strict-mcp-config", false, "Only use MCP servers from --mcp-config in dispatched sessions")
	rootCmd.AddCommand(agentsCmd)

	carapace.Gen(agentsCmd).FlagCompletion(carapace.ActionMap{
		"add-dir":         carapace.ActionDirectories(),
		"cwd":             carapace.ActionDirectories(),
		"effort":          carapace.ActionValues("high", "low", "max", "medium", "ultracode", "xhigh"),
		"mcp-config":      carapace.ActionFiles(),
		"model":           carapace.ActionValues("fable", "haiku", "opus", "sonnet"),
		"permission-mode": carapace.ActionValues("default"),
		"plugin-dir":      carapace.ActionFiles(),
		"setting-sources": carapace.ActionValues("local", "project", "user").UniqueList(","),
		"settings":        carapace.ActionFiles(),
	})
}
