package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var buildBackend_getRequiresForBuildWheelCmd = &cobra.Command{
	Use:   "get-requires-for-build-wheel",
	Short: "PEP 517 hook `get_requires_for_build_wheel`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildBackend_getRequiresForBuildWheelCmd).Standalone()

	buildBackendCmd.AddCommand(buildBackend_getRequiresForBuildWheelCmd)
}
