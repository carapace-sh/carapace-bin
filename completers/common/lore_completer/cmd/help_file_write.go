package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_file_writeCmd = &cobra.Command{
	Use:   "write",
	Short: "Write data to a specific location",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_file_writeCmd).Standalone()

	help_fileCmd.AddCommand(help_file_writeCmd)
}
