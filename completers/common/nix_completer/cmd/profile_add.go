package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var profile_addCmd = &cobra.Command{
	Use:     "add",
	Short:   "add a package to a profile",
	Aliases: []string{"install"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(profile_addCmd).Standalone()

	profile_addCmd.Flags().String("priority", "", "The priority of the package to add")
	profile_addCmd.Flags().String("profile", "", "The profile to operate on")
	profile_addCmd.Flags().Bool("stdin", false, "Read installables from the standard input")
	profileCmd.AddCommand(profile_addCmd)

	common.AddEvaluationFlags(profile_addCmd)
	common.AddFlakeFlags(profile_addCmd)
	common.AddLoggingFlags(profile_addCmd)

	carapace.Gen(profile_addCmd).FlagCompletion(carapace.ActionMap{
		"profile": carapace.ActionDirectories(),
	})

	carapace.Gen(profile_addCmd).PositionalAnyCompletion(nix.ActionInstallables())
}
