package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/hg"
	"github.com/spf13/cobra"
)

var annotateCmd = &cobra.Command{
	Use:     "annotate",
	Short:   "show changeset information by line for each file",
	Aliases: []string{"blame"},
	GroupID: groups[group_file_content_management].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(annotateCmd).Standalone()

	annotateCmd.Flags().BoolP("changeset", "c", false, "list the changeset")
	annotateCmd.Flags().BoolP("date", "d", false, "list the date (short with -q)")
	annotateCmd.Flags().StringArrayP("exclude", "X", nil, "exclude names matching the given patterns")
	annotateCmd.Flags().BoolP("file", "f", false, "list the filename")
	annotateCmd.Flags().Bool("follow", false, "follow copies/renames and list the filename (DEPRECATED)")
	annotateCmd.Flags().BoolP("ignore-all-space", "w", false, "ignore white space when comparing lines")
	annotateCmd.Flags().BoolP("ignore-blank-lines", "B", false, "ignore changes whose lines are all blank")
	annotateCmd.Flags().BoolP("ignore-space-at-eol", "Z", false, "ignore changes in whitespace at EOL")
	annotateCmd.Flags().BoolP("ignore-space-change", "b", false, "ignore changes in the amount of white space")
	annotateCmd.Flags().StringArrayP("include", "I", nil, "include names matching the given patterns")
	annotateCmd.Flags().BoolP("line-number", "l", false, "show line number at the first appearance")
	annotateCmd.Flags().StringArrayP("line-range", "L", nil, "follow line range of specified file (EXPERIMENTAL)")
	annotateCmd.Flags().Bool("no-follow", false, "don't follow copies and renames")
	annotateCmd.Flags().BoolP("number", "n", false, "list the revision number (default)")
	annotateCmd.Flags().StringP("rev", "r", "", "annotate the specified revision")
	annotateCmd.Flags().String("skip", "", "revset to not display (EXPERIMENTAL)")
	annotateCmd.Flags().StringP("template", "T", "", "display with template")
	annotateCmd.Flags().BoolP("text", "a", false, "treat all files as text")
	annotateCmd.Flags().BoolP("user", "u", false, "list the author (long with -v)")
	rootCmd.AddCommand(annotateCmd)

	carapace.Gen(annotateCmd).FlagCompletion(carapace.ActionMap{
		"exclude": carapace.ActionFiles(),
		"include": carapace.ActionFiles(),
		"rev":     hg.ActionRevisions(),
	})

	carapace.Gen(annotateCmd).PositionalAnyCompletion(
		hg.ActionTrackedFiles(),
	)
}
