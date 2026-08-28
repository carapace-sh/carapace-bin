package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var agentlog_publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Share local-only agent sessions for a branch, review, or change",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(agentlog_publishCmd).Standalone()

	agentlog_publishCmd.Flags().Bool("dry-run", false, "Report what would be shared without changing metadata or syncing")
	agentlog_publishCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	agentlogCmd.AddCommand(agentlog_publishCmd)
}