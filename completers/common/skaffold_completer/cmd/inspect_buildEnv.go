package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspect_buildEnvCmd = &cobra.Command{
	Use:   "build-env",
	Short: "Interact with skaffold build environment definitions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspect_buildEnvCmd).Standalone()

	inspect_buildEnvCmd.PersistentFlags().StringSliceP("module", "m", nil, "Names of modules to filter target action by.")
	applyCommonFlagCompletion(inspect_buildEnvCmd)
	inspectCmd.AddCommand(inspect_buildEnvCmd)
}
