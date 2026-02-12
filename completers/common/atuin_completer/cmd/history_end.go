package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var history_endCmd = &cobra.Command{
	Use:   "end",
	Short: "Finishes a new command in the history (adds time, exit code)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(history_endCmd).Standalone()

	history_endCmd.Flags().StringP("duration", "d", "", "")
	history_endCmd.Flags().StringP("exit", "e", "", "")
	history_endCmd.Flags().BoolP("help", "h", false, "Print help")
	history_endCmd.MarkFlagRequired("exit")
	historyCmd.AddCommand(history_endCmd)
}
