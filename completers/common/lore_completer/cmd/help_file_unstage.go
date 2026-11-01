package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_file_unstageCmd = &cobra.Command{
	Use:   "unstage",
	Short: "Unstage changes to a file or directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_file_unstageCmd).Standalone()

	help_fileCmd.AddCommand(help_file_unstageCmd)
}
