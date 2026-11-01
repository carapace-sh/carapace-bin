package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var file_help_dirty_moveCmd = &cobra.Command{
	Use:   "move",
	Short: "Mark a file as moved (dirty)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_help_dirty_moveCmd).Standalone()

	file_help_dirtyCmd.AddCommand(file_help_dirty_moveCmd)
}
