package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var git_help_fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch from a Git remote",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_help_fetchCmd).Standalone()

	git_helpCmd.AddCommand(git_help_fetchCmd)
}
