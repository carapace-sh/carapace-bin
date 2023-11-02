package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_git_exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Update the underlying Git repo with changes made in the repo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_git_exportCmd).Standalone()

	help_gitCmd.AddCommand(help_git_exportCmd)
}
