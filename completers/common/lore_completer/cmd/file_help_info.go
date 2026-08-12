package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var file_help_infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get info about the given file or directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_help_infoCmd).Standalone()

	file_helpCmd.AddCommand(file_help_infoCmd)
}
