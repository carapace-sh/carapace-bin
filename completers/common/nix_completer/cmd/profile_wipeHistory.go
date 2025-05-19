package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var profile_wipeHistoryCmd = &cobra.Command{
	Use:   "wipe-history",
	Short: "delete non-current versions of a profile",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(profile_wipeHistoryCmd).Standalone()

	profile_wipeHistoryCmd.Flags().Bool("dry-run", false, "Show what this command would do without doing it")
	profile_wipeHistoryCmd.Flags().String("older-than", "", "Delete versions older than the specified age")
	profile_wipeHistoryCmd.Flags().String("profile", "", "The profile to operate on")
	profileCmd.AddCommand(profile_wipeHistoryCmd)

	carapace.Gen(profile_wipeHistoryCmd).FlagCompletion(carapace.ActionMap{
		"profile": carapace.ActionDirectories(),
	})
}
