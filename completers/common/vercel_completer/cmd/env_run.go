package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/vercel_completer/cmd/action"
	"github.com/spf13/cobra"
)

var env_runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a command with environment variables from the linked Vercel project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(env_runCmd).Standalone()

	env_runCmd.Flags().String("environment", "", "Deployment environment")
	env_runCmd.Flags().String("git-branch", "", "Specify the Git branch")
	env_runCmd.Flags().String("project", "", "Project name or ID")

	envCmd.AddCommand(env_runCmd)

	carapace.Gen(env_runCmd).FlagCompletion(carapace.ActionMap{
		"environment": action.ActionEnvironments(),
	})

	carapace.Gen(env_runCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
