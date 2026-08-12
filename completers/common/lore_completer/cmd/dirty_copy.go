package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dirty_copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "Mark a file as copied (dirty)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dirty_copyCmd).Standalone()

	dirty_copyCmd.Flags().BoolP("help", "h", false, "Print help")
	dirtyCmd.AddCommand(dirty_copyCmd)

	carapace.Gen(dirty_copyCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
