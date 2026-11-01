package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_repository_metadataCmd = &cobra.Command{
	Use:   "metadata",
	Short: "Repository metadata operations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_repository_metadataCmd).Standalone()

	help_repositoryCmd.AddCommand(help_repository_metadataCmd)
}
