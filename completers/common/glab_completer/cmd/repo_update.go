package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var repo_updateCmd = &cobra.Command{
	Use:   "update [path] [flags]",
	Short: "Update an existing GitLab project or repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_updateCmd).Standalone()

	repo_updateCmd.Flags().Bool("archive", false, "Whether the project should be archived.")
	repo_updateCmd.Flags().String("defaultBranch", "", "New default branch for the project.")
	repo_updateCmd.Flags().StringP("description", "d", "", "New description for the project.")
	repoCmd.AddCommand(repo_updateCmd)

	carapace.Gen(repo_updateCmd).FlagCompletion(carapace.ActionMap{
		"defaultBranch": action.ActionBranches(repo_updateCmd),
	})

	carapace.Gen(repo_updateCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
