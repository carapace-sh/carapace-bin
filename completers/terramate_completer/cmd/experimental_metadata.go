package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var experimental_metadataCmd = &cobra.Command{
	Use:   "metadata",
	Short: "Shows metadata available on the project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(experimental_metadataCmd).Standalone()

	experimentalCmd.AddCommand(experimental_metadataCmd)
}
