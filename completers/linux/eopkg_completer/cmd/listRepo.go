package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listRepoCmd = &cobra.Command{
	Use:     "list-repo",
	Short:   "List repositories",
	Aliases: []string{"lr"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listRepoCmd).Standalone()

	rootCmd.AddCommand(listRepoCmd)
}
