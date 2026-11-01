package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_cherryPick_resolve_helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Print this message or the help of the given subcommand(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_cherryPick_resolve_helpCmd).Standalone()

	revision_cherryPick_resolveCmd.AddCommand(revision_cherryPick_resolve_helpCmd)
}
