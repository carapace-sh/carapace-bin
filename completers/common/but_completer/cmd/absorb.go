package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var absorbCmd = &cobra.Command{
	Use:   "absorb",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(absorbCmd).Standalone()

	absorbCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(absorbCmd)

	carapace.Gen(absorbCmd).PositionalCompletion(
		carapace.Batch(
			but.ActionLocalBranches(),
			// TODO change ids
		).ToA(),
	)
}
