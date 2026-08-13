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

	env_addCmd.Flags().Bool("force", false, "Force add the environment variable")
	env_addCmd.Flags().String("guidance", "", "Receive command suggestions")
	env_addCmd.Flags().Bool("no-sensitive", false, "Mark the environment variable as non-sensitive")
	env_addCmd.Flags().String("project", "", "Project name or ID")
	env_addCmd.Flags().Bool("sensitive", false, "Mark the environment variable as sensitive")
	env_addCmd.Flags().String("value", "", "Value for the environment variable")
	env_addCmd.Flags().String("visibility", "", "Visibility of the environment variable")
	env_addCmd.Flags().Bool("yes", false, "Skip confirmation")

	envCmd.AddCommand(env_addCmd)

	carapace.Gen(env_addCmd).FlagCompletion(carapace.ActionMap{
		"project":    action.ActionProjects(env_addCmd),
		"visibility": carapace.ActionValues("auto", "encrypted", "plain"),
	})

	carapace.Gen(env_addCmd).PositionalCompletion(
		carapace.ActionValues(),
		action.ActionEnvironments(),
		git.ActionRefs(git.RefOption{LocalBranches: true}),
	)
}
