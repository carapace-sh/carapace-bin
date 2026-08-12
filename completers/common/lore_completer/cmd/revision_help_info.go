package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_help_infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get info about a revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_help_infoCmd).Standalone()

	revision_helpCmd.AddCommand(revision_help_infoCmd)
}
