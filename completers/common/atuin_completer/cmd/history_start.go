package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var history_startCmd = &cobra.Command{
	Use:   "start",
	Short: "Begins a new command in the history",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(history_startCmd).Standalone()

	history_startCmd.Flags().String("author", "", "Author of this command, eg `ellie`, `claude`, or `copilot`")
	history_startCmd.Flags().BoolP("help", "h", false, "Print help")
	history_startCmd.Flags().String("intent", "", "Optional intent/rationale for running this command")
	historyCmd.AddCommand(history_startCmd)
}
