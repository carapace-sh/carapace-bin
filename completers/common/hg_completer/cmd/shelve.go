package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/hg_completer/cmd/action"
	"github.com/spf13/cobra"
)

var shelveCmd = &cobra.Command{
	Use:     "shelve",
	Short:   "save and set aside changes from the working directory",
	GroupID: groups[group_working_directory_management].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shelveCmd).Standalone()

	shelveCmd.Flags().BoolP("addremove", "A", false, "mark new/missing files as added/removed before shelving")
	shelveCmd.Flags().Bool("cleanup", false, "delete all shelved changes")
	shelveCmd.Flags().String("date", "", "shelve with the specified commit date")
	shelveCmd.Flags().BoolP("delete", "d", false, "delete the named shelved change(s)")
	shelveCmd.Flags().BoolP("edit", "e", false, "invoke editor on commit messages")
	shelveCmd.Flags().StringArrayP("exclude", "X", nil, "exclude names matching the given patterns")
	shelveCmd.Flags().StringArrayP("include", "I", nil, "include names matching the given patterns")
	shelveCmd.Flags().BoolP("interactive", "i", false, "interactive mode")
	shelveCmd.Flags().BoolP("keep", "k", false, "shelve, but keep changes in the working directory")
	shelveCmd.Flags().BoolP("list", "l", false, "list current shelves")
	shelveCmd.Flags().StringP("message", "m", "", "use text as shelve message")
	shelveCmd.Flags().StringP("name", "n", "", "use the given name for the shelved commit")
	shelveCmd.Flags().BoolP("patch", "p", false, "output patches for changes (provide the names of the shelved changes as positional arguments)")
	shelveCmd.Flags().Bool("stat", false, "output diffstat-style summary of changes (provide the names of the shelved changes as positional arguments)")
	shelveCmd.Flags().BoolP("unknown", "u", false, "store unknown files in the shelve")
	rootCmd.AddCommand(shelveCmd)

	carapace.Gen(shelveCmd).FlagCompletion(carapace.ActionMap{
		"exclude": carapace.ActionFiles(),
		"include": carapace.ActionFiles(),
		"name":    action.ActionShelves(),
	})

	carapace.Gen(shelveCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if shelveCmd.Flags().Changed("delete") || shelveCmd.Flags().Changed("patch") || shelveCmd.Flags().Changed("stat") {
				return action.ActionShelves()
			}
			return carapace.Batch(action.ActionChangedFiles(), action.ActionUntrackedFiles()).ToA()
		}),
	)
}
