package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/vercel_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var env_rmCmd = &cobra.Command{
	Use:     "rm",
	Aliases: []string{"remove"},
	Short:   "Remove an Environment Variable",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(env_rmCmd).Standalone()

	env_rmCmd.Flags().String("project", "", "Project name or ID")
	env_rmCmd.Flags().Bool("yes", false, "Skip confirmation")

	envCmd.AddCommand(env_rmCmd)

	carapace.Gen(env_rmCmd).FlagCompletion(carapace.ActionMap{
		"project": action.ActionProjects(env_rmCmd),
	})

	carapace.Gen(env_rmCmd).PositionalCompletion(
		action.ActionEnvironmentVariables(env_rmCmd),
		action.ActionEnvironments(),
		git.ActionRefs(git.RefOption{LocalBranches: true}),
	)
}
