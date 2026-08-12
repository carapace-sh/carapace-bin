package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var file_metadata_help_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get metadata from a file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_metadata_help_getCmd).Standalone()

	file_metadata_helpCmd.AddCommand(file_metadata_help_getCmd)
}
