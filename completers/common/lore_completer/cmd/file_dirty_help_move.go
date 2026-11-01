package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var file_dirty_help_moveCmd = &cobra.Command{
	Use:   "move",
	Short: "Mark a file as moved (dirty)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_dirty_help_moveCmd).Standalone()

	file_dirty_helpCmd.AddCommand(file_dirty_help_moveCmd)
}
