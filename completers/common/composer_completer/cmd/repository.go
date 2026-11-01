package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/composer_completer/cmd/action"
	"github.com/spf13/cobra"
)

var repositoryCmd = &cobra.Command{
	Use:     "repository",
	Short:   "Manages repositories",
	Aliases: []string{"repo"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repositoryCmd).Standalone()

	repositoryCmd.Flags().String("after", "", "When adding a repository, insert it after the given repository name")
	repositoryCmd.Flags().Bool("append", false, "When adding a repository, append it (lower priority) instead of prepending it")
	repositoryCmd.Flags().String("before", "", "When adding a repository, insert it before the given repository name")
	repositoryCmd.Flags().StringP("file", "f", "", "If you want to choose a different composer.json or config.json")
	repositoryCmd.Flags().BoolP("global", "g", false, "Apply command to the global config file")

	rootCmd.AddCommand(repositoryCmd)

	carapace.Gen(repositoryCmd).FlagCompletion(carapace.ActionMap{
		"after":  action.ActionRepositories(repositoryCmd),
		"before": action.ActionRepositories(repositoryCmd),
		"file":   carapace.ActionFiles(".json"),
	})

	carapace.Gen(repositoryCmd).PositionalCompletion(
		carapace.ActionValues("add", "disable", "enable", "get-url", "list", "remove", "set-url"),
		action.ActionRepositoryNames(repositoryCmd),
		action.ActionRepositoryTypeArgument(),
	)
}
