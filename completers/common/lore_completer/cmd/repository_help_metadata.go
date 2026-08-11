package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_help_metadataCmd = &cobra.Command{
	Use:   "metadata",
	Short: "Repository metadata operations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_help_metadataCmd).Standalone()

	repository_helpCmd.AddCommand(repository_help_metadataCmd)
}
