package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_file_metadataCmd = &cobra.Command{
	Use:   "metadata",
	Short: "Manage metadata of a given file or directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_file_metadataCmd).Standalone()

	help_fileCmd.AddCommand(help_file_metadataCmd)
}
