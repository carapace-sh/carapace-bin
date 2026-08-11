package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_metadataCmd = &cobra.Command{
	Use:   "metadata",
	Short: "Repository metadata operations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_metadataCmd).Standalone()

	repository_metadataCmd.Flags().BoolP("help", "h", false, "Print help")
	repositoryCmd.AddCommand(repository_metadataCmd)
}
