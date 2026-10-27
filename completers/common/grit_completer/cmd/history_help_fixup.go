package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var history_help_fixupCmd = &cobra.Command{
	Use:   "fixup",
	Short: "Apply staged changes onto a commit and replay descendants",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(history_help_fixupCmd).Standalone()

	history_helpCmd.AddCommand(history_help_fixupCmd)
}
