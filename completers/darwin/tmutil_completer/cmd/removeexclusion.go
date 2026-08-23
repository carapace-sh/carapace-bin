package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var removeexclusionCmd = &cobra.Command{
	Use:   "removeexclusion",
	Short: "configure Time Machine to back up an item",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(removeexclusionCmd).Standalone()
	rootCmd.AddCommand(removeexclusionCmd)

	removeexclusionCmd.Flags().BoolP("fixed-path", "p", false, "Configure fixed-path exclusion")
	removeexclusionCmd.Flags().BoolP("volume", "v", false, "Configure volume exclusion")

	carapace.Gen(removeexclusionCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}