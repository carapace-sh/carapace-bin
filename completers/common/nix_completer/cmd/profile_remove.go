package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/common"
	"github.com/spf13/cobra"
)

var profile_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove installed packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(profile_removeCmd).Standalone()

	profile_removeCmd.Flags().String("profile", "", "The profile to operate on")
	profileCmd.AddCommand(profile_removeCmd)

	common.AddEvaluationFlags(profile_removeCmd)
	common.AddLoggingFlags(profile_removeCmd)

	carapace.Gen(profile_removeCmd).FlagCompletion(carapace.ActionMap{
		"profile": carapace.ActionDirectories(),
	})

	// TODO positional completion
}
