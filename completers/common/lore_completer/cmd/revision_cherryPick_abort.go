package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_cherryPick_abortCmd = &cobra.Command{
	Use:   "abort",
	Short: "Abort a cherry-pick",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_cherryPick_abortCmd).Standalone()

	revision_cherryPick_abortCmd.Flags().BoolP("help", "h", false, "Print help")
	revision_cherryPickCmd.AddCommand(revision_cherryPick_abortCmd)
}
