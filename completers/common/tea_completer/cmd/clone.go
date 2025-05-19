package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var cloneCmd = &cobra.Command{
	Use:     "clone",
	Short:   "Clone a repository locally",
	GroupID: "HELPERS",
	Aliases: []string{"C"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloneCmd).Standalone()

	cloneCmd.Flags().StringP("depth", "d", "", "num commits to fetch, defaults to all")
	cloneCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	rootCmd.AddCommand(cloneCmd)

	carapace.Gen(cloneCmd).FlagCompletion(carapace.ActionMap{
		"login": tea.ActionLogins(),
	})

	carapace.Gen(cloneCmd).PositionalCompletion(
		carapace.ActionValues(), // TODO repo slug
		carapace.ActionDirectories(),
	)
}
