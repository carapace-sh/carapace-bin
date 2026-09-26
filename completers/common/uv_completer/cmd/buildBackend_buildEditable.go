package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var buildBackend_buildEditableCmd = &cobra.Command{
	Use:   "build-editable",
	Short: "PEP 660 hook `build_editable`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildBackend_buildEditableCmd).Standalone()

	buildBackend_buildEditableCmd.Flags().String("metadata-directory", "", "")
	buildBackendCmd.AddCommand(buildBackend_buildEditableCmd)
	carapace.Gen(buildBackend_buildEditableCmd).FlagCompletion(carapace.ActionMap{
		"metadata-directory": carapace.ActionFiles(),
	})
	carapace.Gen(buildBackend_buildEditableCmd).PositionalCompletion(carapace.ActionFiles())
}
