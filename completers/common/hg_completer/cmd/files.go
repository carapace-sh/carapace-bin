package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/hg"
	"github.com/spf13/cobra"
)

var filesCmd = &cobra.Command{
	Use:     "files",
	Short:   "list tracked files",
	GroupID: groups[group_working_directory_management].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(filesCmd).Standalone()

	filesCmd.Flags().StringArrayP("exclude", "X", nil, "exclude names matching the given patterns")
	filesCmd.Flags().StringArrayP("include", "I", nil, "include names matching the given patterns")
	filesCmd.Flags().BoolP("print0", "0", false, "end filenames with NUL, for use with xargs")
	filesCmd.Flags().StringP("rev", "r", "", "search the repository as it is in REV")
	filesCmd.Flags().BoolP("subrepos", "S", false, "recurse into subrepositories")
	filesCmd.Flags().StringP("template", "T", "", "display with template")
	rootCmd.AddCommand(filesCmd)

	carapace.Gen(filesCmd).FlagCompletion(carapace.ActionMap{
		"exclude": carapace.ActionFiles(),
		"include": carapace.ActionFiles(),
		"rev":     hg.ActionRevisions(),
	})

	carapace.Gen(filesCmd).PositionalAnyCompletion(
		hg.ActionTrackedFiles(),
	)
}
