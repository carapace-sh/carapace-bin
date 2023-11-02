package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a commit in the working copy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_editCmd).Standalone()

	helpCmd.AddCommand(help_editCmd)
}
