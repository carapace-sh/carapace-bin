package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/vercel_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var env_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an Environment Variable",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(env_addCmd).Standalone()

	envCmd.AddCommand(env_addCmd)

	carapace.Gen(env_addCmd).PositionalCompletion(
		carapace.ActionValues(),
		action.ActionEnvironments(),
		git.ActionRefs(git.RefOption{LocalBranches: true}),
	)
}
