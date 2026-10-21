package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/scoop"
	"github.com/spf13/cobra"
)

var whichCmd = &cobra.Command{
	Use:   "which",
	Short: "locate a shim/executable (similar to 'which' on Linux)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(whichCmd).Standalone()
	rootCmd.AddCommand(whichCmd)

	carapace.Gen(whichCmd).PositionalCompletion(
		carapace.Batch(
			scoop.ActionShims(),
			carapace.ActionExecutables(),
		).ToA().Unique(),
	)
}
