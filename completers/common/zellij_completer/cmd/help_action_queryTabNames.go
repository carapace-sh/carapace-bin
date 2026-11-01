package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_queryTabNamesCmd = &cobra.Command{
	Use:   "query-tab-names",
	Short: "Query all tab names",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_queryTabNamesCmd).Standalone()

	help_actionCmd.AddCommand(help_action_queryTabNamesCmd)
}
