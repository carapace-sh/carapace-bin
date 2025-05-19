package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var modifyMetadataCmd = &cobra.Command{
	Use:   "modify-metadata",
	Short: "Modify metadata file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(modifyMetadataCmd).Standalone()

	modifyMetadataCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.AddCommand(modifyMetadataCmd)

	carapace.Gen(modifyMetadataCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
