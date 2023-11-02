package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_git_importCmd = &cobra.Command{
	Use:   "import",
	Short: "Update repo with changes made in the underlying Git repo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_git_importCmd).Standalone()

	help_gitCmd.AddCommand(help_git_importCmd)
}
