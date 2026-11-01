package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_stackPanesCmd = &cobra.Command{
	Use:   "stack-panes",
	Short: "Stack pane ids Ids are a space separated list of pane ids. They should either be in the form of `terminal_<int>` (eg. terminal_1), `plugin_<int>` (eg. plugin_1) or bare integers in which case they'll be considered terminals (eg. 1 is the equivalent of terminal_1)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_stackPanesCmd).Standalone()

	action_stackPanesCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	actionCmd.AddCommand(action_stackPanesCmd)
}
