package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/vercel_completer/cmd/action"
	"github.com/spf13/cobra"
)

var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "Link a local directory to a Vercel project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(linkCmd).Standalone()

	linkCmd.Flags().BoolP("confirm", "c", false, "(deprecated)")
	linkCmd.Flags().StringP("project", "p", "", "Set the project name or ID to link")
	linkCmd.Flags().StringP("repo", "r", "", "Link multiple projects from the Git repository (alpha)")
	linkCmd.Flags().String("team", "", "Set the team ID or slug")
	linkCmd.Flags().BoolP("yes", "y", false, "Skip questions")

	rootCmd.AddCommand(linkCmd)

	carapace.Gen(linkCmd).FlagCompletion(carapace.ActionMap{
		"project": action.ActionProjects(linkCmd),
	})

	carapace.Gen(linkCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
