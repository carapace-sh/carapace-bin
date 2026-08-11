package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get info about a revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_infoCmd).Standalone()

	revision_infoCmd.Flags().Bool("delta", false, "Show delta information")
	revision_infoCmd.Flags().BoolP("help", "h", false, "Print help")
	revision_infoCmd.Flags().Bool("metadata", false, "Show file metadata information")
	revisionCmd.AddCommand(revision_infoCmd)
}
