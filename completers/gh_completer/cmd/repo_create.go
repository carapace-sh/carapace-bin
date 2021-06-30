package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var repo_createCmd = &cobra.Command{
	Use:   "create [<name>]",
	Short: "Create a new repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	repo_createCmd.Flags().BoolP("confirm", "y", false, "Skip the confirmation prompt")
	repo_createCmd.Flags().StringP("description", "d", "", "Description of the repository")
	repo_createCmd.Flags().Bool("enable-issues", true, "Enable issues in the new repository")
	repo_createCmd.Flags().Bool("enable-wiki", true, "Enable wiki in the new repository")
	repo_createCmd.Flags().StringP("gitignore", "g", "", "Specify a gitignore template for the repository")
	repo_createCmd.Flags().StringP("homepage", "h", "", "Repository home page `URL`")
	repo_createCmd.Flags().Bool("internal", false, "Make the new repository internal")
	repo_createCmd.Flags().StringP("license", "l", "", "Specify an Open Source License for the repository")
	repo_createCmd.Flags().Bool("private", false, "Make the new repository private")
	repo_createCmd.Flags().Bool("public", false, "Make the new repository public")
	repo_createCmd.Flags().StringP("team", "t", "", "The `name` of the organization team to be granted access")
	repo_createCmd.Flags().StringP("template", "p", "", "Make the new repository based on a template `repository`")
	repoCmd.AddCommand(repo_createCmd)

	carapace.Gen(repo_createCmd).FlagCompletion(carapace.ActionMap{
		"gitignore": action.ActionGitignoreTemplates(repo_createCmd),
		"license":   action.ActionLicenses(repo_createCmd),
		// TODO team
		"template": action.ActionOwnerRepositories(repo_createCmd),
	})
}
