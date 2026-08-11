package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var file_metadata_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set metadata on for a staged file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_metadata_setCmd).Standalone()

	file_metadata_setCmd.Flags().Bool("binary", false, "Indicator that values are paths to files")
	file_metadata_setCmd.Flags().BoolP("help", "h", false, "Print help")
	file_metadataCmd.AddCommand(file_metadata_setCmd)
}
