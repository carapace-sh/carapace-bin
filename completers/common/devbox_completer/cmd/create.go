package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create [dir] --template <template>",
	Short: "Initialize a directory as a devbox project using a template",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createCmd).Standalone()

	createCmd.Flags().StringP("repo", "r", "", "Git repository HTTPS URL to import template files from. Example: https://github.com/jetify-com/devbox")
	createCmd.Flags().Bool("show-all", false, "show all available templates")
	createCmd.Flags().StringP("subdir", "s", "", "Subdirectory of the Git repository in which the template files reside. Example: examples/tutorial")
	createCmd.Flags().StringP("template", "t", "", "template to initialize the project with")
	createCmd.Flag("repo").Hidden = true
	createCmd.Flag("subdir").Hidden = true
	rootCmd.AddCommand(createCmd)

	// TODO flagcompletion
	carapace.Gen(createCmd).FlagCompletion(carapace.ActionMap{
		"repo": git.ActionRepositorySearch(git.SearchOpts{}.Default()),
	})

	carapace.Gen(createCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
