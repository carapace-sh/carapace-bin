package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_revision_infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get info about a revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_revision_infoCmd).Standalone()

	help_revisionCmd.AddCommand(help_revision_infoCmd)
}
