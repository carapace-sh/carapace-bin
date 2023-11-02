package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var git_help_pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push to a Git remote",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_help_pushCmd).Standalone()

	git_helpCmd.AddCommand(git_help_pushCmd)
}
