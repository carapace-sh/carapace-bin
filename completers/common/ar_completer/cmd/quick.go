package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quickCmd = &cobra.Command{
	Use:     "quick",
	Aliases: []string{"q"},
	Short:   "quick append file(s) to the archive",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quickCmd).Standalone()

	quickCmd.Flags().BoolS("f", "f", false, "truncate inserted file names")
	addGenericOptions(quickCmd)

	carapace.Gen(quickCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
