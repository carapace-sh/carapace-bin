package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var buildBackend_prepareMetadataForBuildWheelCmd = &cobra.Command{
	Use:   "prepare-metadata-for-build-wheel",
	Short: "PEP 517 hook `prepare_metadata_for_build_wheel`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildBackend_prepareMetadataForBuildWheelCmd).Standalone()

	buildBackendCmd.AddCommand(buildBackend_prepareMetadataForBuildWheelCmd)
	carapace.Gen(buildBackend_prepareMetadataForBuildWheelCmd).PositionalCompletion(carapace.ActionFiles())
}
