package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_help_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a repository in the given directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_help_createCmd).Standalone()

	repository_helpCmd.AddCommand(repository_help_createCmd)
}
