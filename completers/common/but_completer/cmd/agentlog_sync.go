package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var agentlog_syncCmd = &cobra.Command{
	Use:    "sync",
	Short:  "Sync GitMeta metadata",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(agentlog_syncCmd).Standalone()

	agentlog_syncCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	agentlogCmd.AddCommand(agentlog_syncCmd)
}
