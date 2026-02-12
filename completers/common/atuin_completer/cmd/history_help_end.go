package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var history_help_endCmd = &cobra.Command{
	Use:   "end",
	Short: "Finishes a new command in the history (adds time, exit code)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(history_help_endCmd).Standalone()

	history_helpCmd.AddCommand(history_help_endCmd)
}
