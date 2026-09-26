package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var buildBackend_prepareMetadataForBuildEditableCmd = &cobra.Command{
	Use:   "prepare-metadata-for-build-editable",
	Short: "PEP 660 hook `prepare_metadata_for_build_editable`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildBackend_prepareMetadataForBuildEditableCmd).Standalone()

	buildBackendCmd.AddCommand(buildBackend_prepareMetadataForBuildEditableCmd)
	carapace.Gen(buildBackend_prepareMetadataForBuildEditableCmd).PositionalCompletion(carapace.ActionFiles())
}
