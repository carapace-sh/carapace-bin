package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_rebaseCmd = &cobra.Command{
	Use:   "rebase",
	Short: "Replay local commits onto a new base (atomic ref update)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_rebaseCmd).Standalone()

	helpCmd.AddCommand(help_rebaseCmd)
}
