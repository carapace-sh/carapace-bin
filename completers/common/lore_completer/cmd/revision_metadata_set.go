package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_metadata_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set metadata on for a staged revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_metadata_setCmd).Standalone()

	revision_metadata_setCmd.Flags().Bool("binary", false, "Indicator that values are paths to files")
	revision_metadata_setCmd.Flags().BoolP("help", "h", false, "Print help")
	revision_metadataCmd.AddCommand(revision_metadata_setCmd)
}
