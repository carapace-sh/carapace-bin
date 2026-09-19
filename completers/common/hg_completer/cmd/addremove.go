package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var addremoveCmd = &cobra.Command{
	Use:     "addremove",
	Short:   "add all new files, delete all missing files",
	GroupID: groups[group_working_directory_management].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addremoveCmd).Standalone()

	addremoveCmd.Flags().BoolP("dry-run", "n", false, "do not perform actions, just print output")
	addremoveCmd.Flags().StringArrayP("exclude", "X", nil, "exclude names matching the given patterns")
	addremoveCmd.Flags().StringArrayP("include", "I", nil, "include names matching the given patterns")
	addremoveCmd.Flags().StringP("similarity", "s", "", "guess renamed files by similarity (0<=s<=100)")
	addremoveCmd.Flags().BoolP("subrepos", "S", false, "recurse into subrepositories")
	rootCmd.AddCommand(addremoveCmd)

	carapace.Gen(addremoveCmd).FlagCompletion(carapace.ActionMap{
		"exclude": carapace.ActionFiles(),
		"include": carapace.ActionFiles(),
	})
}
