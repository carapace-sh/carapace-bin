package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var setOptionCmd = &cobra.Command{
	Use:     "set-option",
	Aliases: []string{"set"},
	Short:   "set a session option",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setOptionCmd).Standalone()

	setOptionCmd.Flags().BoolS("F", "F", false, "expand formats in the option value")
	setOptionCmd.Flags().BoolS("U", "U", false, "unset a pane option across all panes in the window")
	setOptionCmd.Flags().BoolS("a", "a", false, "append to string options")
	setOptionCmd.Flags().BoolS("g", "g", false, "set a global session option")
	setOptionCmd.Flags().BoolS("o", "o", false, "prevent setting of an option that is already set")
	setOptionCmd.Flags().BoolS("p", "p", false, "change pane (no session) options")
	setOptionCmd.Flags().BoolS("q", "q", false, "suppress errors about unknown or ambiguous options")
	setOptionCmd.Flags().BoolS("s", "s", false, "change server (not session) options")
	setOptionCmd.Flags().StringS("t", "t", "", "specify target session")
	setOptionCmd.Flags().BoolS("u", "u", false, "unset a non-global option")
	setOptionCmd.Flags().BoolS("w", "w", false, "change window (not session) options")
	rootCmd.AddCommand(setOptionCmd)

	carapace.Gen(setOptionCmd).FlagCompletion(carapace.ActionMap{
		"t": tmux.ActionSessions(),
	})
}
