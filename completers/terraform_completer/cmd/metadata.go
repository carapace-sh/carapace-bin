package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var metadataCmd = &cobra.Command{
	Use:   "metadata",
	Short: "Metadata related commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(metadataCmd).Standalone()

	rootCmd.AddCommand(metadataCmd)
}
