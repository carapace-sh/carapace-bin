package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var repos_createFromTemplateCmd = &cobra.Command{
	Use:     "create-from-template",
	Short:   "Create a repository based on an existing template",
	Aliases: []string{"ct"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repos_createFromTemplateCmd).Standalone()

	repos_createFromTemplateCmd.Flags().Bool("avatar", false, "copy repo avatar from template")
	repos_createFromTemplateCmd.Flags().Bool("content", false, "copy git content from template")
	repos_createFromTemplateCmd.Flags().String("description", "", "add custom description to repo")
	repos_createFromTemplateCmd.Flags().Bool("githooks", false, "copy git hooks from template")
	repos_createFromTemplateCmd.Flags().Bool("labels", false, "copy repo labels from template")
	repos_createFromTemplateCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	repos_createFromTemplateCmd.Flags().StringP("name", "n", "", "name of new repo")
	repos_createFromTemplateCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	repos_createFromTemplateCmd.Flags().StringP("owner", "O", "", "name of repo owner")
	repos_createFromTemplateCmd.Flags().Bool("private", false, "make new repo private")
	repos_createFromTemplateCmd.Flags().StringP("template", "t", "", "source template to copy from")
	repos_createFromTemplateCmd.Flags().Bool("topics", false, "copy topics from template")
	repos_createFromTemplateCmd.Flags().Bool("webhooks", false, "copy webhooks from template")
	reposCmd.AddCommand(repos_createFromTemplateCmd)

	// TODO completion
	carapace.Gen(repos_createFromTemplateCmd).FlagCompletion(carapace.ActionMap{
		"login":  tea.ActionLogins(),
		"output": tea.ActionOutputFormats(),
	})
}
