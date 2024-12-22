package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var profile_installCmd = &cobra.Command{
	Use:   "install",
	Short: "install a package into a profile",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(profile_installCmd).Standalone()

	profile_installCmd.Flags().String("priority", "", "The priority of the package to install")
	profile_installCmd.Flags().String("profile", "", "The profile to operate on")
	profile_installCmd.Flags().Bool("stdin", false, "Read installables from the standard input")
	profileCmd.AddCommand(profile_installCmd)

	addEvaluationFlags(profile_installCmd)
	addFlakeFlags(profile_installCmd)
	addLoggingFlags(profile_installCmd)

	carapace.Gen(profile_installCmd).FlagCompletion(carapace.ActionMap{
		"inputs-from": carapace.Batch(
			carapace.ActionDirectories(),
			nix.ActionFlakes(),
		).ToA(),
		"output-lock-file":    carapace.ActionFiles(),
		"profile":             carapace.ActionDirectories(),
		"reference-lock-file": carapace.ActionFiles("lock"),
	})

	carapace.Gen(profile_installCmd).PositionalAnyCompletion(nix.ActionInstallables())
}
