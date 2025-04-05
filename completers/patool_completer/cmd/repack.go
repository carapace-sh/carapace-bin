package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repackCmd = &cobra.Command{
	Use:   "repack",
	Short: "repack an archive to a different format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repackCmd).Standalone()

	rootCmd.AddCommand(repackCmd)

	carapace.Gen(repackCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
