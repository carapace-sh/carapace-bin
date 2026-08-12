package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_help_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List repositories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_help_listCmd).Standalone()

	repository_helpCmd.AddCommand(repository_help_listCmd)
}
