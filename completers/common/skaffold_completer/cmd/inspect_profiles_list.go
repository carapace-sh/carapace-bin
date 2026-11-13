package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspect_profiles_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Print the list of profile names that can be invoked with the --profile flag in other skaffold commands.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspect_profiles_listCmd).Standalone()

	inspect_profiles_listCmd.Flags().String("build-env", "", "If specified as one of \"local\", \"googleCloudBuild\" or \"cluster\", then filter the output profiles list to that build environment.")
	applyCommonFlagCompletion(inspect_profiles_listCmd)
	inspect_profilesCmd.AddCommand(inspect_profiles_listCmd)
}
