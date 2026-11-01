package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/eopkg"
	"github.com/spf13/cobra"
)

var repoPriorityCmd = &cobra.Command{
	Use:     "repo-priority <repo> <priority>",
	Aliases: []string{"rp"},
	Short:   "set a repository's priority",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repoPriorityCmd).Standalone()

	rootCmd.AddCommand(repoPriorityCmd)

	carapace.Gen(repoPriorityCmd).PositionalCompletion(
		eopkg.ActionRepositories(),
		carapace.ActionValues(),
	)
}
