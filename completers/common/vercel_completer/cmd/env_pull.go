package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/vercel_completer/cmd/action"
	"github.com/spf13/cobra"
)

var env_pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull all Development Environment Variables from the cloud and write to a file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(env_pullCmd).Standalone()

	env_pullCmd.Flags().String("environment", "", "Deployment environment [development]")
	env_pullCmd.Flags().String("git-branch", "", "Specify the Git branch")
	env_pullCmd.Flags().String("id", "", "Identifier of the environment to pull")
	env_pullCmd.Flags().Bool("prod", false, "Pull production environment")
	env_pullCmd.Flags().String("project", "", "Project name or ID")
	env_pullCmd.Flags().BoolP("yes", "y", false, "Skip questions")

	envCmd.AddCommand(env_pullCmd)

	carapace.Gen(env_pullCmd).FlagCompletion(carapace.ActionMap{
		"environment": action.ActionEnvironments(),
	})

	carapace.Gen(env_pullCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
