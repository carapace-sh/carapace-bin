package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var git_help_importCmd = &cobra.Command{
	Use:   "import",
	Short: "Update repo with changes made in the underlying Git repo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_help_importCmd).Standalone()

	git_helpCmd.AddCommand(git_help_importCmd)
}
