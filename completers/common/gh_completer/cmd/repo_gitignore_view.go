package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var repo_gitignore_viewCmd = &cobra.Command{
	Use:   "view <template>",
	Short: "View an available repository gitignore template",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_gitignore_viewCmd).Standalone()

	repo_gitignoreCmd.AddCommand(repo_gitignore_viewCmd)

	carapace.Gen(repo_gitignore_viewCmd).PositionalCompletion(
		action.ActionGitignoreTemplates(repo_gitignore_viewCmd),
	)
}
