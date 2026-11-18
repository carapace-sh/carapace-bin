package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup Atuin features",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setupCmd).Standalone()

	setupCmd.Flags().BoolP("help", "h", false, "Print help")
	setupCmd.Flags().Bool("no-agent-history", false, "Don't capture the commands your agents run")
	setupCmd.Flags().Bool("no-ctrl-r", false, "Leave ctrl-r to your shell")
	setupCmd.Flags().Bool("no-import", false, "Don't import the shell history and threads already on this machine")
	setupCmd.Flags().Bool("no-mcp", false, "Don't let your agents search Atuin over MCP")
	setupCmd.Flags().Bool("no-shell-history", false, "Don't capture your shell history")
	setupCmd.Flags().Bool("no-threads", false, "Don't capture your conversations with your agents")
	setupCmd.Flags().Bool("no-up-arrow", false, "Leave the up arrow to your shell")
	setupCmd.Flags().String("sync", "", "Where to sync: `cloud`, `off`, or your own server's address")
	setupCmd.Flags().Bool("teams", false, "Turn on Atuin for Teams")
	setupCmd.Flags().Bool("yes", false, "Set up without asking, as the flags say: sync is off unless `--sync` says otherwise")
	setupCmd.Flag("teams").Hidden = true
	rootCmd.AddCommand(setupCmd)
}
