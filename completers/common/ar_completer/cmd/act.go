package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var actCmd = &cobra.Command{
	Use:     "act",
	Aliases: []string{"s"},
	Short:   "act as ranlib",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(actCmd).Standalone()

	addGenericOptions(actCmd)

	carapace.Gen(actCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
