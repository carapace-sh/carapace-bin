package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var setWindowoptionCmd = &cobra.Command{
	Use:     "set-window-option",
	Aliases: []string{"setw"},
	Short:   "set a window option",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setWindowoptionCmd).Standalone()

	setWindowoptionCmd.Flags().BoolS("F", "F", false, "expand formats in the option value")
	setWindowoptionCmd.Flags().BoolS("a", "a", false, "append to string options")
	setWindowoptionCmd.Flags().BoolS("g", "g", false, "set a global window option")
	setWindowoptionCmd.Flags().BoolS("o", "o", false, "prevent setting of an option that is already set")
	setWindowoptionCmd.Flags().BoolS("q", "q", false, "suppress errors about unknown or ambiguous options")
	setWindowoptionCmd.Flags().StringS("t", "t", "", "specify target window")
	setWindowoptionCmd.Flags().BoolS("u", "u", false, "unset a non-global option")
	rootCmd.AddCommand(setWindowoptionCmd)

	carapace.Gen(setWindowoptionCmd).FlagCompletion(carapace.ActionMap{
		"t": tmux.ActionWindows(),
	})
}
