package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var file_help_dirty_copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "Mark a file as copied (dirty)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_help_dirty_copyCmd).Standalone()

	file_help_dirtyCmd.AddCommand(file_help_dirty_copyCmd)
}
