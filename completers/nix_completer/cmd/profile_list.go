package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var profile_listCmd = &cobra.Command{
	Use:   "list",
	Short: "list installed packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(profile_listCmd).Standalone()

	profile_listCmd.Flags().String("profile", "", "The profile to operate on")
	profileCmd.AddCommand(profile_listCmd)

	addEvaluationFlags(profile_listCmd)
	addLoggingFlags(profile_listCmd)

	carapace.Gen(profile_listCmd).FlagCompletion(carapace.ActionMap{
		"profile": carapace.ActionDirectories(),
	})
}
