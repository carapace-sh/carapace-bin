package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var file_metadataCmd = &cobra.Command{
	Use:   "metadata",
	Short: "Manage metadata of a given file or directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_metadataCmd).Standalone()

	file_metadataCmd.Flags().BoolP("help", "h", false, "Print help")
	fileCmd.AddCommand(file_metadataCmd)
}
