package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_prevCmd = &cobra.Command{
	Use:   "prev",
	Short: "Move the working copy commit to the parent of the current revision.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_prevCmd).Standalone()

	helpCmd.AddCommand(help_prevCmd)
}
