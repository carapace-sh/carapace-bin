package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/hg"
	"github.com/spf13/cobra"
)

var diffCmd = &cobra.Command{
	Use:     "diff",
	Short:   "diff repository (or selected files)",
	GroupID: groups[group_file_content_management].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(diffCmd).Standalone()

	diffCmd.Flags().Bool("binary", false, "generate binary diffs in git mode (default)")
	diffCmd.Flags().StringP("change", "c", "", "change made by revision")
	diffCmd.Flags().StringArrayP("exclude", "X", nil, "exclude names matching the given patterns")
	diffCmd.Flags().String("from", "", "revision to diff from")
	diffCmd.Flags().BoolP("git", "g", false, "use git extended diff format (DEFAULT: diff.git)")
	diffCmd.Flags().BoolP("ignore-all-space", "w", false, "ignore white space when comparing lines")
	diffCmd.Flags().BoolP("ignore-blank-lines", "B", false, "ignore changes whose lines are all blank")
	diffCmd.Flags().Bool("ignore-changes-from-ancestors", false, "only compare the change made by the selected revision (EXPERIMENTAL)")
	diffCmd.Flags().BoolP("ignore-space-at-eol", "Z", false, "ignore changes in whitespace at EOL")
	diffCmd.Flags().BoolP("ignore-space-change", "b", false, "ignore changes in the amount of white space")
	diffCmd.Flags().StringArrayP("include", "I", nil, "include names matching the given patterns")
	diffCmd.Flags().Bool("nodates", false, "omit dates from diff headers")
	diffCmd.Flags().Bool("noprefix", false, "omit a/ and b/ prefixes from filenames")
	diffCmd.Flags().StringArrayP("rev", "r", nil, "revision (DEPRECATED)")
	diffCmd.Flags().Bool("reverse", false, "produce a diff that undoes the changes")
	diffCmd.Flags().String("root", "", "produce diffs relative to subdirectory")
	diffCmd.Flags().BoolP("show-function", "p", false, "show which function each change is in (DEFAULT: diff.showfunc)")
	diffCmd.Flags().Bool("stat", false, "output diffstat-style summary of changes")
	diffCmd.Flags().BoolP("subrepos", "S", false, "recurse into subrepositories")
	diffCmd.Flags().BoolP("text", "a", false, "treat all files as text")
	diffCmd.Flags().String("to", "", "revision to diff to")
	diffCmd.Flags().StringP("unified", "U", "", "number of lines of context to show")
	rootCmd.AddCommand(diffCmd)

	carapace.Gen(diffCmd).FlagCompletion(carapace.ActionMap{
		"change":  hg.ActionRevisions(),
		"exclude": carapace.ActionFiles(),
		"from":    hg.ActionRevisions(),
		"include": carapace.ActionFiles(),
		"rev":     hg.ActionRevisions(),
		"root":    carapace.ActionDirectories(),
		"to":      hg.ActionRevisions(),
	})

	carapace.Gen(diffCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
