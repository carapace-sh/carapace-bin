package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var commandPromptCmd = &cobra.Command{
	Use:   "command-prompt",
	Short: "open the tmux command prompt in a client",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(commandPromptCmd).Standalone()

	commandPromptCmd.Flags().BoolS("1", "1", false, "only accept one key press")
	commandPromptCmd.Flags().BoolS("C", "C", false, "continue updating panes while prompt is displayed")
	commandPromptCmd.Flags().BoolS("F", "F", false, "expand template as a format")
	commandPromptCmd.Flags().StringS("I", "I", "", "comma-separated list of initial inputs")
	commandPromptCmd.Flags().BoolS("N", "N", false, "accept only numeric key presses")
	commandPromptCmd.Flags().BoolS("P", "P", false, "open prompt inside a pane instead of on the status line")
	commandPromptCmd.Flags().StringS("T", "T", "", "prompt type")
	commandPromptCmd.Flags().BoolS("b", "b", false, "show prompt in background")
	commandPromptCmd.Flags().BoolS("e", "e", false, "BSpace cancels an empty prompt")
	commandPromptCmd.Flags().BoolS("i", "i", false, "execute command every time prompt input changes")
	commandPromptCmd.Flags().BoolS("k", "k", false, "like -1 but translates key press to key name")
	commandPromptCmd.Flags().BoolS("l", "l", false, "disable splitting of inputs/prompts at commas")
	commandPromptCmd.Flags().StringS("p", "p", "", "comma-separated list of prompts")
	commandPromptCmd.Flags().StringS("t", "t", "", "specify target client")
	rootCmd.AddCommand(commandPromptCmd)

	carapace.Gen(commandPromptCmd).FlagCompletion(carapace.ActionMap{
		"T": carapace.ActionValues("command", "search"),
		"t": tmux.ActionClients(),
	})
}
