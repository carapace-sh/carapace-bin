package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_newCmd = &cobra.Command{
	Use:   "new",
	Short: "Insert a blank commit before the specified commit, or at the top of a stack",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_newCmd).Standalone()

	helpCmd.AddCommand(help_newCmd)
}
