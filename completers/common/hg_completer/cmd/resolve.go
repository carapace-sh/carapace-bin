package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/hg_completer/cmd/action"
	"github.com/spf13/cobra"
)

var resolveCmd = &cobra.Command{
	Use:     "resolve",
	Short:   "redo merges or set/view the merge status of files",
	GroupID: groups[group_working_directory_management].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resolveCmd).Standalone()

	resolveCmd.Flags().BoolP("all", "a", false, "select all unresolved files")
	resolveCmd.Flags().StringArrayP("exclude", "X", nil, "exclude names matching the given patterns")
	resolveCmd.Flags().StringArrayP("include", "I", nil, "include names matching the given patterns")
	resolveCmd.Flags().BoolP("list", "l", false, "list state of files needing merge")
	resolveCmd.Flags().BoolP("mark", "m", false, "mark files as resolved")
	resolveCmd.Flags().BoolP("no-status", "n", false, "hide status prefix")
	resolveCmd.Flags().Bool("re-merge", false, "re-merge files")
	resolveCmd.Flags().StringP("template", "T", "", "display with template")
	resolveCmd.Flags().StringP("tool", "t", "", "specify merge tool")
	resolveCmd.Flags().BoolP("unmark", "u", false, "mark files as unresolved")
	rootCmd.AddCommand(resolveCmd)

	carapace.Gen(resolveCmd).FlagCompletion(carapace.ActionMap{
		"exclude": carapace.ActionFiles(),
		"include": carapace.ActionFiles(),
		"tool":    action.ActionMergeTools(),
	})

	carapace.Gen(resolveCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
