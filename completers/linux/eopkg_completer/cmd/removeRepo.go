package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var removeRepoCmd = &cobra.Command{
	Use:     "remove-repo <name>",
	Aliases: []string{"rr"},
	Short:   "remove a repository from the system",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(removeRepoCmd).Standalone()

	rootCmd.AddCommand(removeRepoCmd)
}
