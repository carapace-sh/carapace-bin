package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var file_help_resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset changes to a path or file to the current revision, discarding your local changes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_help_resetCmd).Standalone()

	file_helpCmd.AddCommand(file_help_resetCmd)
}
