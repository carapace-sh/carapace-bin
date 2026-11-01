package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var buildBackend_getRequiresForBuildSdistCmd = &cobra.Command{
	Use:   "get-requires-for-build-sdist",
	Short: "PEP 517 hook `get_requires_for_build_sdist`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildBackend_getRequiresForBuildSdistCmd).Standalone()

	buildBackendCmd.AddCommand(buildBackend_getRequiresForBuildSdistCmd)
}
