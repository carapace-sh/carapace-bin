package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var repo_createCmd = &cobra.Command{
	Use:   "create [path] [flags]",
	Short: "Create a new GitLab project/repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_createCmd).Standalone()

	repo_createCmd.Flags().String("defaultBranch", "", "Default branch of the project. Defaults to `master` if not provided.")
	repo_createCmd.Flags().StringP("description", "d", "", "Description of the new project.")
	repo_createCmd.Flags().StringP("group", "g", "", "Namespace or group for the new project. Defaults to the current user's namespace.")
	repo_createCmd.Flags().Bool("internal", false, "Make project internal: visible to any authenticated user. Default.")
	repo_createCmd.Flags().StringP("name", "n", "", "Name of the new project.")
	repo_createCmd.Flags().BoolP("private", "p", false, "Make project private: visible only to project members.")
	repo_createCmd.Flags().BoolP("public", "P", false, "Make project public: visible without any authentication.")
	repo_createCmd.Flags().Bool("readme", false, "Initialize project with `README.md`.")
	repo_createCmd.Flags().String("remoteName", "", "Remote name for the Git repository you're in. Defaults to `origin` if not provided.")
	repo_createCmd.Flags().BoolP("skipGitInit", "s", false, "Skip run 'git init'.")
	repo_createCmd.Flags().StringSliceP("tag", "t", []string{}, "The list of tags for the project.")
	repoCmd.AddCommand(repo_createCmd)

	carapace.Gen(repo_createCmd).FlagCompletion(carapace.ActionMap{
		"group": action.ActionGroups(repo_createCmd),
	})

	carapace.Gen(repo_createCmd).PositionalCompletion(
		action.ActionRepo(repo_createCmd),
	)
}
