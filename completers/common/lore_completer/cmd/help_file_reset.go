package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_file_resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset changes to a path or file to the current revision, discarding your local changes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_file_resetCmd).Standalone()

	help_fileCmd.AddCommand(help_file_resetCmd)
}
