package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/common"
	"github.com/spf13/cobra"
)

var profile_historyCmd = &cobra.Command{
	Use:   "history",
	Short: "show all versions of a profile",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(profile_historyCmd).Standalone()

	profile_historyCmd.Flags().String("profile", "", "The profile to operate on")
	profileCmd.AddCommand(profile_historyCmd)

	common.AddEvaluationFlags(profile_historyCmd)
	common.AddLoggingFlags(profile_historyCmd)

	carapace.Gen(profile_historyCmd).FlagCompletion(carapace.ActionMap{
		"profile": carapace.ActionDirectories(),
	})
}
