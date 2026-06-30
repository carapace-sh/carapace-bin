package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lCmd = &cobra.Command{
	Use:   "l",
	Short: "List contents of archive",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lCmd).Standalone()

	addCommonFlags(lCmd)
	addCommonFlagCompletions(lCmd)
	rootCmd.AddCommand(lCmd)

	carapace.Gen(lCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
