package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var git_renameCmd = &cobra.Command{
	Use:   "rename REPO NEWNAME",
	Short: "Rename a Git repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_renameCmd).Standalone()

	gitCmd.AddCommand(git_renameCmd)
}
