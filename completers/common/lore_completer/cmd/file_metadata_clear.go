package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var file_metadata_clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear metadata for a staged file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_metadata_clearCmd).Standalone()

	file_metadata_clearCmd.Flags().BoolP("help", "h", false, "Print help")
	file_metadataCmd.AddCommand(file_metadata_clearCmd)
}
