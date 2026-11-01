package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var listKeysCmd = &cobra.Command{
	Use:     "list-keys",
	Aliases: []string{"lsk"},
	Short:   "list all key-bindings",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listKeysCmd).Standalone()

	listKeysCmd.Flags().BoolS("1", "1", false, "list only the first matching key")
	listKeysCmd.Flags().StringS("F", "F", "", "specify format")
	listKeysCmd.Flags().BoolS("N", "N", false, "list only keys with attached notes")
	listKeysCmd.Flags().StringS("O", "O", "", "initial sort order")
	listKeysCmd.Flags().StringS("P", "P", "", "specify a prefix to print before each key")
	listKeysCmd.Flags().StringS("T", "T", "", "specify key table")
	listKeysCmd.Flags().BoolS("a", "a", false, "list the command for keys that do have a note")
	listKeysCmd.Flags().BoolS("r", "r", false, "reverse the sort order")
	rootCmd.AddCommand(listKeysCmd)

	carapace.Gen(listKeysCmd).FlagCompletion(carapace.ActionMap{
		"O": carapace.ActionValues("key", "modifier", "name"),
		"T": tmux.ActionKeyTables(),
	})
}
