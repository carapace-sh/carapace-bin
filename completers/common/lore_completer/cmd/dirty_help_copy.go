package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dirty_help_copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "Mark a file as copied (dirty)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dirty_help_copyCmd).Standalone()

	dirty_helpCmd.AddCommand(dirty_help_copyCmd)
}
