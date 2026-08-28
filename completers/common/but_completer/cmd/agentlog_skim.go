package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var agentlog_skimCmd = &cobra.Command{
	Use:   "skim",
	Short: "Skim prior agent work",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(agentlog_skimCmd).Standalone()

	agentlog_skimCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	agentlogCmd.AddCommand(agentlog_skimCmd)

	carapace.Gen(agentlog_skimCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"branch", "Branch ref target",
			"review", "GitButler review target, including pull request / merge request style reviews",
			"change", "GitButler change id target",
		),
	)
}
