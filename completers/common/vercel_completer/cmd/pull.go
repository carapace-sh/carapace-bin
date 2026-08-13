package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/vercel_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull latest environment variables and project settings from Vercel",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pullCmd).Standalone()

	pullCmd.Flags().String("environment", "", "Deployment environment [development]")
	pullCmd.Flags().String("git-branch", "", "Specify the Git branch")
	pullCmd.Flags().Bool("prod", false, "Pull production environment")
	pullCmd.Flags().String("project", "", "Project name or ID")
	pullCmd.Flags().BoolP("yes", "y", false, "Skip questions")

	rootCmd.AddCommand(pullCmd)

	carapace.Gen(pullCmd).FlagCompletion(carapace.ActionMap{
		"environment": action.ActionEnvironments(),
		"project":     action.ActionProjects(pullCmd),
	})

	carapace.Gen(pullCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
