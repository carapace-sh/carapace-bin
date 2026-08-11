package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_branch_metadata_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set metadata on the branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_branch_metadata_setCmd).Standalone()

	help_branch_metadataCmd.AddCommand(help_branch_metadata_setCmd)
}
