package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new, empty change and edit it in the working copy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_newCmd).Standalone()

	helpCmd.AddCommand(help_newCmd)
}
