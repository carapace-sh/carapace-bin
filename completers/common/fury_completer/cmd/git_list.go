package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var git_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List repos in this account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_listCmd).Standalone()

	gitCmd.AddCommand(git_listCmd)
}
