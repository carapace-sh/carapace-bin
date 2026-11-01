package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_revision_cherryPick_abortCmd = &cobra.Command{
	Use:   "abort",
	Short: "Abort a cherry-pick",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_revision_cherryPick_abortCmd).Standalone()

	help_revision_cherryPickCmd.AddCommand(help_revision_cherryPick_abortCmd)
}
