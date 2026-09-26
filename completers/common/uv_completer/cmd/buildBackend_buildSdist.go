package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var buildBackend_buildSdistCmd = &cobra.Command{
	Use:   "build-sdist",
	Short: "PEP 517 hook `build_sdist`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildBackend_buildSdistCmd).Standalone()

	buildBackendCmd.AddCommand(buildBackend_buildSdistCmd)
	carapace.Gen(buildBackend_buildSdistCmd).PositionalCompletion(carapace.ActionFiles())
}
