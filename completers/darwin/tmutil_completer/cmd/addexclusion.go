package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var addexclusionCmd = &cobra.Command{
	Use:   "addexclusion",
	Short: "configure an exclusion from Time Machine backups",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addexclusionCmd).Standalone()
	rootCmd.AddCommand(addexclusionCmd)

	addexclusionCmd.Flags().BoolP("fixed-path", "p", false, "Configure fixed-path exclusion")
	addexclusionCmd.Flags().BoolP("volume", "v", false, "Configure volume exclusion")

	carapace.Gen(addexclusionCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}