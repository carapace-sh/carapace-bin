package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var agentlog_hookCmd = &cobra.Command{
	Use:   "hook",
	Short: "Capture an agent transcript from hook input",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(agentlog_hookCmd).Standalone()

	agentlog_hookCmd.Flags().String("agent", "", "")
	agentlog_hookCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	agentlogCmd.AddCommand(agentlog_hookCmd)
}