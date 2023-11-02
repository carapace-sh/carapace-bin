package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var git_help_exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Update the underlying Git repo with changes made in the repo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_help_exportCmd).Standalone()

	git_helpCmd.AddCommand(git_help_exportCmd)
}
