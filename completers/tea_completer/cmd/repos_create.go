package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var repos_createCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create a repository",
	Aliases: []string{"c"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repos_createCmd).Standalone()

	repos_createCmd.Flags().String("branch", "", "use custom default branch (need --init)")
	repos_createCmd.Flags().String("description", "", "add description to repo")
	repos_createCmd.Flags().String("gitignores", "", "list of gitignore templates (need --init)")
	repos_createCmd.Flags().Bool("init", false, "initialize repo")
	repos_createCmd.Flags().String("labels", "", "name of label set to add")
	repos_createCmd.Flags().String("license", "", "add license (need --init)")
	repos_createCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	repos_createCmd.Flags().String("name", "", "name of new repo")
	repos_createCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	repos_createCmd.Flags().StringP("owner", "O", "", "name of repo owner")
	repos_createCmd.Flags().Bool("private", false, "make repo private")
	repos_createCmd.Flags().String("readme", "", "use readme template (need --init)")
	repos_createCmd.Flags().Bool("template", false, "make repo a template repo")
	repos_createCmd.Flags().String("trustmodel", "", "select trust model (committer,collaborator,collaborator+committer)")
	reposCmd.AddCommand(repos_createCmd)

	// TODO completion
	carapace.Gen(repos_createCmd).FlagCompletion(carapace.ActionMap{
		"login":      tea.ActionLogins(),
		"output":     tea.ActionOutputFormats(),
		"trustmodel": carapace.ActionValues("committer", "collaborator", "collaborator+committer"),
	})
}
