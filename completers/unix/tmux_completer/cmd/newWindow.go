package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var newWindowCmd = &cobra.Command{
	Use:     "new-window",
	Aliases: []string{"neww"},
	Short:   "create a new window",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(newWindowCmd).Standalone()

	newWindowCmd.Flags().BoolS("E", "E", false, "create an empty pane without a running command")
	newWindowCmd.Flags().StringS("F", "F", "", "specify output format")
	newWindowCmd.Flags().BoolS("P", "P", false, "print information about new window after it has been created")
	newWindowCmd.Flags().BoolS("S", "S", false, "select window if name already exists")
	newWindowCmd.Flags().BoolS("a", "a", false, "insert new window at next index after target")
	newWindowCmd.Flags().BoolS("b", "b", false, "insert new window at next index before target")
	newWindowCmd.Flags().StringS("c", "c", "", "specify working directory for the session")
	newWindowCmd.Flags().BoolS("d", "d", false, "don't make the new window become the active one")
	newWindowCmd.Flags().StringS("e", "e", "", "specify environment variable")
	newWindowCmd.Flags().BoolS("k", "k", false, "destroy it if the specified window exists")
	newWindowCmd.Flags().StringS("n", "n", "", "specify a window name")
	newWindowCmd.Flags().StringS("t", "t", "", "specify target window")
	rootCmd.AddCommand(newWindowCmd)

	carapace.Gen(newWindowCmd).FlagCompletion(carapace.ActionMap{
		"c": carapace.ActionDirectories(),
		"t": tmux.ActionWindows(),
	})

	carapace.Gen(newWindowCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
