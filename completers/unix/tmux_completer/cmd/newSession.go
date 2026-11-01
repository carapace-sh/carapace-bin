package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var newSessionCmd = &cobra.Command{
	Use:     "new-session",
	Aliases: []string{"new"},
	Short:   "create a new session",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(newSessionCmd).Standalone()

	newSessionCmd.Flags().BoolS("A", "A", false, "attach to existing session if it already exists")
	newSessionCmd.Flags().BoolS("D", "D", false, "with -A, detach other clients attached to session")
	newSessionCmd.Flags().BoolS("E", "E", false, "don't apply update-environment option")
	newSessionCmd.Flags().StringS("F", "F", "", "specify output format")
	newSessionCmd.Flags().BoolS("P", "P", false, "print information about new session after it is created")
	newSessionCmd.Flags().BoolS("X", "X", false, "with -D, send SIGHUP to the parent of the attached client")
	newSessionCmd.Flags().StringS("c", "c", "", "specify working directory for the session")
	newSessionCmd.Flags().BoolS("d", "d", false, "don't attach new session to current terminal")
	newSessionCmd.Flags().StringS("e", "e", "", "specify environment variable")
	newSessionCmd.Flags().StringS("f", "f", "", "specify client flags")
	newSessionCmd.Flags().StringS("n", "n", "", "specify initial window name")
	newSessionCmd.Flags().StringS("s", "s", "", "name the session")
	newSessionCmd.Flags().StringS("t", "t", "", "specify target session")
	newSessionCmd.Flags().StringS("x", "x", "", "specify width")
	newSessionCmd.Flags().StringS("y", "y", "", "specify height")
	rootCmd.AddCommand(newSessionCmd)

	carapace.Gen(newSessionCmd).FlagCompletion(carapace.ActionMap{
		"c": carapace.ActionDirectories(),
		"f": tmux.ActionClientFlags().UniqueList(","),
		"t": tmux.ActionSessions(),
	})

	carapace.Gen(newSessionCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
