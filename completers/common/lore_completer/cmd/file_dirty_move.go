package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var file_dirty_moveCmd = &cobra.Command{
	Use:   "move",
	Short: "Mark a file as moved (dirty)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_dirty_moveCmd).Standalone()

	file_dirty_moveCmd.Flags().BoolP("help", "h", false, "Print help")
	file_dirtyCmd.AddCommand(file_dirty_moveCmd)

	carapace.Gen(file_dirty_moveCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
