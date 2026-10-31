package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspect_buildEnv_modifyCmd = &cobra.Command{
	Use:   "modify",
	Short: "Modify the build environment for the default pipeline or an existing profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspect_buildEnv_modifyCmd).Standalone()

	inspect_buildEnv_modifyCmd.PersistentFlags().StringP("profile", "p", "", "If specified then the modify actions are performed on all pipelines from profiles that match flag value. It must match at least one existing profile name. If unspecified then only the default pipeline in the target 'skaffold.yaml' file are modified")
	inspect_buildEnv_modifyCmd.PersistentFlags().Bool("strict", true, "If set to 'false' then do not fail on mismatch of build env type or missing definitions. If set to 'true' then all default pipelines or those matching the `--profile` flag need to match the build env type exactly. Defaults to 'true'")
	applyCommonFlagCompletion(inspect_buildEnv_modifyCmd)
	inspect_buildEnvCmd.AddCommand(inspect_buildEnv_modifyCmd)
}
