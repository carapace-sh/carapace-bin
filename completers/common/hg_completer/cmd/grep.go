package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/hg"
	"github.com/spf13/cobra"
)

var grepCmd = &cobra.Command{
	Use:     "grep",
	Short:   "search for a pattern in specified files",
	GroupID: groups[group_file_content_management].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(grepCmd).Standalone()

	grepCmd.Flags().Bool("all", false, "an alias to --diff (DEPRECATED)")
	grepCmd.Flags().Bool("all-files", false, "include all files in the changeset while grepping (DEPRECATED)")
	grepCmd.Flags().BoolP("date", "d", false, "list the date (short with -q)")
	grepCmd.Flags().Bool("diff", false, "search revision differences for when the pattern was added or removed")
	grepCmd.Flags().StringArrayP("exclude", "X", nil, "exclude names matching the given patterns")
	grepCmd.Flags().BoolP("files-with-matches", "l", false, "print only filenames and revisions that match")
	grepCmd.Flags().BoolP("follow", "f", false, "follow changeset history, or file history across copies and renames")
	grepCmd.Flags().BoolP("ignore-case", "i", false, "ignore case when matching")
	grepCmd.Flags().StringArrayP("include", "I", nil, "include names matching the given patterns")
	grepCmd.Flags().BoolP("line-number", "n", false, "print matching line numbers")
	grepCmd.Flags().BoolP("print0", "0", false, "end fields with NUL")
	grepCmd.Flags().StringArrayP("rev", "r", nil, "search files changed within revision range")
	grepCmd.Flags().StringP("template", "T", "", "display with template")
	grepCmd.Flags().BoolP("text", "a", false, "treat all files as text")
	grepCmd.Flags().BoolP("user", "u", false, "list the author (long with -v)")
	rootCmd.AddCommand(grepCmd)

	carapace.Gen(grepCmd).FlagCompletion(carapace.ActionMap{
		"exclude": carapace.ActionFiles(),
		"include": carapace.ActionFiles(),
		"rev":     hg.ActionRevisions(),
		"user":    hg.ActionUsers(),
	})

	carapace.Gen(grepCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
