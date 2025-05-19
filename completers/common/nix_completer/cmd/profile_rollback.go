package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var profile_rollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "roll back to the previous version or a specified version of a profile",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(profile_rollbackCmd).Standalone()

	profile_rollbackCmd.Flags().Bool("dry-run", false, "Show what this command would do without doing it")
	profile_rollbackCmd.Flags().String("profile", "", "The profile to operate on")
	profile_rollbackCmd.Flags().String("to", "", "The profile version to roll back to")
	profileCmd.AddCommand(profile_rollbackCmd)

	// TODO complete `to`
	carapace.Gen(profile_rollbackCmd).FlagCompletion(carapace.ActionMap{
		"profile": carapace.ActionDirectories(),
	})
}
