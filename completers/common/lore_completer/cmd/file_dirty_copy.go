package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var file_dirty_copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "Mark a file as copied (dirty)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_dirty_copyCmd).Standalone()

	file_dirty_copyCmd.Flags().BoolP("help", "h", false, "Print help")
	file_dirtyCmd.AddCommand(file_dirty_copyCmd)
}
