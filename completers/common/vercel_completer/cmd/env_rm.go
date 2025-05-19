package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/vercel_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var env_rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove an Environment Variable",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(env_rmCmd).Standalone()

	envCmd.AddCommand(env_rmCmd)

	carapace.Gen(env_rmCmd).PositionalCompletion(
		action.ActionEnvironmentVariables(env_rmCmd),
		action.ActionEnvironments(),
		git.ActionRefs(git.RefOption{LocalBranches: true}),
	)
}
