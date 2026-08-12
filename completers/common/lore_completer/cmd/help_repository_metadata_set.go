package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_repository_metadata_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set metadata on the repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_repository_metadata_setCmd).Standalone()

	help_repository_metadataCmd.AddCommand(help_repository_metadata_setCmd)
}
