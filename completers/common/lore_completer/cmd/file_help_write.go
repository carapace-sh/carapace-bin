package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var file_help_writeCmd = &cobra.Command{
	Use:   "write",
	Short: "Write data to a specific location",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_help_writeCmd).Standalone()

	file_helpCmd.AddCommand(file_help_writeCmd)
}
