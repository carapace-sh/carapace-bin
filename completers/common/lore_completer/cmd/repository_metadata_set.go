package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_metadata_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set metadata on the repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_metadata_setCmd).Standalone()

	repository_metadata_setCmd.Flags().Bool("binary", false, "Indicator that values are paths to binary files")
	repository_metadata_setCmd.Flags().BoolP("help", "h", false, "Print help")
	repository_metadata_setCmd.Flags().Bool("numeric", false, "Indicator that values are numeric (u64)")
	repository_metadataCmd.AddCommand(repository_metadata_setCmd)
}
