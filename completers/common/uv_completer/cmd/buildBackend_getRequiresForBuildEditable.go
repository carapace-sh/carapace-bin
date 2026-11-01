package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var buildBackend_getRequiresForBuildEditableCmd = &cobra.Command{
	Use:   "get-requires-for-build-editable",
	Short: "PEP 660 hook `get_requires_for_build_editable`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildBackend_getRequiresForBuildEditableCmd).Standalone()

	buildBackendCmd.AddCommand(buildBackend_getRequiresForBuildEditableCmd)
}
