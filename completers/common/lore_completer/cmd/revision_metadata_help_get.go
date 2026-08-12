package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_metadata_help_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get metadata from a revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_metadata_help_getCmd).Standalone()

	revision_metadata_helpCmd.AddCommand(revision_metadata_help_getCmd)
}
