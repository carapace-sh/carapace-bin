package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var file_help_unstageCmd = &cobra.Command{
	Use:   "unstage",
	Short: "Unstage changes to a file or directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_help_unstageCmd).Standalone()

	file_helpCmd.AddCommand(file_help_unstageCmd)
}
