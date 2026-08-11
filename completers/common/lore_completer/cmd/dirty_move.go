package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dirty_moveCmd = &cobra.Command{
	Use:   "move",
	Short: "Mark a file as moved (dirty)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dirty_moveCmd).Standalone()

	dirty_moveCmd.Flags().BoolP("help", "h", false, "Print help")
	dirtyCmd.AddCommand(dirty_moveCmd)
}
