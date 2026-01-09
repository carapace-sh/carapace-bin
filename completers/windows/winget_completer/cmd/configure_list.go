package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configure_listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "Shows configuration history",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configure_listCmd).Standalone()

	configure_listCmd.Flags().BoolP("history", "h", false, "Select items from history")
	configure_listCmd.Flags().StringP("output", "o", "", "File where the result is to be written")
	configure_listCmd.Flags().Bool("remove", false, "Remove the item from history")
	configureCmd.AddCommand(configure_listCmd)

	carapace.Gen(configure_listCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionFiles(),
	})
}
