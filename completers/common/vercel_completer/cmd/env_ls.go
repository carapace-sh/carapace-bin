package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/vercel_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var env_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all variables for the specified Environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(env_lsCmd).Standalone()

	envCmd.AddCommand(env_lsCmd)

	carapace.Gen(env_lsCmd).PositionalCompletion(
		action.ActionEnvironments(),
		git.ActionRefs(git.RefOption{LocalBranches: true}),
	)
}
