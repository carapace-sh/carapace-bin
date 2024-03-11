package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var displayCmd = &cobra.Command{
	Use:     "display",
	Aliases: []string{"t"},
	Short:   "display contents of the archive",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(displayCmd).Standalone()

	displayCmd.Flags().BoolS("O", "O", false, "display contents of the archive")
	addGenericOptions(displayCmd)

	carapace.Gen(displayCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
