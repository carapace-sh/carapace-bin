package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_metadata_help_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get metadata from the repository (omit key to list all)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_metadata_help_getCmd).Standalone()

	repository_metadata_helpCmd.AddCommand(repository_metadata_help_getCmd)
}
