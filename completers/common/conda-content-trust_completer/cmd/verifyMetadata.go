package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var verifyMetadataCmd = &cobra.Command{
	Use:   "verify-metadata",
	Short: "Verify metadata file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(verifyMetadataCmd).Standalone()

	verifyMetadataCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.AddCommand(verifyMetadataCmd)

	carapace.Gen(verifyMetadataCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)
}
