package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var removeRepoCmd = &cobra.Command{
	Use:     "remove-repo",
	Short:   "Remove a repository",
	Aliases: []string{"rr"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(removeRepoCmd).Standalone()

	rootCmd.AddCommand(removeRepoCmd)
}
