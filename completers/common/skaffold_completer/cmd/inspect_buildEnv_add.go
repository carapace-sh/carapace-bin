package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspect_buildEnv_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new build environment to the default pipeline or to a new or existing profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspect_buildEnv_addCmd).Standalone()

	inspect_buildEnv_addCmd.PersistentFlags().StringP("profile", "p", "", "Profile name to add the new build env definition in. If the profile name doesn't exist then the profile will be created in all the target configs. If this flag is not specified then the build env is added to the default pipeline of the target configs.")
	applyCommonFlagCompletion(inspect_buildEnv_addCmd)
	inspect_buildEnvCmd.AddCommand(inspect_buildEnv_addCmd)
}
