package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var agentlog_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show a session, or one turn in detail",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(agentlog_showCmd).Standalone()

	agentlog_showCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	agentlog_showCmd.Flags().String("limit", "", "Maximum turns or turn records to return")
	agentlog_showCmd.Flags().String("turn", "", "Show detailed records for this turn key")
	agentlogCmd.AddCommand(agentlog_showCmd)
}