package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_rebaseCmd = &cobra.Command{
	Use:   "rebase",
	Short: "Move revisions to different parent(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_rebaseCmd).Standalone()

	helpCmd.AddCommand(help_rebaseCmd)
}
