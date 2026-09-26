package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var buildBackend_buildWheelCmd = &cobra.Command{
	Use:   "build-wheel",
	Short: "PEP 517 hook `build_wheel`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildBackend_buildWheelCmd).Standalone()

	buildBackend_buildWheelCmd.Flags().String("metadata-directory", "", "")
	buildBackendCmd.AddCommand(buildBackend_buildWheelCmd)
	carapace.Gen(buildBackend_buildWheelCmd).FlagCompletion(carapace.ActionMap{
		"metadata-directory": carapace.ActionFiles(),
	})
	carapace.Gen(buildBackend_buildWheelCmd).PositionalCompletion(carapace.ActionFiles())
}
