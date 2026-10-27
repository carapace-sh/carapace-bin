package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_history_fixupCmd = &cobra.Command{
	Use:   "fixup",
	Short: "Apply staged changes onto a commit and replay descendants",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_history_fixupCmd).Standalone()

	help_historyCmd.AddCommand(help_history_fixupCmd)
}
