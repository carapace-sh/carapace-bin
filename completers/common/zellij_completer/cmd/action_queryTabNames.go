package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_queryTabNamesCmd = &cobra.Command{
	Use:   "query-tab-names",
	Short: "Query all tab names",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_queryTabNamesCmd).Standalone()

	action_queryTabNamesCmd.Flags().BoolP("help", "h", false, "Print help")
	actionCmd.AddCommand(action_queryTabNamesCmd)
}
