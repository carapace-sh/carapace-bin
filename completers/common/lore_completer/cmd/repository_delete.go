package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_deleteCmd).Standalone()

	repository_deleteCmd.Flags().BoolP("help", "h", false, "Print help")
	repositoryCmd.AddCommand(repository_deleteCmd)

	carapace.Gen(repository_deleteCmd).PositionalCompletion(
		carapace.ActionValues(), // url
	)
}
