package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_help_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_help_deleteCmd).Standalone()

	repository_helpCmd.AddCommand(repository_help_deleteCmd)
}
