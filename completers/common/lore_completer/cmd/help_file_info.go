package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_file_infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get info about the given file or directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_file_infoCmd).Standalone()

	help_fileCmd.AddCommand(help_file_infoCmd)
}
