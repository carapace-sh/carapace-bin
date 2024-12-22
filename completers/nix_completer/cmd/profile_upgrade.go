package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var profile_upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "upgrade packages using their most recent flake",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(profile_upgradeCmd).Standalone()

	profile_upgradeCmd.Flags().String("profile", "", "The profile to operate on")
	profileCmd.AddCommand(profile_upgradeCmd)

	addEvaluationFlags(profile_upgradeCmd)
	addFlakeFlags(profile_upgradeCmd)
	addLoggingFlags(profile_upgradeCmd)

	carapace.Gen(profile_upgradeCmd).FlagCompletion(carapace.ActionMap{
		"inputs-from": carapace.Batch(
			carapace.ActionDirectories(),
			nix.ActionFlakes(),
		).ToA(),
		"output-lock-file":    carapace.ActionFiles(),
		"profile":             carapace.ActionDirectories(),
		"reference-lock-file": carapace.ActionFiles("lock"),
	})
}
