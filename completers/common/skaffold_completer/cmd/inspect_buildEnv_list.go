package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspect_buildEnv_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Print the list of active build environments.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspect_buildEnv_listCmd).Standalone()

	inspect_buildEnv_listCmd.Flags().StringSliceP("profile", "p", nil, "Profile names to activate")
	inspect_buildEnv_listCmd.Flags().Bool("propagate-profiles", true, "Setting '--propagate-profiles=false' disables propagating profiles set by the '--profile' flag across config dependencies. This mean that only profiles defined directly in the target 'skaffold.yaml' file are activated.")
	applyCommonFlagCompletion(inspect_buildEnv_listCmd)
	inspect_buildEnvCmd.AddCommand(inspect_buildEnv_listCmd)
}
