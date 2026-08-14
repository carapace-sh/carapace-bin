package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var listPanesCmd = &cobra.Command{
	Use:     "list-panes",
	Aliases: []string{"lsp"},
	Short:   "list panes of a window",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listPanesCmd).Standalone()

	listPanesCmd.Flags().StringS("F", "F", "", "specify output format")
	listPanesCmd.Flags().StringS("O", "O", "", "initial sort order")
	listPanesCmd.Flags().BoolS("a", "a", false, "list all panes the server possesses")
	listPanesCmd.Flags().StringS("f", "f", "", "filter items")
	listPanesCmd.Flags().BoolS("r", "r", false, "reverse sort order")
	listPanesCmd.Flags().BoolS("s", "s", false, "if specified, -t chooses a session")
	listPanesCmd.Flags().StringS("t", "t", "", "specify target")
	rootCmd.AddCommand(listPanesCmd)

	carapace.Gen(listPanesCmd).FlagCompletion(carapace.ActionMap{
		"O": carapace.ActionValues("name", "index", "size", "creation", "activity"),
		"t": tmux.ActionSessions(),
	})
}
