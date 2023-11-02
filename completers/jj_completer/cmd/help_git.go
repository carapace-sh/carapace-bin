package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_gitCmd = &cobra.Command{
	Use:   "git",
	Short: "Commands for working with the underlying Git repo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_gitCmd).Standalone()

	helpCmd.AddCommand(help_gitCmd)
}
