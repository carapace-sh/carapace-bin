package common

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func AddInputAlteringFlags(cmd *cobra.Command) {
	cmd.Flags().String("experimental_resolved_file_instead_of_workspace", "", "If non-empty read the specified resolved file instead of the WORKSPACE file Tags: changes_inputs")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"experimental_resolved_file_instead_of_workspace": carapace.ActionFiles(),
	})
}
