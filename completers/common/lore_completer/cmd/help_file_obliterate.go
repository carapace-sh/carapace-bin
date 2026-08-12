package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_file_obliterateCmd = &cobra.Command{
	Use:   "obliterate",
	Short: "Obliterate a file or fragment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_file_obliterateCmd).Standalone()

	help_fileCmd.AddCommand(help_file_obliterateCmd)
}
