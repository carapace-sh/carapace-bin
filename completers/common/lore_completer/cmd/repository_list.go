package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List repositories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_listCmd).Standalone()

	repository_listCmd.Flags().BoolP("help", "h", false, "Print help")
	repositoryCmd.AddCommand(repository_listCmd)
}
