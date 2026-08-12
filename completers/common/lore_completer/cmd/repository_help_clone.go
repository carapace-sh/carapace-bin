package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_help_cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone a remote repository into the given path",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_help_cloneCmd).Standalone()

	repository_helpCmd.AddCommand(repository_help_cloneCmd)
}
