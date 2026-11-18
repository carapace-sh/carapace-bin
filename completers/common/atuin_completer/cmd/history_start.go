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
	history_startCmd.Flags().String("author-kind", "", "Whether a human or an AI agent ran this command")
	history_startCmd.Flags().Bool("command-from-env", false, "Collects the command from the `ATUIN_COMMAND_LINE` environment variable, which does not need escaping and is more compatible between OS and shells")
	history_startCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	history_startCmd.Flags().Bool("hook", false, "Passed by shell hooks; this flag disables logging to avoid corrupting the terminal and to minimize the amount of time the command takes to run")
	history_startCmd.Flags().String("intent", "", "Optional intent/rationale for running this command")
	history_startCmd.Flag("command-from-env").Hidden = true
	history_startCmd.Flag("hook").Hidden = true
	historyCmd.AddCommand(history_startCmd)

	carapace.Gen(history_startCmd).FlagCompletion(carapace.ActionMap{
		"author-kind": carapace.ActionValuesDescribed(
			"agent", "An AI agent ran this command",
			"user", "A human ran this command",
		),
	})
}
