package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/common"
	"github.com/spf13/cobra"
)

var profile_diffClosuresCmd = &cobra.Command{
	Use:   "diff-closures",
	Short: "show the closure difference between each version of a profile",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(profile_diffClosuresCmd).Standalone()

	profile_diffClosuresCmd.Flags().String("profile", "", "The profile to operate on")
	profileCmd.AddCommand(profile_diffClosuresCmd)

	common.AddLoggingFlags(profile_diffClosuresCmd)

	carapace.Gen(profile_diffClosuresCmd).FlagCompletion(carapace.ActionMap{
		"profile": carapace.ActionDirectories(),
	})
}
